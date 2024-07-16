package action

import (
	"encoding/json"

	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
)

type LinkGetLinksParams struct {
}

func LinkGetLinks(options jsonrpc.Options) (json.RawMessage, error) {
	links := make([]entity.Link, 0)

	options.Conn.Model(entity.User{ID: options.UserID}).Order("ID desc").Association("Links").Find(&links)

	data, _ := json.Marshal(&links)

	return data, nil
}
