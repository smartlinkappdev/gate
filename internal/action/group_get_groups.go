package action

import (
	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
	"encoding/json"
	"fmt"
)

type GroupGetGroupsParams struct {
	Type string `json:"type"`
}

type GroupGetGroupsDataGroup struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	UsersCount  int            `json:"users_count"`
	LinksCount  int            `json:"links_count"`
	Users       []*entity.User `json:"users"`
	Links       []*entity.Link `json:"links"`
}

type GroupGetGroupsData struct {
	My []GroupGetGroupsDataGroup `json:"my"`
}

type GetGroupsResponse struct {
	My  []entity.Group `json:"my"`
	All []entity.Group `json:"all"`
}

func GroupGetGroups(options jsonrpc.Options) (json.RawMessage, error) {
	params := GroupGetGroupsParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	my := make([]entity.Group, 0)

	query := options.Conn.Preload("Users").Preload("Links")
	switch params.Type {
	case "my":
		{
			query.Model(&entity.User{ID: options.UserID}).Order("id DESC").Association("Groups").Find(&my)
		}
	case "all":
		{
			query.Model(&entity.Group{}).Order("id DESC").Find(&my)
		}
	case "owner":
		{
			query.Model(&entity.Group{}).Where("owner_id = ?", options.UserID).Order("id DESC").Find(&my)
		}
	}

	//query.Find(&my)

	ggr := &GetGroupsResponse{My: my}
	fmt.Println(ggr)
	return json.Marshal(ggr)
}
