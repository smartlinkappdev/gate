package action

import (
	"encoding/json"

	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
)

type GroupCreateGroupParams struct {
	Name        string `form:"name"`
	Description string `form:"description"`
}

func GroupCreateGroup(options jsonrpc.Options) (json.RawMessage, error) {
	params := GroupCreateGroupParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	group := entity.Group{
		Name:        params.Name,
		OwnerID:     options.UserID,
		Description: params.Description,
	}

	user := entity.User{}
	options.Conn.Find(&user, "id = ?", options.UserID)

	group.Users = append(group.Users, &user)
	options.Conn.Create(&group)

	result, err := json.Marshal(&group)
	if err != nil {
		return nil, err
	}

	return result, nil
}
