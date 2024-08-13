package action

import (
	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
	"encoding/json"
)

type GroupInviteUsersParams struct {
	ID    int   `json:"id"`
	Users []int `json:"users"`
}

func GroupInviteUsers(options jsonrpc.Options) (json.RawMessage, error) {
	group := entity.Group{}

	params := GroupInviteUsersParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	users := make([]*entity.User, len(params.Users))
	for i, id := range params.Users {
		users[i] = &entity.User{ID: id}
	}

	options.Conn.Model(&group).Where("id = ?", params.ID).Find(&group)
	err = options.Conn.Model(&group).Association("Users").Append(&users)
	return nil, err
}
