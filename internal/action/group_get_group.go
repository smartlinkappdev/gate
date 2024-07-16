package action

import (
	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
	"encoding/json"
)

type GroupGetGroupParams struct {
	ID int `json:"id"`
}

func GroupGetGroup(options jsonrpc.Options) (json.RawMessage, error) {
	params := GroupGetGroupParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	var group entity.Group
	options.Conn.Model(&entity.Group{}).Where("id = ?", params.ID).First(&group)

	return json.Marshal(group)
}
