package action

import (
	"cmd/gate/main.go/internal/jsonrpc"
	"cmd/gate/main.go/internal/ym/model"
	"encoding/json"
	"fmt"
	"time"
)

type MetricGetMetricParams struct {
	Type      string `json:"type"`
	Dimension string `json:"dimension"`
	Metric    string `json:"metric"`
	Interval  string `json:"interval"`
	Range     []time.Time
}

func MetricGetMetric(options jsonrpc.Options) (json.RawMessage, error) {
	params := MetricGetMetricParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("PARAMS", params)

	m := model.YmDatum{}
	//res := make([]model.YmDatum, 5)
	//options.Conn2.Model(&m).Limit(5).Find(&res)

	type Result struct {
		Date    string
		UrlPath string `gorm:"url_path"`
		Total   int
	}

	r := make([]Result, 0)
	err = options.Conn2.Model(&m).Select("EXTRACT(EPOCH FROM timestamp) as date, url_path, sum(users) as total").Group("EXTRACT(EPOCH FROM timestamp), url_path").Where("url_path = ?", "sms/mbp").Where("timestamp BETWEEN ? AND ?", "2024-05-25", "2024-05-27").Find(&r).Error
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("METRIC:", res)
	//fmt.Println("METRIC_66666:", r)
	result, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return result, nil
}
