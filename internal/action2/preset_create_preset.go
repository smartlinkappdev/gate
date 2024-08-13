package action2

import (
	"encoding/json"

	"cmd/gate/main.go/internal/jsonrpc"
)

type PresetCreatePresetsParams struct {
}

func PresetCreatePreset(options jsonrpc.Options) (json.RawMessage, error) {

	params := PresetCreatePresetsParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
