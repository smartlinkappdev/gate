package action

import (
	"cmd/gate/main.go/internal/jsonrpc"
	"cmd/gate/main.go/internal/ym/model"
	"encoding/json"
	"fmt"
	"strings"
)

type MetricGetTableParams struct {
	Type       string   `json:"type"`
	Dimensions []string `json:"dimensions"`
	Metrics    []string `json:"metrics"`
	Interval   string   `json:"interval"`
	Range      struct {
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
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func MetricGetTable(options jsonrpc.Options) (json.RawMessage, error) {
	params := MetricGetTableParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("PARAMS", params)

	m := model.YmDatum{}

	type Result struct {
		Date            string `json:"date,omitempty"`
		Day             string `json:"day,omitempty" gorm:"day"`
		Hour            string `json:"hour,omitempty"`
		Browser         string `json:"browser,omitempty"`
		Device          string `json:"device,omitempty"`
		URLPath         string `json:"url_path,omitempty"`
		UTMSource       string `json:"utm_source,omitempty"`
		UTMMedium       string `json:"utm_medium,omitempty"`
		UTMCampaign     string `json:"utm_campaign,omitempty"`
		OperatingSystem string `json:"operating_system,omitempty"`
		Pageviews       int    `json:"pageviews,omitempty"`
		Users           int    `json:"users,omitempty"`
	}

	r := make([]Result, 0)

	queryParams := make([]string, 0)

	var results []string
	for _, col := range params.Metrics {
		result := fmt.Sprintf("sum(%s) as %s", col, col)
		results = append(results, result)
	}

	finalString := strings.Join(results, ", ")

	dimensionStrings := []string{}
	groupStrings := []string{}

	for _, str := range params.Dimensions {
		if str == "day" {
			dimensionStrings = append(dimensionStrings, "TO_CHAR(timestamp, 'YYYY-MM-DD') as day")
			//groupStrings = append(groupStrings, "TO_CHAR(timestamp, 'YYYY-MM-DD')")

		} else if str == "hour" {
			dimensionStrings = append(dimensionStrings, "TO_CHAR(timestamp, 'HH24:00') as hour")
			//groupStrings = append(groupStrings, "TO_CHAR(timestamp, 'HH24:00')")
		} else {
			dimensionStrings = append(dimensionStrings, str)
			//groupStrings = append(groupStrings, str)
			groupStrings = append(groupStrings, str)
		}

	}

	queryParams = append(queryParams, strings.Join(dimensionStrings, ", "), finalString)

	// Начало формирования запроса
	//query := options.Conn2.Model(&m).Select(queryParams).Group(strings.Join(dimensionStrings, ", "))
	query := options.Conn2.Model(&m).Select(queryParams).Group(fmt.Sprintf("timestamp, %s", strings.Join(groupStrings, ", ")))
	queryTotal := options.Conn2.Model(&m).Select(fmt.Sprintf(finalString))

	if len(params.Filter.URLPath.In) > 0 {
		query = query.Where("url_path IN ?", params.Filter.URLPath.In)
		queryTotal = queryTotal.Where("url_path IN ?", params.Filter.URLPath.In)
	}
	if len(params.Filter.URLPath.NotIn) > 0 {
		query = query.Not("url_path IN ?", params.Filter.URLPath.NotIn)
		queryTotal = queryTotal.Not("url_path IN ?", params.Filter.URLPath.NotIn)
	}

	if len(params.Filter.Browser.In) > 0 {
		query = query.Where("browser IN ?", params.Filter.Browser.In)
		queryTotal = queryTotal.Where("browser IN ?", params.Filter.Browser.In)
	}
	if len(params.Filter.Browser.NotIn) > 0 {
		query = query.Not("browser IN ?", params.Filter.Browser.NotIn)
		queryTotal = queryTotal.Not("browser IN ?", params.Filter.Browser.NotIn)
	}

	if len(params.Filter.OperatingSystem.In) > 0 {
		query = query.Where("operating_system IN ?", params.Filter.OperatingSystem.In)
		queryTotal = queryTotal.Where("operating_system IN ?", params.Filter.OperatingSystem.In)
	}
	if len(params.Filter.OperatingSystem.NotIn) > 0 {
		query = query.Not("operating_system IN ?", params.Filter.OperatingSystem.NotIn)
		queryTotal = queryTotal.Not("operating_system IN ?", params.Filter.OperatingSystem.NotIn)
	}

	if len(params.Filter.Device.In) > 0 {
		query = query.Where("device IN ?", params.Filter.Device.In)
		queryTotal = queryTotal.Where("device IN ?", params.Filter.Device.In)
	}
	if len(params.Filter.Device.NotIn) > 0 {
		query = query.Not("device IN ?", params.Filter.Device.NotIn)
		queryTotal = queryTotal.Not("device IN ?", params.Filter.Device.NotIn)
	}

	if len(params.Filter.UTMSource.In) > 0 {
		query = query.Where("utm_source IN ?", params.Filter.UTMSource.In)
		queryTotal = queryTotal.Where("utm_source IN ?", params.Filter.UTMSource.In)
	}
	if len(params.Filter.UTMSource.NotIn) > 0 {
		query = query.Not("utm_source IN ?", params.Filter.UTMSource.NotIn)
		queryTotal = queryTotal.Not("utm_source IN ?", params.Filter.UTMSource.NotIn)
	}

	if len(params.Filter.UTMMedium.In) > 0 {
		query = query.Where("utm_medium IN ?", params.Filter.UTMMedium.In)
		queryTotal = queryTotal.Where("utm_medium IN ?", params.Filter.UTMMedium.In)
	}
	if len(params.Filter.UTMMedium.NotIn) > 0 {
		query = query.Not("utm_medium IN ?", params.Filter.UTMMedium.NotIn)
		queryTotal = queryTotal.Not("utm_medium IN ?", params.Filter.UTMMedium.NotIn)
	}

	if len(params.Filter.UTMCampaign.In) > 0 {
		query = query.Where("utm_campaign IN ?", params.Filter.UTMCampaign.In)
		queryTotal = queryTotal.Where("utm_campaign IN ?", params.Filter.UTMCampaign.In)
	}
	if len(params.Filter.UTMCampaign.NotIn) > 0 {
		query = query.Not("utm_campaign IN ?", params.Filter.UTMCampaign.NotIn)
		queryTotal = queryTotal.Not("utm_campaign IN ?", params.Filter.UTMCampaign.NotIn)
	}

	var totalRows int64
	err = query.
		Where("timestamp BETWEEN ? AND ?", params.Range.GTE, params.Range.LTE).
		Count(&totalRows).Error
	if err != nil {
		fmt.Println("Error counting rows:", err)
	}

	err = query.
		Where("timestamp BETWEEN ? AND ?", params.Range.GTE, params.Range.LTE).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&r).Error
	if err != nil {
		fmt.Println(err)
	}

	type Totals struct {
		Rows      int `gorm:"rows" json:"rows,omitempty"`
		Users     int `json:"users,omitempty"`
		Pageviews int `json:"pageviews,omitempty"`
	}

	type Res struct {
		Rows   []Result `json:"rows"`
		Totals `json:"totals,omitempty"`
	}

	res := Res{
		Rows:   r,
		Totals: Totals{},
	}

	err = queryTotal.
		Where("timestamp BETWEEN ? AND ?", params.Range.GTE, params.Range.LTE).
		Find(&res.Totals).Error
	if err != nil {
		fmt.Println(err)
	}

	res.Totals.Rows = int(totalRows)

	result, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return result, nil
}
