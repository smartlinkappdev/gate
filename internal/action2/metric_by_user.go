package action2

import (
	"cmd/gate/main.go/internal/jsonrpc"
	"encoding/json"
)

type MetricByUserParams struct {
	Type      string `json:"type"` // my || all || owner
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

func MetricByUser(options jsonrpc.Options) (json.RawMessage, error) {
	return nil, nil
}
