package model

import (
	"database/sql"
	"net/http"
	"src/model/entities"
	"strings"
)

type ApiSettingModel interface {
	FetchApiSetting(r *http.Request) ([]*entities.ApiSetting, error)
	AddApiSetting(r entities.ApiSetting) (sql.Result, error)
	UpdateApiSetting(r *http.Request, e entities.ApiSetting) (sql.Result, error)
	DeleteApiSetting(r *http.Request) (sql.Result, error)
}

type apiSettingModel struct {
}

func CreateApiSettingModel() ApiSettingModel {
	return &apiSettingModel{}
}

func (am *apiSettingModel) FetchApiSetting(r *http.Request) ([]*entities.ApiSetting, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	sql := `SELECT api_setting_id, api_id, request_method, endpoint, execution_interval_sec FROM ApiSetting`
	var params []interface{}
	var whereSql = ""

	apiIds := strings.Split(r.FormValue("api_id"), ",")
	if len(apiIds) > 0 && apiIds[0] != "" {
		whereSql += " WHERE api_id IN (" + strings.Repeat("?,", len(apiIds)-1) + "?)"
		for _, id := range apiIds {
			params = append(params, id)
		}
	}

	apiSettingIds := strings.Split(r.FormValue("api_setting_id"), ",")
	if len(apiSettingIds) > 0 && apiSettingIds[0] != "" {
		if whereSql == "" {
			whereSql += " WHERE"
		} else {
			whereSql += " AND"
		}
		whereSql += " api_setting_id IN (" + strings.Repeat("?,", len(apiSettingIds)-1) + "?)"
		for _, id := range apiSettingIds {
			params = append(params, id)
		}
	}

	rows, err := Db.Query(sql+whereSql, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var apiSetting []*entities.ApiSetting

	for rows.Next() {
		var apiSettingId int
		var apiId int
		var requestMethod string
		var endpoint string
		var executionIntervalSec int

		if err := rows.Scan(&apiSettingId, &apiId, &requestMethod, &endpoint, &executionIntervalSec); err != nil {
			return nil, err
		}

		apiSetting = append(apiSetting, &entities.ApiSetting{
			ApiSettingId:         apiSettingId,
			ApiId:                apiId,
			RequestMethod:        requestMethod,
			Endpoint:             endpoint,
			ExecutionIntervalSec: executionIntervalSec,
		})
	}

	return apiSetting, nil
}

func (tm *apiSettingModel) AddApiSetting(e entities.ApiSetting) (sql.Result, error) {
	req := entities.ApiSetting{
		ApiId:                e.ApiId,
		RequestMethod:        e.RequestMethod,
		Endpoint:             e.Endpoint,
		ExecutionIntervalSec: e.ExecutionIntervalSec,
	}

	sql := `INSERT INTO ApiSetting(api_id, request_method, endpoint, execution_interval_sec) VALUES(?, ?, ?, ?)`

	result, err := Db.Exec(sql, req.ApiId, req.RequestMethod, req.Endpoint, req.ExecutionIntervalSec)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *apiSettingModel) UpdateApiSetting(r *http.Request, e entities.ApiSetting) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	req := entities.ApiSetting{
		RequestMethod:        e.RequestMethod,
		Endpoint:             e.Endpoint,
		ExecutionIntervalSec: e.ExecutionIntervalSec,
	}

	sql := `UPDATE ApiSetting SET request_method = ?, endpoint = ?, execution_interval_sec = ?`

	params := []interface{}{req.RequestMethod, req.Endpoint, req.ExecutionIntervalSec}

	apiSettingIds := strings.Split(r.FormValue("api_setting_id"), ",")
	if len(apiSettingIds) > 0 && apiSettingIds[0] != "" {
		sql += " WHERE api_setting_id IN (" + strings.Repeat("?,", len(apiSettingIds)-1) + "?)"
		for _, id := range apiSettingIds {
			params = append(params, id)
		}
	}

	result, err := Db.Exec(sql, params...)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *apiSettingModel) DeleteApiSetting(r *http.Request) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	sql := `DELETE FROM ApiSetting`
	var params []interface{}
	var whereSql = ""

	apiIds := strings.Split(r.FormValue("api_id"), ",")
	if len(apiIds) > 0 && apiIds[0] != "" {
		whereSql += " WHERE api_id IN (" + strings.Repeat("?,", len(apiIds)-1) + "?)"
		for _, id := range apiIds {
			params = append(params, id)
		}
	}

	apiSettingIds := strings.Split(r.FormValue("api_setting_id"), ",")
	if len(apiSettingIds) > 0 && apiSettingIds[0] != "" {
		if whereSql == "" {
			whereSql += " WHERE"
		} else {
			whereSql += " AND"
		}
		whereSql += " api_setting_id IN (" + strings.Repeat("?,", len(apiSettingIds)-1) + "?)"
		for _, id := range apiSettingIds {
			params = append(params, id)
		}
	}

	result, err := Db.Exec(sql+whereSql, params...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
