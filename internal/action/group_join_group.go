package action

import (
	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
	"encoding/json"
	"fmt"
)

type GroupJoinGroupParams struct {
	ID int `json:"id"`
}

func GroupJoinGroup(options jsonrpc.Options) (json.RawMessage, error) {
	group := entity.Group{}

	params := GroupJoinGroupParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	fmt.Println("params", params)

	options.Conn.Model(&group).Where("id = ?", params.ID).Find(&group)
	err = options.Conn.Model(&group).Association("Users").Append(&entity.User{ID: options.UserID})
	return nil, err
}
