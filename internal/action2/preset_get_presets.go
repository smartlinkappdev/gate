package action2

import (
	"encoding/json"

	"cmd/gate/main.go/internal/jsonrpc"
)

type PresetGetPresetsParams struct {
}

func UserGetUser(options jsonrpc.Options) (json.RawMessage, error) {

	params := PresetGetPresetsParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
