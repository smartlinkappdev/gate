package action

import (
	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type LinkCreateLinkParams struct {
	Name string `json:"name"`
}

func LinkCreateLink(options jsonrpc.Options) (json.RawMessage, error) {
	params := LinkCreateLinkParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	link := entity.Link{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
		Name:      params.Name,
		OwnerID:   options.UserID,
	}

	link.Users = append(link.Users, &entity.User{ID: options.UserID})
	options.Conn.Create(&link)
	return nil, nil
}
