package model

import (
	"database/sql"
	"net/http"
	"src/model/entities"
	"strings"
)

type ApiHeaderSettingModel interface {
	FetchApiHeaderSetting(r *http.Request) ([]*entities.ApiHeaderSetting, error)
	AddApiHeaderSetting(r entities.ApiHeaderSetting) (sql.Result, error)
	UpdateApiHeaderSetting(r *http.Request, e entities.ApiHeaderSetting) (sql.Result, error)
	DeleteApiHeaderSetting(r *http.Request) (sql.Result, error)
}

type apiHeaderSettingModel struct {
}

func CreateApiHeaderSettingModel() ApiHeaderSettingModel {
	return &apiHeaderSettingModel{}
}

func (am *apiHeaderSettingModel) FetchApiHeaderSetting(r *http.Request) ([]*entities.ApiHeaderSetting, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	sql := `SELECT api_setting_id, api_header_setting_no, api_header_key, api_header_value FROM ApiHeaderSetting`
	var params []interface{}
	var whereSql = ""

	apiSettingIds := strings.Split(r.FormValue("api_setting_id"), ",")
	if len(apiSettingIds) > 0 && apiSettingIds[0] != "" {
		whereSql += " WHERE api_setting_id IN (" + strings.Repeat("?,", len(apiSettingIds)-1) + "?)"
		for _, id := range apiSettingIds {
			params = append(params, id)
		}
	}

	apiHeaderSettingNos := strings.Split(r.FormValue("api_header_setting_no"), ",")
	if len(apiHeaderSettingNos) > 0 && apiHeaderSettingNos[0] != "" {
		if whereSql == "" {
			whereSql += " WHERE"
		} else {
			whereSql += " AND"
		}
		whereSql += " api_header_setting_no IN (" + strings.Repeat("?,", len(apiHeaderSettingNos)-1) + "?)"
		for _, id := range apiHeaderSettingNos {
			params = append(params, id)
		}
	}

	rows, err := Db.Query(sql+whereSql, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var apiHeaderSetting []*entities.ApiHeaderSetting

	for rows.Next() {
		var apiSettingId int
		var apiHeaderSettingNo int
		var apiHeaderKey string
		var apiHeaderValue string

		if err := rows.Scan(&apiSettingId, &apiHeaderSettingNo, &apiHeaderKey, &apiHeaderValue); err != nil {
			return nil, err
		}

		apiHeaderSetting = append(apiHeaderSetting, &entities.ApiHeaderSetting{
			ApiSettingId:       apiSettingId,
			ApiHeaderSettingNo: apiHeaderSettingNo,
			ApiHeaderKey:       apiHeaderKey,
			ApiHeaderValue:     apiHeaderValue,
		})
	}

	return apiHeaderSetting, nil
}

func (tm *apiHeaderSettingModel) AddApiHeaderSetting(e entities.ApiHeaderSetting) (sql.Result, error) {
	req := entities.ApiHeaderSetting{
		ApiSettingId:       e.ApiSettingId,
		ApiHeaderSettingNo: e.ApiHeaderSettingNo,
		ApiHeaderKey:       e.ApiHeaderKey,
		ApiHeaderValue:     e.ApiHeaderValue,
	}

	sql := `INSERT INTO ApiHeaderSetting(api_setting_id, api_header_setting_no, api_header_key, api_header_value) VALUES(?, ?, ?, ?)`

	result, err := Db.Exec(sql, req.ApiSettingId, req.ApiHeaderSettingNo, req.ApiHeaderKey, req.ApiHeaderValue)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *apiHeaderSettingModel) UpdateApiHeaderSetting(r *http.Request, e entities.ApiHeaderSetting) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	req := entities.ApiHeaderSetting{
		ApiHeaderKey:   e.ApiHeaderKey,
		ApiHeaderValue: e.ApiHeaderValue,
	}

	sql := `UPDATE ApiHeaderSetting SET api_header_key = ?, api_header_value = ?`
	params := []interface{}{req.ApiHeaderKey, req.ApiHeaderValue}
	var whereSql = ""

	apiSettingIds := strings.Split(r.FormValue("api_setting_id"), ",")
	if len(apiSettingIds) > 0 && apiSettingIds[0] != "" {
		whereSql += " WHERE api_setting_id IN (" + strings.Repeat("?,", len(apiSettingIds)-1) + "?)"
		for _, id := range apiSettingIds {
			params = append(params, id)
		}
	}

	apiHeaderSettingNos := strings.Split(r.FormValue("api_header_setting_no"), ",")
	if len(apiHeaderSettingNos) > 0 && apiHeaderSettingNos[0] != "" {
		if whereSql == "" {
			whereSql += " WHERE"
		} else {
			whereSql += " AND"
		}
		whereSql += " api_header_setting_no IN (" + strings.Repeat("?,", len(apiHeaderSettingNos)-1) + "?)"
		for _, id := range apiHeaderSettingNos {
			params = append(params, id)
		}
	}

	result, err := Db.Exec(sql+whereSql, params...)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *apiHeaderSettingModel) DeleteApiHeaderSetting(r *http.Request) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	sql := `DELETE FROM ApiHeaderSetting`
	var params []interface{}
	var whereSql = ""

	apiSettingIds := strings.Split(r.FormValue("api_setting_id"), ",")
	if len(apiSettingIds) > 0 && apiSettingIds[0] != "" {
		whereSql += " WHERE api_setting_id IN (" + strings.Repeat("?,", len(apiSettingIds)-1) + "?)"
		for _, id := range apiSettingIds {
			params = append(params, id)
		}
	}

	apiHeaderSettingNos := strings.Split(r.FormValue("api_header_setting_no"), ",")
	if len(apiHeaderSettingNos) > 0 && apiHeaderSettingNos[0] != "" {
		if whereSql == "" {
			whereSql += " WHERE"
		} else {
			whereSql += " AND"
		}
		whereSql += " api_header_setting_no IN (" + strings.Repeat("?,", len(apiHeaderSettingNos)-1) + "?)"
		for _, id := range apiHeaderSettingNos {
			params = append(params, id)
		}
	}

	result, err := Db.Exec(sql+whereSql, params...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
