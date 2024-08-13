package action

import (
	"encoding/json"
	"fmt"

	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
)

type UsersGetUsersParams struct {
	Email string
}

func UserGetUsers(options jsonrpc.Options) (json.RawMessage, error) {
	params := UsersGetUsersParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	var users []entity.User

	if params.Email != "" {
		options.Conn.Where("id != ?", options.UserID).Where("email ILIKE ?", fmt.Sprintf("%%%s%%", params.Email)).Order("id DESC").Find(&users)
		result, err := json.Marshal(users)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	options.Conn.Not("id = ?", options.UserID).Order("id ASC").Find(&users)

	result, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	return result, nil
}
