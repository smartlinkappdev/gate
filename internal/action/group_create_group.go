package action

import (
	"encoding/json"

	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
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

	user := entity.User{}
	options.Conn.Find(&user, "id = ?", options.UserID)

	group := entity.Group{
		Name:        params.Name,
		OwnerID:     options.UserID,
		Description: params.Description,
	}

	for _, uID := range params.Users {
		group.Users = append(group.Users, &entity.User{ID: uID})
	}

	for _, lID := range params.Links {
		group.Links = append(group.Links, &entity.Link{ID: lID})
	}

	group.Users = append(group.Users, &user)

	options.Conn.Create(&group)

	return json.Marshal(&group)
}
