package action

import (
	"encoding/json"
	"time"

	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
)

type LinkGetLinkParams struct {
	ID int `json:"id"`
}

type LinkGetLinkResult struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
}

func LinkGetLink(options jsonrpc.Options) (json.RawMessage, error) {
	params := LinkGetLinkParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	link := entity.Link{}
	options.Conn.Model(link).Where("id = ?", params.ID).First(&link)

	result := LinkGetLinkResult{
		ID:          link.ID,
		CreatedAt:   link.CreatedAt,
		Name:        link.Name,
		Path:        link.Path,
		Description: link.Description,
		Image:       link.Image,
	}

	data, err := json.Marshal(&result)
	if err != nil {
		return nil, err
	}

	return data, nil
}
