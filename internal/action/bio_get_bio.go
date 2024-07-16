package action

import (
	"encoding/json"
	"fmt"

	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
)

type BioGetBioData struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	LinksCount  int    `json:"links_count"`
	GroupsCount int    `json:"groups_count"`
	ID          int    `json:"id"`
}

func BioGetBio(options jsonrpc.Options) (json.RawMessage, error) {
	id := options.UserID

	user := entity.User{}
	options.Conn.Model(&user).Where("id = ?", id).First(&user)

	LinksCount := int(options.Conn.Model(&entity.User{ID: id}).Association("Links").Count())
	GroupsCount := int(options.Conn.Model(&user).Association("Groups").Count())

	data := BioGetBioData{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Role:        user.Role,
		LinksCount:  LinksCount,
		GroupsCount: GroupsCount,
		ID:          options.UserID,
	}

	fmt.Println("USER:", user)
	result, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}
