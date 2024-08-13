package action

import (
	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
	"cmd/gate/main.go/internal/ym/model"
	"encoding/json"
	"fmt"
	"sync"
)

type GroupCreateGroupParams struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Users       []int  `form:"users"`
	Links       []int  `form:"links"`
}

func GroupCreateGroup(options jsonrpc.Options) (json.RawMessage, error) {
	params := GroupCreateGroupParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	//user := entity.User{}
	//options.Conn.Find(&user, "id = ?", options.UserID)

	// Получение метаинформации по ссылкам
	//type meta struct {
	//	URLPath []string `gorm:"url_path" json:"dimension"`
	//	Browser []string `json:"browser"`
	//}
	//m := meta{}

	var urlPaths []string
	options.Conn.Model(&entity.Link{}).
		Select("name").
		Where("id IN ?", params.Links).
		Find(&urlPaths)

	wg := sync.WaitGroup{}

	browser := make([]string, 0)
	device := make([]string, 0)
	UTMSource := make([]string, 0)
	UTMMedium := make([]string, 0)

	wg.Add(1)
	go func() {
		defer wg.Done()
		options.Conn2.Model(&model.YmDatum{}).Select("DISTINCT browser").Where("url_path IN ?", urlPaths).Find(&browser)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		options.Conn2.Model(&model.YmDatum{}).Select("DISTINCT device").Where("url_path IN ?", urlPaths).Find(&device)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		options.Conn2.Model(&model.YmDatum{}).Select("DISTINCT utm_source").Where("url_path IN ?", urlPaths).Find(&UTMSource)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		options.Conn2.Model(&model.YmDatum{}).Select("DISTINCT utm_medium").Where("url_path IN ?", urlPaths).Find(&UTMMedium)
	}()

	wg.Wait()

	raw, err := json.Marshal(struct {
		Browser   []string `json:"browser"`
		Device    []string `json:"device"`
		UTMSource []string `json:"utm_source"`
		UTMMedium []string `json:"utm_medium"`
	}{browser, device, UTMSource, UTMMedium})
	if err != nil {
		fmt.Println(err)
	}

	filter := entity.Filter{
		Name:        "Test filter",
		Description: "",
		Data:        raw,
	}

	preset := entity.Preset{
		Name:        "Основной",
		Description: "Базовый набор фильтров для группы",
		Owner:       &entity.User{ID: options.UserID},
		Filter:      &filter,
		Users:       nil,
		Links:       nil,
		Groups:      nil,
		Filters:     nil,
	}

	group := entity.Group{
		Name:        params.Name,
		OwnerID:     options.UserID,
		Description: params.Description,
		Preset:      &preset,
	}

	// Добавление пользователей в группу
	group.Users = append(group.Users, &entity.User{ID: options.UserID})
	for _, uID := range params.Users {
		group.Users = append(group.Users, &entity.User{ID: uID})
	}

	// Добавление ссылок в группу
	for _, lID := range params.Links {
		group.Links = append(group.Links, &entity.Link{ID: lID})
	}

	options.Conn.Create(&group)
	return json.Marshal(&group)
}
