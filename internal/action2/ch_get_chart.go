package action2

import (
	"cmd/gate/main.go/internal/jsonrpc"
	"cmd/gate/main.go/internal/ym/model"
	"encoding/json"
	"fmt"
)

type CHGetChartParams struct {
	Type      string `json:"type"`
	Dimension string `json:"dimension"`
	Metric    string `json:"metric"`
	Interval  string `json:"interval"`
	Range     struct {
		GTE string `json:"gte"`
		LTE string `json:"lte"`
	}
	Filter struct {
		URLPath struct {
			In    []string `json:"in"`
			NotIn []string `json:"notIn"`
		} `json:"url_path"`
		Browser struct {
			In    []string `json:"in"`
			NotIn []string `json:"notIn"`
		} `json:"browser"`
		Device struct {
			In    []string `json:"in"`
			NotIn []string `json:"notIn"`
		} `json:"device"`
		OperatingSystem struct {
			In    []string `json:"in"`
			NotIn []string `json:"notIn"`
		} `json:"operating_system"`
		UTMSource struct {
			In    []string `json:"in"`
			NotIn []string `json:"notIn"`
		} `json:"utm_source"`
		UTMMedium struct {
			In    []string `json:"in"`
			NotIn []string `json:"notIn"`
		} `json:"utm_medium"`
		UTMCampaign struct {
			In    []string `json:"in"`
			NotIn []string `json:"notIn"`
		} `json:"utm_campaign"`
	} `json:"filter"`
	Limit int `json:"limit"`
}

func CHGetChart(options jsonrpc.Options) (json.RawMessage, error) {
	params := CHGetChartParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("PARAMS", params)

	m := model.Smartlink{}
	ms := make([]interface{}, 0)

	options.Conn3ch.Model(&m).Where("date1 BETWEEN ? AND ?", params.Range.GTE, params.Range.LTE).Find(&ms)

	result, err := json.Marshal(ms)
	if err != nil {
		return nil, err
	}
	return result, nil
}
