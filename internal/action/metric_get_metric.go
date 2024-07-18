package action

import (
	"cmd/gate/main.go/internal/entity"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"cmd/gate/main.go/internal/jsonrpc"
	"cmd/gate/main.go/internal/ym/model"
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

	type Result struct {
		Date      string
		Dimension string
		Total     int
	}

	r := make([]Result, 0)

	dateString := "DATE(timestamp)"
	if params.Interval == "hour" {
		dateString = "EXTRACT(EPOCH FROM timestamp)"
	}

	paths := make([]string, 0)
	options.Conn.Model(entity.User{ID: options.UserID}).Select("name").Association("Links").Find(&paths)

	queryParams := make([]string, 0)
	queryParams = append(queryParams, fmt.Sprintf("%s as date", dateString), fmt.Sprintf("%s as dimension", params.Dimension), "sum(users) as total")

	err = options.Conn2.Model(&m).
		Select(queryParams).
		Group(fmt.Sprintf("%s, dimension", dateString)).
		Where("url_path IN ?", paths).
		Where("timestamp BETWEEN ? AND ?", params.Range[0], params.Range[1]).
		Find(&r).Error
	if err != nil {
		fmt.Println(err)
	}

	d := make(map[string]map[int]int)
	// Заполнение вложенной карты
	for _, result := range r {
		if d[result.Dimension] == nil {
			d[result.Dimension] = make(map[int]int)
		}

		var t int
		if params.Interval == "day" {
			tt, err := time.Parse(time.RFC3339, result.Date)
			if err != nil {
				fmt.Println("Error parsing date:", err)
			}

			// Получение количества секунд с начала эпохи (UNIX timestamp)
			t = int(tt.Unix())
		} else {
			t, err = strconv.Atoi(result.Date[0:10])
			if err != nil {
				fmt.Println(err)
			}
		}

		d[result.Dimension][t*1000] = result.Total
	}

	result, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	return result, nil
}
