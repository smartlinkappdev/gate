package action

import (
	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
	"encoding/json"
)

type GroupDeleteUserParams struct {
	GroupID int `json:"group_id"`
	UserID  int `json:"user_id"`
}

func GroupDeleteUser(options jsonrpc.Options) (json.RawMessage, error) {
	group := entity.Group{}

	params := GroupDeleteUserParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	options.Conn.Model(&group).Where("id = ?", params.GroupID).Find(&group)
	err = options.Conn.Model(&group).Association("Users").Delete(&entity.User{ID: params.UserID})
	return nil, nil
}
