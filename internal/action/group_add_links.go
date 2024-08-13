package action

import (
	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
	"encoding/json"
)

type GroupAddLinksParams struct {
	ID    int   `json:"id"`
	Links []int `json:"links"`
}

func GroupAddLinks(options jsonrpc.Options) (json.RawMessage, error) {
	group := entity.Group{}

	params := GroupAddLinksParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	links := make([]*entity.Link, len(params.Links))
	for i, id := range params.Links {
		links[i] = &entity.Link{ID: id}
	}

	options.Conn.Model(&group).Where("id = ?", params.ID).Find(&group)
	err = options.Conn.Model(&group).Association("Links").Append(&links)
	return nil, err
}
