package action2

import (
	"cmd/gate/main.go/internal/jsonrpc"
	"cmd/gate/main.go/internal/sl/model"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type CHGetMetaParams struct {
	Dimension string `json:"dimension"`
	Filter    string `json:"filter"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	Range     struct {
		GTE string `json:"gte"`
		LTE string `json:"lte"`
	} `json:"range"`
}

func CHGetMeta(options jsonrpc.Options) (json.RawMessage, error) {
	hash := sha256.New()
	hash.Write(options.Params)
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	ctx := context.Background()
	value, err := options.Rdb.Get(ctx, hashString).Result()
	if value != "" {
		return []byte(value), err
	}

	params := CHGetMetaParams{}
	err = json.Unmarshal(options.Params, &params)
	if err != nil {
		return nil, err
	}

	ch := model.DmikAggMarketing{}
	query := options.Conn3ch.Model(&ch).Where("date_event BETWEEN ? AND ?", params.Range.GTE, params.Range.LTE)

	if params.Limit != 0 {
		query = query.Limit(params.Limit)
	}

	if params.Offset != 0 {
		query = query.Offset(params.Offset)
	}

	if params.Filter != "" {
		qString := fmt.Sprintf("%s LIKE ?", params.Dimension)
		filter := "%" + params.Filter + "%"
		query = query.Where(qString, filter)
	}

	result1 := make([]string, 0)
	query.Distinct(fmt.Sprintf("COALESCE(%s, '')", params.Dimension)).Find(&result1)

	result, err := json.Marshal(result1)
	if err != nil {
		return nil, err
	}

	options.Rdb.Set(ctx, hashString, string(result), time.Minute*15)

	return result, nil
}
