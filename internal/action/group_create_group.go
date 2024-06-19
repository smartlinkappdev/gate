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

	user := entity.User{}
	options.Conn.Find(&user, "id = ?", options.UserID)

	group := entity.Group{
		Name:        params.Name,
		OwnerID:     options.UserID,
		Description: params.Description,
	}

	group.Users = append(group.Users, &user)
	options.Conn.Create(&group)

	return json.Marshal(&group)
}
