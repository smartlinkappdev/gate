package action2

import (
	"encoding/json"

	"cmd/gate/main.go/internal/jsonrpc"
)

type MetaGetMetaParams struct {
}

func MetaGetMeta(options jsonrpc.Options) (json.RawMessage, error) {

	params := MetaGetMetaParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
