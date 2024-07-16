package action

import (
	"encoding/json"
	"fmt"

	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
)

type UserGetUserParams struct {
	ID int `json:"id"`
}

type UserGetUserData struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	LinksCount  int    `json:"links_count"`
	GroupsCount int    `json:"groups_count"`
}

func UserGetUser(options jsonrpc.Options) (json.RawMessage, error) {
	id := options.UserID

	params := UserGetUserParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	if params.ID != 0 {
		id = params.ID
	}

	user := entity.User{}
	options.Conn.Model(&user).Where("id = ?", id).First(&user)

	LinksCount := int(options.Conn.Model(&user).Association("Links").Count())
	GroupsCount := int(options.Conn.Model(&user).Association("Groups").Count())

	data := UserGetUserData{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Role:        user.Role,
		LinksCount:  LinksCount,
		GroupsCount: GroupsCount,
	}

	fmt.Println("USER:", user)
	result, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}
