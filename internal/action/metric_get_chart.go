package action

import (
	"cmd/gate/main.go/internal/jsonrpc"
	"cmd/gate/main.go/internal/ym/model"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"time"
)

type MetricGetChartParams struct {
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

func MetricGetChart(options jsonrpc.Options) (json.RawMessage, error) {
	params := MetricGetChartParams{}
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

	queryParams := make([]string, 0)
	queryParams = append(queryParams, fmt.Sprintf("%s as date", dateString), fmt.Sprintf("%s as dimension", params.Dimension), fmt.Sprintf("sum(%s) as total", params.Metric))

	// Начало формирования запроса
	query := options.Conn2.Model(&m).Select(queryParams).Group(fmt.Sprintf("%s, dimension", dateString))

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

	d := make(map[string]map[int]int)

	forLimit := make(map[string]int)

	res := make(map[string]map[int]int)
	res["total"] = make(map[int]int)

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

		res["total"][t*1000] += result.Total

		forLimit[result.Dimension] += result.Total
		d[result.Dimension][t*1000] = result.Total
	}

	topKeys := getTopNKeys(forLimit, params.Limit)
	fmt.Println(topKeys)

	for _, key := range topKeys {
		res[key] = d[key]
	}

	result, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func getTopNKeys(data map[string]int, n int) []string {
	// Создаем слайс для хранения пар (ключ, значение)
	type kv struct {
		Key   string
		Value int
	}
	var kvList []kv

	// Заполняем слайс парами (ключ, значение)
	for k, v := range data {
		kvList = append(kvList, kv{k, v})
	}

	// Сортируем слайс по значениям в порядке убывания
	sort.Slice(kvList, func(i, j int) bool {
		return kvList[i].Value > kvList[j].Value
	})

	// Создаем слайс для хранения ключей с максимальными значениями
	var topKeys []string
	for i := 0; i < n && i < len(kvList); i++ {
		topKeys = append(topKeys, kvList[i].Key)
	}

	return topKeys
}
