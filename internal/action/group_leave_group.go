package action

import (
	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
	"encoding/json"
)

type GroupLeftGroupParams struct {
	ID int `json:"id"`
}

func GroupLeftGroup(options jsonrpc.Options) (json.RawMessage, error) {
	group := entity.Group{}

	params := GroupJoinGroupParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	options.Conn.Model(&group).Where("id = ?", params.ID).Find(&group)

	err = options.Conn.Model(&group).Association("Users").Delete(&entity.User{ID: options.UserID})
	return nil, err
}
