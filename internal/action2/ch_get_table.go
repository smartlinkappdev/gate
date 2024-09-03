package action2

import (
	"cmd/gate/main.go/internal/jsonrpc"
	"cmd/gate/main.go/internal/sl/model"
	"encoding/json"
	"fmt"
	"strings"
)

type Filter struct {
	In    []string `json:"in"`
	NotIn []string `json:"notIn"`
}

type CHGetTableParams struct {
	Type       string   `json:"type"`
	Dimensions []string `json:"dimensions"`
	Metrics    []string `json:"metrics"`
	Interval   string   `json:"interval"`
	Range      struct {
		GTE string `json:"gte"`
		LTE string `json:"lte"`
	}
	Filter  map[string]Filter `json:"filter"`
	Filter1 struct {
		//DateEvent          Filter  `json:"date_event,omitempty" gorm:"date_event"`
		EventName          Filter `json:"event_name,omitempty" gorm:"event_name"`
		OsName             Filter `json:"os_name,omitempty" gorm:"os_name"`
		Product            Filter `json:"product,omitempty" gorm:"product"`
		Object             Filter `json:"object,omitempty" gorm:"object"`
		Utm                Filter `json:"utm,omitempty" gorm:"utm"`
		MediaSource        Filter `json:"media_source,omitempty" gorm:"media_source"`
		UtmSource          Filter `json:"utm_source,omitempty" gorm:"utm_source"`
		UtmMedium          Filter `json:"utm_medium,omitempty" gorm:"utm_medium"`
		UtmCampaign        Filter `json:"utm_campaign,omitempty" gorm:"utm_campaign"`
		UtmTerm            Filter `json:"utm_term,omitempty" gorm:"utm_term"`
		UtmContent         Filter `json:"utm_content,omitempty" gorm:"utm_content"`
		ID6                Filter `json:"id_6,omitempty" gorm:"id_6"`
		ID64               Filter `json:"id_6_4,omitempty" gorm:"id_6_4"`
		ID4                Filter `json:"id_4,omitempty" gorm:"id_4"`
		IDPerformance      Filter `json:"id_performance,omitempty" gorm:"id_performance"`
		IDText             Filter `json:"id_text,omitempty" gorm:"id_text"`
		Rk                 Filter `json:"rk,omitempty" gorm:"rk"`
		TranscriptEvents   Filter `json:"transcript_events,omitempty" gorm:"transcript_events"`
		Platform           Filter `json:"platform,omitempty" gorm:"platform"`
		NatureType         Filter `json:"nature_type,omitempty" gorm:"nature_type"`
		Team               Filter `json:"team,omitempty" gorm:"team"`
		Description        Filter `json:"description,omitempty" gorm:"description"`
		DescriptionHTML    Filter `json:"description_html,omitempty" gorm:"description_html"`
		ScreenPath         Filter `json:"screen_path,omitempty" gorm:"screen_path"`
		ScreenPathHTML     Filter `json:"screen_path_html,omitempty" gorm:"screen_path_html"`
		PropertiesDesc     Filter `json:"properties_desc,omitempty" gorm:"properties_desc"`
		PropertiesDescHTML Filter `json:"properties_desc_html,omitempty" gorm:"properties_desc_html"`
		CommentHTML        Filter `json:"comment_html,omitempty" gorm:"comment_html"`
		SectionCode        Filter `json:"section_code,omitempty" gorm:"section_code"`
		SectionName        Filter `json:"section_name,omitempty" gorm:"section_name"`
		CategoryCode       Filter `json:"category_code,omitempty" gorm:"category_code"`
		CategoryName       Filter `json:"category_name,omitempty" gorm:"category_name"`
		MinVersion         Filter `json:"min_version,omitempty" gorm:"min_version"`
		MaxVersion         Filter `json:"max_version,omitempty" gorm:"max_version"`
		EventType          Filter `json:"event_type,omitempty" gorm:"event_type"`
		Tags               Filter `json:"tags,omitempty" gorm:"tags"`
		Module             Filter `json:"module,omitempty" gorm:"module"`
		Flow               Filter `json:"flow,omitempty" gorm:"flow"`
		EventForAnalytic   Filter `json:"event_for_analytic,omitempty" gorm:"event_for_analytic"`
	} `json:"filter1"`
	OrderBy []map[string]string `json:"orderBy"`
	Limit   int                 `json:"limit"`
	Offset  int                 `json:"offset"`
}

