package model

import (
	"database/sql"
	"net/http"
	"src/model/entities"
	"strings"
	"time"
)

type ApiResultHistoryModel interface {
	FetchApiResultHistory(r *http.Request) ([]*entities.ApiResultHistory, error)
	AddApiResultHistory(r entities.ApiResultHistory) (sql.Result, error)
	UpdateApiResultHistory(r *http.Request, e entities.ApiResultHistory) (sql.Result, error)
	DeleteApiResultHistory(r *http.Request) (sql.Result, error)
}

type apiResultHistoryModel struct {
}

func CreateApiResultHistoryModel() ApiResultHistoryModel {
	return &apiResultHistoryModel{}
}

func (am *apiResultHistoryModel) FetchApiResultHistory(r *http.Request) ([]*entities.ApiResultHistory, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	sql := `SELECT api_result_id, api_id, request_endpoint, request_param_string, request_date_time, response_status_code, response_data, response_date_time FROM ApiResultHistory`
	var params []interface{}
	var whereSql = ""

	apiIds := strings.Split(r.FormValue("api_result_id"), ",")
	if len(apiIds) > 0 && apiIds[0] != "" {
		whereSql += " WHERE api_result_id IN (" + strings.Repeat("?,", len(apiIds)-1) + "?)"
		for _, id := range apiIds {
			params = append(params, id)
		}
	}

	apiResultHistoryIds := strings.Split(r.FormValue("api_id"), ",")
	if len(apiResultHistoryIds) > 0 && apiResultHistoryIds[0] != "" {
		if whereSql == "" {
			whereSql += " WHERE"
		} else {
			whereSql += " AND"
		}
		whereSql += " api_id IN (" + strings.Repeat("?,", len(apiResultHistoryIds)-1) + "?)"
		for _, id := range apiResultHistoryIds {
			params = append(params, id)
		}
	}

	rows, err := Db.Query(sql+whereSql, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var apiResultHistory []*entities.ApiResultHistory

	for rows.Next() {
		var apiResultId int
		var apiId int
		var requestEndpoint string
		var requestParamString string
		var requestDateTime time.Time
		var responseStatusCode string
		var responseData string
		var responseDateTime time.Time

		if err := rows.Scan(&apiResultId, &apiId, &requestEndpoint, &requestParamString, &requestDateTime, &responseStatusCode, &responseData, &responseDateTime); err != nil {
			return nil, err
		}

		apiResultHistory = append(apiResultHistory, &entities.ApiResultHistory{
			ApiResultId:        apiResultId,
			ApiId:              apiId,
			RequestEndpoint:    requestEndpoint,
			RequestParamString: requestParamString,
			RequestDateTime:    requestDateTime,
			ResponseStatusCode: responseStatusCode,
			ResponseData:       responseData,
			ResponseDateTime:   responseDateTime,
		})
	}

	return apiResultHistory, nil
}

func (tm *apiResultHistoryModel) AddApiResultHistory(e entities.ApiResultHistory) (sql.Result, error) {
	req := entities.ApiResultHistory{
		ApiId:              e.ApiId,
		RequestEndpoint:    e.RequestEndpoint,
		RequestParamString: e.RequestParamString,
		RequestDateTime:    e.RequestDateTime,
		ResponseStatusCode: e.ResponseStatusCode,
		ResponseData:       e.ResponseData,
		ResponseDateTime:   e.ResponseDateTime,
	}

	sql := `INSERT INTO ApiResultHistory(api_id, request_endpoint, request_param_string, request_date_time, response_status_code, response_data, response_date_time) VALUES(?, ?, ?, ?, ?, ?, ?)`

	result, err := Db.Exec(sql, req.ApiId, req.RequestEndpoint, req.RequestParamString, req.RequestDateTime, req.ResponseStatusCode, req.ResponseData, req.ResponseDateTime)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *apiResultHistoryModel) UpdateApiResultHistory(r *http.Request, e entities.ApiResultHistory) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	req := entities.ApiResultHistory{
		RequestEndpoint:    e.RequestEndpoint,
		RequestParamString: e.RequestParamString,
		RequestDateTime:    e.RequestDateTime,
		ResponseStatusCode: e.ResponseStatusCode,
		ResponseData:       e.ResponseData,
		ResponseDateTime:   e.ResponseDateTime,
	}

	sql := `UPDATE ApiResultHistory SET request_method = ?, endpoint = ?, execution_interval_sec = ?`

	params := []interface{}{req.RequestEndpoint, req.RequestParamString, req.RequestDateTime}

	apiResultHistoryIds := strings.Split(r.FormValue("api_setting_id"), ",")
	if len(apiResultHistoryIds) > 0 && apiResultHistoryIds[0] != "" {
		sql += " WHERE api_setting_id IN (" + strings.Repeat("?,", len(apiResultHistoryIds)-1) + "?)"
		for _, id := range apiResultHistoryIds {
			params = append(params, id)
		}
	}

	result, err := Db.Exec(sql, params...)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *apiResultHistoryModel) DeleteApiResultHistory(r *http.Request) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	sql := `DELETE FROM ApiResultHistory`
	var params []interface{}
	var whereSql = ""

	apiIds := strings.Split(r.FormValue("api_id"), ",")
	if len(apiIds) > 0 && apiIds[0] != "" {
		whereSql += " WHERE api_id IN (" + strings.Repeat("?,", len(apiIds)-1) + "?)"
		for _, id := range apiIds {
			params = append(params, id)
		}
	}

	apiResultHistoryIds := strings.Split(r.FormValue("api_setting_id"), ",")
	if len(apiResultHistoryIds) > 0 && apiResultHistoryIds[0] != "" {
		if whereSql == "" {
			whereSql += " WHERE"
		} else {
			whereSql += " AND"
		}
		whereSql += " api_setting_id IN (" + strings.Repeat("?,", len(apiResultHistoryIds)-1) + "?)"
		for _, id := range apiResultHistoryIds {
			params = append(params, id)
		}
	}

	result, err := Db.Exec(sql+whereSql, params...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
