package action

import (
	"cmd/gate/main.go/internal/jsonrpc"
	"cmd/gate/main.go/internal/ym/model"
	"encoding/json"
	"fmt"
)

type MetricDownloadBarParams struct {
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

func MetricDownloadBar(options jsonrpc.Options) (json.RawMessage, error) {
	params := MetricDownloadBarParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("PARAMS", params)

	m := model.YmDatum{}

	type Result struct {
		Dimension string
		PageViews int
		Users     int
	}

	r := make([]Result, 0)

	queryParams := make([]string, 0)
	queryParams = append(queryParams, fmt.Sprintf("%s as dimension", params.Dimension), "sum(pageviews) as page_views", "sum(users) as users")

	// Начало формирования запроса
	query := options.Conn2.Model(&m).Select(queryParams).Group(fmt.Sprintf("dimension"))

	if len(params.Filter.URLPath.In) > 0 {
		query = query.Where("url_path IN ?", params.Filter.URLPath.In)
	}
	if len(params.Filter.URLPath.NotIn) > 0 {
		query = query.Not("url_path IN ?", params.Filter.URLPath.NotIn)
	}

	if len(params.Filter.Browser.In) > 0 {
		query = query.Where("browser IN ?", params.Filter.Browser.In)
	}
	if len(params.Filter.Browser.NotIn) > 0 {
		query = query.Not("browser IN ?", params.Filter.Browser.NotIn)
	}

	if len(params.Filter.OperatingSystem.In) > 0 {
		query = query.Where("operating_system IN ?", params.Filter.OperatingSystem.In)
	}
	if len(params.Filter.OperatingSystem.NotIn) > 0 {
		query = query.Not("operating_system IN ?", params.Filter.OperatingSystem.NotIn)
	}

	if len(params.Filter.Device.In) > 0 {
		query = query.Where("device IN ?", params.Filter.Device.In)
	}
	if len(params.Filter.Device.NotIn) > 0 {
		query = query.Not("device IN ?", params.Filter.Device.NotIn)
	}

	if len(params.Filter.UTMSource.In) > 0 {
		query = query.Where("utm_source IN ?", params.Filter.UTMSource.In)
	}
	if len(params.Filter.UTMSource.NotIn) > 0 {
		query = query.Not("utm_source IN ?", params.Filter.UTMSource.NotIn)
	}

	if len(params.Filter.UTMMedium.In) > 0 {
		query = query.Where("utm_medium IN ?", params.Filter.UTMMedium.In)
	}
	if len(params.Filter.UTMMedium.NotIn) > 0 {
		query = query.Not("utm_medium IN ?", params.Filter.UTMMedium.NotIn)
	}

	if len(params.Filter.UTMCampaign.In) > 0 {
		query = query.Where("utm_campaign IN ?", params.Filter.UTMCampaign.In)
	}
	if len(params.Filter.UTMCampaign.NotIn) > 0 {
		query = query.Not("utm_campaign IN ?", params.Filter.UTMCampaign.NotIn)
	}

	err = query.
		Where("timestamp BETWEEN ? AND ?", params.Range.GTE, params.Range.LTE).
		Find(&r).Error
	if err != nil {
		fmt.Println(err)
	}

	result, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return result, nil
}
