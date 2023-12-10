package model

import (
	"database/sql"
	"net/http"
	"src/model/entities"
	"strings"
)

type ApiParamSettingModel interface {
	FetchApiParamSetting(r *http.Request) ([]*entities.ApiParamSetting, error)
	AddApiParamSetting(r entities.ApiParamSetting) (sql.Result, error)
	UpdateApiParamSetting(r *http.Request, e entities.ApiParamSetting) (sql.Result, error)
	DeleteApiParamSetting(r *http.Request) (sql.Result, error)
}

type apiParamSettingModel struct {
}

func CreateApiParamSettingModel() ApiParamSettingModel {
	return &apiParamSettingModel{}
}

func (am *apiParamSettingModel) FetchApiParamSetting(r *http.Request) ([]*entities.ApiParamSetting, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	sql := `SELECT api_setting_id, api_param_setting_no, api_param_key, api_param_value FROM ApiParamSetting`
	var params []interface{}
	var whereSql = ""

	apiSettingIds := strings.Split(r.FormValue("api_setting_id"), ",")
	if len(apiSettingIds) > 0 && apiSettingIds[0] != "" {
		whereSql += " WHERE api_setting_id IN (" + strings.Repeat("?,", len(apiSettingIds)-1) + "?)"
		for _, id := range apiSettingIds {
			params = append(params, id)
		}
	}

	apiParamSettingNos := strings.Split(r.FormValue("api_param_setting_no"), ",")
	if len(apiParamSettingNos) > 0 && apiParamSettingNos[0] != "" {
		if whereSql == "" {
			whereSql += " WHERE"
		} else {
			whereSql += " AND"
		}
		whereSql += " api_param_setting_no IN (" + strings.Repeat("?,", len(apiParamSettingNos)-1) + "?)"
		for _, id := range apiParamSettingNos {
			params = append(params, id)
		}
	}

	rows, err := Db.Query(sql+whereSql, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var apiParamSetting []*entities.ApiParamSetting

	for rows.Next() {
		var apiSettingId int
		var apiParamSettingNo int
		var apiParamKey string
		var apiParamValue string

		if err := rows.Scan(&apiSettingId, &apiParamSettingNo, &apiParamKey, &apiParamValue); err != nil {
			return nil, err
		}

		apiParamSetting = append(apiParamSetting, &entities.ApiParamSetting{
			ApiSettingId:      apiSettingId,
			ApiParamSettingNo: apiParamSettingNo,
			ApiParamKey:       apiParamKey,
			ApiParamValue:     apiParamValue,
		})
	}

	return apiParamSetting, nil
}

func (tm *apiParamSettingModel) AddApiParamSetting(e entities.ApiParamSetting) (sql.Result, error) {
	req := entities.ApiParamSetting{
		ApiSettingId:      e.ApiSettingId,
		ApiParamSettingNo: e.ApiParamSettingNo,
		ApiParamKey:       e.ApiParamKey,
		ApiParamValue:     e.ApiParamValue,
	}

	sql := `INSERT INTO ApiParamSetting(api_setting_id, api_param_setting_no, api_param_key, api_param_value) VALUES(?, ?, ?, ?)`

	result, err := Db.Exec(sql, req.ApiSettingId, req.ApiParamSettingNo, req.ApiParamKey, req.ApiParamValue)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *apiParamSettingModel) UpdateApiParamSetting(r *http.Request, e entities.ApiParamSetting) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	req := entities.ApiParamSetting{
		ApiParamKey:   e.ApiParamKey,
		ApiParamValue: e.ApiParamValue,
	}

	sql := `UPDATE ApiParamSetting SET api_param_key = ?, api_param_value = ?`
	params := []interface{}{req.ApiParamKey, req.ApiParamValue}
	var whereSql = ""

	apiSettingIds := strings.Split(r.FormValue("api_setting_id"), ",")
	if len(apiSettingIds) > 0 && apiSettingIds[0] != "" {
		whereSql += " WHERE api_setting_id IN (" + strings.Repeat("?,", len(apiSettingIds)-1) + "?)"
		for _, id := range apiSettingIds {
			params = append(params, id)
		}
	}

	apiParamSettingNos := strings.Split(r.FormValue("api_param_setting_no"), ",")
	if len(apiParamSettingNos) > 0 && apiParamSettingNos[0] != "" {
		if whereSql == "" {
			whereSql += " WHERE"
		} else {
			whereSql += " AND"
		}
		whereSql += " api_param_setting_no IN (" + strings.Repeat("?,", len(apiParamSettingNos)-1) + "?)"
		for _, id := range apiParamSettingNos {
			params = append(params, id)
		}
	}

	result, err := Db.Exec(sql+whereSql, params...)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *apiParamSettingModel) DeleteApiParamSetting(r *http.Request) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	sql := `DELETE FROM ApiParamSetting`
	var params []interface{}
	var whereSql = ""

	apiSettingIds := strings.Split(r.FormValue("api_setting_id"), ",")
	if len(apiSettingIds) > 0 && apiSettingIds[0] != "" {
		whereSql += " WHERE api_setting_id IN (" + strings.Repeat("?,", len(apiSettingIds)-1) + "?)"
		for _, id := range apiSettingIds {
			params = append(params, id)
		}
	}

	apiParamSettingNos := strings.Split(r.FormValue("api_param_setting_no"), ",")
	if len(apiParamSettingNos) > 0 && apiParamSettingNos[0] != "" {
		if whereSql == "" {
			whereSql += " WHERE"
		} else {
			whereSql += " AND"
		}
		whereSql += " api_param_setting_no IN (" + strings.Repeat("?,", len(apiParamSettingNos)-1) + "?)"
		for _, id := range apiParamSettingNos {
			params = append(params, id)
		}
	}

	result, err := Db.Exec(sql+whereSql, params...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
