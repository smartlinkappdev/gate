package action

import (
	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
	"encoding/json"
	"fmt"
)

type GroupGetGroupParams struct {
	ID int `json:"id"`
}

type GroupGetGroupResult struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       struct {
		ID int `json:"id"`
	} `json:"owner"`
	Users []struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	} `json:"users"`
	Preset entity.Preset `json:"preset"`
}

func GroupGetGroup(options jsonrpc.Options) (json.RawMessage, error) {
	params := GroupGetGroupParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	//group := GroupGetGroupResult{}

	var group entity.Group
	options.Conn.Model(&entity.Group{}).Where("id = ?", params.ID).Preload("Links").Preload("Preset").Preload("Owner").Preload("Users").First(&group)

	result := GroupGetGroupResult{
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		Owner: struct {
			ID int `json:"id"`
		}{
			ID: group.OwnerID,
		},
		Users: make([]struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
		}, len(group.Users)),
	}

	for i, user := range group.Users {
		result.Users[i] = struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
		}{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		}
	}

	fmt.Printf("%+v\n", result)

	return json.Marshal(group)
}