func CHGetTable(options jsonrpc.Options) (json.RawMessage, error) {
	params := CHGetTableParams{}
	err := json.Unmarshal(options.Params, &params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("PARAMS", params)

	mm := model.DmikAggMarketing{
		DateEvent:          "",
		EventName:          "",
		OsName:             "",
		Product:            "",
		Object:             "",
		Utm:                "",
		MediaSource:        "",
		UtmSource:          "",
		UtmMedium:          "",
		UtmCampaign:        "",
		UtmTerm:            "",
		UtmContent:         "",
		Id6:                "",
		Id64:               "",
		Id4:                "",
		IDPerformance:      "",
		IDText:             "",
		Rk:                 "",
		TranscriptEvents:   "",
		Platform:           "",
		NatureType:         "",
		Team:               "",
		Description:        "",
		DescriptionHTML:    "",
		ScreenPath:         "",
		ScreenPathHTML:     "",
		PropertiesDesc:     "",
		PropertiesDescHTML: "",
		CommentHTML:        "",
		SectionCode:        "",
		SectionName:        "",
		CategoryCode:       "",
		CategoryName:       "",
		MinVersion:         "",
		MaxVersion:         "",
		EventType:          "",
		Tags:               "",
		Module:             "",
		Flow:               "",
		EventForAnalytic:   "",
		UsersSession:       0,
		EventsSession:      0,
		Users:              0,
		Events:             0,
	}

	type Result struct {
		ID                 int     `json:"id,omitempty" gorm:"id"`
		DateEvent          string  `json:"date_event,omitempty" gorm:"date_event"`
		EventName          string  `json:"event_name,omitempty" gorm:"event_name"`
		OsName             string  `json:"os_name,omitempty" gorm:"os_name"`
		Product            string  `json:"product,omitempty" gorm:"product"`
		Object             string  `json:"object,omitempty" gorm:"object"`
		Utm                string  `json:"utm,omitempty" gorm:"utm"`
		MediaSource        string  `json:"media_source,omitempty" gorm:"media_source"`
		UtmSource          string  `json:"utm_source,omitempty" gorm:"utm_source"`
		UtmMedium          string  `json:"utm_medium,omitempty" gorm:"utm_medium"`
		UtmCampaign        string  `json:"utm_campaign,omitempty" gorm:"utm_campaign"`
		UtmTerm            string  `json:"utm_term,omitempty" gorm:"utm_term"`
		UtmContent         string  `json:"utm_content,omitempty" gorm:"utm_content"`
		ID6                string  `json:"id_6,omitempty" gorm:"id_6"`
		Id64               string  `json:"id_6_4,omitempty" gorm:"id_6_4"`
		ID4                string  `json:"id_4,omitempty" gorm:"id_4"`
		IDPerformance      string  `json:"id_performance,omitempty" gorm:"id_performance"`
		IDText             string  `json:"id_text,omitempty" gorm:"id_text"`
		Rk                 string  `json:"rk,omitempty" gorm:"rk"`
		TranscriptEvents   string  `json:"transcript_events,omitempty" gorm:"transcript_events"`
		Platform           string  `json:"platform,omitempty" gorm:"platform"`
		NatureType         string  `json:"nature_type,omitempty" gorm:"nature_type"`
		Team               string  `json:"team,omitempty" gorm:"team"`
		Description        string  `json:"description,omitempty" gorm:"description"`
		DescriptionHTML    string  `json:"description_html,omitempty" gorm:"description_html"`
		ScreenPath         string  `json:"screen_path,omitempty" gorm:"screen_path"`
		ScreenPathHTML     string  `json:"screen_path_html,omitempty" gorm:"screen_path_html"`
		PropertiesDesc     *string `json:"properties_desc,omitempty" gorm:"properties_desc"`
		PropertiesDescHTML string  `json:"properties_desc_html,omitempty" gorm:"properties_desc_html"`
		CommentHTML        string  `json:"comment_html,omitempty" gorm:"comment_html"`
		SectionCode        string  `json:"section_code,omitempty" gorm:"section_code"`
		SectionName        string  `json:"section_name,omitempty" gorm:"section_name"`
		CategoryCode       string  `json:"category_code,omitempty" gorm:"category_code"`
		CategoryName       string  `json:"category_name,omitempty" gorm:"category_name"`
		MinVersion         string  `json:"min_version,omitempty" gorm:"min_version"`
		MaxVersion         string  `json:"max_version,omitempty" gorm:"max_version"`
		EventType          string  `json:"event_type,omitempty" gorm:"event_type"`
		Tags               string  `json:"tags,omitempty" gorm:"tags"`
		Module             string  `json:"module,omitempty" gorm:"module"`
		Flow               string  `json:"flow,omitempty" gorm:"flow"`
		EventForAnalytic   string  `json:"event_for_analytic,omitempty" gorm:"event_for_analytic"`
		UsersSession       int64   `json:"users_session,omitempty" gorm:"users_session"`
		EventsSession      int64   `json:"events_session,omitempty" gorm:"events_session"`
		Users              int64   `json:"users,omitempty" gorm:"users"`
		Events             int64   `json:"events,omitempty" gorm:"events"`
	}

	r := make([]Result, 0)

	queryParams := make([]string, 0)

	var results []string
	for _, col := range params.Metrics {
		result := fmt.Sprintf("CAST(sum(%s) AS BIGINT) as %s", col, col)
		results = append(results, result)
	}

	finalString := strings.Join(results, ", ")

	dimensionStrings := []string{}
	groupStrings := []string{}

	for _, str := range params.Dimensions {
		if str == "id" {
			dimensionStrings = append(dimensionStrings, "row_number() OVER () as id")
		} else {
			dimensionStrings = append(dimensionStrings, str)
			groupStrings = append(groupStrings, str)
		}
		//groupStrings = append(groupStrings, str)
	}

	queryParams = append(queryParams, strings.Join(dimensionStrings, ", "), finalString)

	// Начало формирования запроса
	//query := options.Conn2.Model(&m).Select(queryParams).Group(strings.Join(dimensionStrings, ", "))
	query := options.Conn3ch.Model(&mm).Select(queryParams).Group(fmt.Sprintf("%s", strings.Join(groupStrings, ", ")))
	queryTotal := options.Conn2.Model(&mm).Select(fmt.Sprintf(finalString))

	//if len(params.Filter.ID4.In) > 0 {
	//	query = query.Where("id_4 IN ?", params.Filter.ID4.In)
	//	queryTotal = queryTotal.Where("id_4 IN ?", params.Filter.ID4.In)
	//}
	//if len(params.Filter.ID4.NotIn) > 0 {
	//	query = query.Not("id_4 IN ?", params.Filter.ID4.NotIn)
	//	queryTotal = queryTotal.Not("id_4 IN ?", params.Filter.ID4.NotIn)
	//}
	//
	//if len(params.Filter.ID6.In) > 0 {
	//	query = query.Where("id_6 IN ?", params.Filter.ID6.In)
	//	queryTotal = queryTotal.Where("id_6 IN ?", params.Filter.ID6.In)
	//}
	//if len(params.Filter.ID6.NotIn) > 0 {
	//	query = query.Not("id_6 IN ?", params.Filter.ID6.NotIn)
	//	queryTotal = queryTotal.Not("id_6 IN ?", params.Filter.ID6.NotIn)
	//}
	//
	//if len(params.Filter.ID64.In) > 0 {
	//	query = query.Where("id_64 IN ?", params.Filter.ID64.In)
	//	queryTotal = queryTotal.Where("id_64 IN ?", params.Filter.ID64.In)
	//}
	//if len(params.Filter.ID64.NotIn) > 0 {
	//	query = query.Not("id_6_4 IN ?", params.Filter.ID64.NotIn)
	//	queryTotal = queryTotal.Not("id_6_4 IN ?", params.Filter.ID64.NotIn)
	//}
	//
	//if len(params.Filter.IDPerformance.In) > 0 {
	//	query = query.Where("id_performance IN ?", params.Filter.IDPerformance.In)
	//	queryTotal = queryTotal.Where("id_performance IN ?", params.Filter.IDPerformance.In)
	//}
	//if len(params.Filter.IDPerformance.NotIn) > 0 {
	//	query = query.Not("id_performance IN ?", params.Filter.IDPerformance.NotIn)
	//	queryTotal = queryTotal.Not("id_performance IN ?", params.Filter.IDPerformance.NotIn)
	//}
	//
	//if len(params.Filter.IDText.In) > 0 {
	//	query = query.Where("id_text IN ?", params.Filter.IDText.In)
	//	queryTotal = queryTotal.Where("id_text IN ?", params.Filter.IDText.In)
	//}
	//if len(params.Filter.IDText.NotIn) > 0 {
	//	query = query.Not("id_text IN ?", params.Filter.IDText.NotIn)
	//	queryTotal = queryTotal.Not("id_text IN ?", params.Filter.IDText.NotIn)
	//}
	//
	//if len(params.Filter.Object.In) > 0 {
	//	query = query.Where("object in (?)", params.Filter.Object.In)
	//	queryTotal = queryTotal.Where("object in (?)", params.Filter.Object.In)
	//}
	//if len(params.Filter.Object.NotIn) > 0 {
	//	query = query.Not("object IN ?", params.Filter.Object.NotIn)
	//	queryTotal = queryTotal.Not("object IN ?", params.Filter.Object.NotIn)
	//}
	//
	//if len(params.Filter.OsName.In) > 0 {
	//	query = query.Where("os_name IN ?", params.Filter.OsName.In)
	//	queryTotal = queryTotal.Where("os_name IN ?", params.Filter.OsName.In)
	//}
	//if len(params.Filter.OsName.NotIn) > 0 {
	//	query = query.Not("os_name IN ?", params.Filter.OsName.NotIn)
	//	queryTotal = queryTotal.Not("os_name IN ?", params.Filter.OsName.NotIn)
	//}
	//
	//if len(params.Filter.EventName.In) > 0 {
	//	query = query.Where("event_name IN ?", params.Filter.EventName.In)
	//	queryTotal = queryTotal.Where("event_name IN ?", params.Filter.EventName.In)
	//}
	//if len(params.Filter.EventName.NotIn) > 0 {
	//	query = query.Not("event_name IN ?", params.Filter.EventName.NotIn)
	//	queryTotal = queryTotal.Not("event_name IN ?", params.Filter.EventName.NotIn)
	//}
	//
	//if len(params.Filter.CategoryCode.In) > 0 {
	//	query = query.Where("category_code IN ?", params.Filter.CategoryCode.In)
	//	queryTotal = queryTotal.Where("category_code IN ?", params.Filter.CategoryCode.In)
	//}
	//if len(params.Filter.CategoryCode.NotIn) > 0 {
	//	query = query.Not("category_code IN ?", params.Filter.CategoryCode.NotIn)
	//	queryTotal = queryTotal.Not("category_code IN ?", params.Filter.CategoryCode.NotIn)
	//}
	//
	//if len(params.Filter.CategoryName.In) > 0 {
	//	query = query.Where("category_name IN ?", params.Filter.CategoryName.In)
	//	queryTotal = queryTotal.Where("category_name IN ?", params.Filter.CategoryName.In)
	//}
	//if len(params.Filter.CategoryName.NotIn) > 0 {
	//	query = query.Not("category_name IN ?", params.Filter.CategoryName.NotIn)
	//	queryTotal = queryTotal.Not("category_name IN ?", params.Filter.CategoryName.NotIn)
	//}
	//
	//if len(params.Filter.CommentHTML.In) > 0 {
	//	query = query.Where("comment_html IN ?", params.Filter.CommentHTML.In)
	//	queryTotal = queryTotal.Where("comment_html IN ?", params.Filter.CommentHTML.In)
	//}
	//if len(params.Filter.CommentHTML.NotIn) > 0 {
	//	query = query.Not("comment_html IN ?", params.Filter.CommentHTML.NotIn)
	//	queryTotal = queryTotal.Not("comment_html IN ?", params.Filter.CommentHTML.NotIn)
	//}

	type ilter struct {
		Utm                Filter `json:"utm,omitempty" gorm:"utm"`
		MediaSource        Filter `json:"media_source,omitempty" gorm:"media_source"`
		UtmSource          Filter `json:"utm_source,omitempty" gorm:"utm_source"`
		UtmMedium          Filter `json:"utm_medium,omitempty" gorm:"utm_medium"`
		UtmCampaign        Filter `json:"utm_campaign,omitempty" gorm:"utm_campaign"`
		UtmTerm            Filter `json:"utm_term,omitempty" gorm:"utm_term"`
		UtmContent         Filter `json:"utm_content,omitempty" gorm:"utm_content"`
		Rk                 Filter `json:"rk,omitempty" gorm:"rk"`
		TranscriptEvents   Filter `json:"transcript_events,omitempty" gorm:"transcript_events"`
		Platform           Filter `json:"platform,omitempty" gorm:"platform"`
		NatureType         Filter `json:"nature_type,omitempty" gorm:"nature_type"`
		Team               Filter `json:"team,omitempty" gorm:"team"`
		Description        Filter `json:"description,omitempty" gorm:"description"`
		DescriptionHTML    Filter `json:"description_html,omitempty" gorm:"description_html"`
		ScreenPath         Filter `json:"screen_path,omitempty" gorm:"screen_path"`
		ScreenPathHTML     Filter `json:"screen_path_html,omitempty" gorm:"screen_path_html"`
		PropertiesDesc     Filter `json:"properties_desc,omitempty" gorm:"properties_desc"`
		PropertiesDescHTML Filter `json:"properties_desc_html,omitempty" gorm:"properties_desc_html"`
		SectionCode        Filter `json:"section_code,omitempty" gorm:"section_code"`
		SectionName        Filter `json:"section_name,omitempty" gorm:"section_name"`
		MinVersion         Filter `json:"min_version,omitempty" gorm:"min_version"`
		MaxVersion         Filter `json:"max_version,omitempty" gorm:"max_version"`
		EventType          Filter `json:"event_type,omitempty" gorm:"event_type"`
		Tags               Filter `json:"tags,omitempty" gorm:"tags"`
		Module             Filter `json:"module,omitempty" gorm:"module"`
		Flow               Filter `json:"flow,omitempty" gorm:"flow"`
		EventForAnalytic   Filter `json:"event_for_analytic,omitempty" gorm:"event_for_analytic"`
	}

	for key, value := range params.Filter {
		if len(value.In) > 0 {
			query = query.Where(fmt.Sprintf("%s IN ?", key), value.In)
			queryTotal = queryTotal.Where(fmt.Sprintf("%s IN ?", key), value.In)
		}
		if len(value.NotIn) > 0 {
			query = query.Not(fmt.Sprintf("%s IN ?", key), value.NotIn)
			queryTotal = queryTotal.Not(fmt.Sprintf("%s IN ?", key), value.NotIn)
		}
	}

	type Totals struct {
		Rows          int   `gorm:"rows" json:"rows,omitempty"`
		UsersSession  int64 `json:"users_session,omitempty" gorm:"users_session"`
		EventsSession int64 `json:"events_session,omitempty" gorm:"events_session"`
		Users         int64 `json:"users,omitempty" gorm:"users"`
		Events        int64 `json:"events,omitempty" gorm:"events"`
	}

	var totalRows int64

	err = query.
		Where("date_event BETWEEN ? AND ?", params.Range.GTE, params.Range.LTE).
		Count(&totalRows).Error
	if err != nil {
		fmt.Println("Error counting rows:", err)
	}

	if params.Limit != 0 {
		query = query.Limit(params.Limit)
	}

	for _, order := range params.OrderBy {
		for key, value := range order {
			query = query.Order(fmt.Sprintf("%s %s", key, value))
		}
	}

	err = query.
		Where("date_event BETWEEN ? AND ?", params.Range.GTE, params.Range.LTE).
		Offset(params.Offset).
		Find(&r).Error
	if err != nil {
		fmt.Println(err)
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
		Where("date_event BETWEEN ? AND ?", params.Range.GTE, params.Range.LTE).
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
