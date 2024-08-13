package action

import (
	"encoding/json"

	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
)

type LinkGetLinksParams struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
}

func LinkGetLinks(options jsonrpc.Options) (json.RawMessage, error) {
	params := LinkGetLinksParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	links := make([]entity.Link, 0)

	query := options.Conn

	switch params.Type {
	case "Мои":
		{
			query = query.Model(entity.User{ID: options.UserID})
			query.Order("ID desc").Association("Links").Find(&links)

			data, _ := json.Marshal(&links)
			return data, nil
		}
	case "Доступно":
		{

		}
	case "Все":
		{
			query = query.Model(entity.Link{}).Order("ID desc").Limit(10).Offset(params.Offset).Find(&links)
			data, _ := json.Marshal(&links)
			return data, nil
		}

	}

	query.Order("ID desc").Association("Links").Find(&links)
	data, _ := json.Marshal(&links)

	return data, nil
}
