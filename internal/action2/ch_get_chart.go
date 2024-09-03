package action2

import (
	"cmd/gate/main.go/internal/jsonrpc"
	"cmd/gate/main.go/internal/sl/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"strconv"
)

type CHGetChartParams struct {
	Dimension string `json:"dimension"`
	Metric    string `json:"metric"`
	Range     struct {
		GTE string `json:"gte"`
		LTE string `json:"lte"`
	}
	Filter map[string]Filter `json:"filter"`
	Limit  int               `json:"limit"`
}

func CHGetChart(options jsonrpc.Options) (json.RawMessage, error) {
	params := CHGetChartParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	m := model.DmikAggMarketing{}

	type Row struct {
		Date      string `json:"date" gorm:"date_event"`
		Dimension string `json:"dimension"`
		Metric    int64  `json:"metric" gorm:"metric"`
	}
	ms := make([]Row, 0)

	selectString := fmt.Sprintf("toUnixTimestamp(date_event) as date, %s as dimension, sum(%s) as metric", params.Dimension, params.Metric)
	dateString := fmt.Sprintf("date_event BETWEEN '%s' AND '%s'", params.Range.GTE, params.Range.LTE)

	query := options.Conn3ch.Model(&m).Select(selectString).Where(dateString)

	for key, value := range params.Filter {
		if len(value.In) > 0 {
			query = query.Where(fmt.Sprintf("%s IN ?", key), value.In)
		}
		if len(value.NotIn) > 0 {
			query = query.Not(fmt.Sprintf("%s IN ?", key), value.NotIn)
		}
	}

	// этот способ используется так как формат метрик в базе big.Int
	rows, err := query.Group("dimension, date").Order("date, metric desc").Rows()
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)

	for rows.Next() {
		r := Row{}
		metric := big.Int{}
		err = rows.Scan(&r.Date, &r.Dimension, &metric)
		if err != nil {
			return nil, err
		}
		r.Metric = metric.Int64()
		ms = append(ms, r)
	}

	d := make(map[string]map[int]int64)

	forLimit := make(map[string]int64)

	res := make(map[string]map[int]int64)
	res["total"] = make(map[int]int64)

	// Заполнение вложенной карты
	for _, result := range ms {
		if d[result.Dimension] == nil {
			d[result.Dimension] = make(map[int]int64)
		}

		var t int

		t, err = strconv.Atoi(result.Date[0:10])
		if err != nil {
			fmt.Println(err)
		}

		res["total"][t*1000] += result.Metric

		forLimit[result.Dimension] += result.Metric
		d[result.Dimension][t*1000] = result.Metric
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
	//
	//result, err := json.Marshal(ms)
	//if err != nil {
	//	return nil, err
	//}
	//return result, nil
}

func getTopNKeys(data map[string]int64, n int) []string {
	// Создаем слайс для хранения пар (ключ, значение)
	type kv struct {
		Key   string
		Value int64
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
