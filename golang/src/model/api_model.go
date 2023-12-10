package model

import (
	"database/sql"
	"net/http"
	"src/model/entities"
	"strings"
)

type ApiModel interface {
	FetchApi(r *http.Request) ([]*entities.Api, error)
	AddApi(r entities.Api) (sql.Result, error)
	UpdateApi(r *http.Request, e entities.Api) (sql.Result, error)
	DeleteApi(r *http.Request) (sql.Result, error)
}

type apiModel struct {
}

func CreateApiModel() ApiModel {
	return &apiModel{}
}

func (am *apiModel) FetchApi(r *http.Request) ([]*entities.Api, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	sql := `SELECT api_id, api_name, api_description FROM Api`
	var params []interface{}

	apiIds := strings.Split(r.FormValue("api_id"), ",")
	if len(apiIds) > 0 && apiIds[0] != "" {
		sql += " WHERE api_id IN (" + strings.Repeat("?,", len(apiIds)-1) + "?)"
		for _, id := range apiIds {
			params = append(params, id)
		}
	}

	rows, err := Db.Query(sql, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var api []*entities.Api

	for rows.Next() {
		var apiId int
		var apiName string
		var apiDescription string

		if err := rows.Scan(&apiId, &apiName, &apiDescription); err != nil {
			return nil, err
		}

		api = append(api, &entities.Api{
			ApiId:          apiId,
			ApiName:        apiName,
			ApiDescription: apiDescription,
		})
	}

	return api, nil
}

func (tm *apiModel) AddApi(e entities.Api) (sql.Result, error) {
	req := entities.Api{
		ApiName:        e.ApiName,
		ApiDescription: e.ApiDescription,
	}

	sql := `INSERT INTO Api(api_name, api_description) VALUES(?, ?)`

	result, err := Db.Exec(sql, req.ApiName, req.ApiDescription)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *apiModel) UpdateApi(r *http.Request, e entities.Api) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	req := entities.Api{
		ApiName:        e.ApiName,
		ApiDescription: e.ApiDescription,
	}

	sql := `UPDATE Api SET api_name = ?, api_description = ?`

	params := []interface{}{req.ApiName, req.ApiDescription}

	apiIds := strings.Split(r.FormValue("api_id"), ",")
	if len(apiIds) > 0 && apiIds[0] != "" {
		sql += " WHERE api_id IN (" + strings.Repeat("?,", len(apiIds)-1) + "?)"
		for _, id := range apiIds {
			params = append(params, id)
		}
	}

	result, err := Db.Exec(sql, params...)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *apiModel) DeleteApi(r *http.Request) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	sql := `DELETE FROM Api`
	var params []interface{}

	apiIds := strings.Split(r.FormValue("api_id"), ",")
	if len(apiIds) > 0 && apiIds[0] != "" {
		sql += " WHERE api_id IN (" + strings.Repeat("?,", len(apiIds)-1) + "?)"
		for _, id := range apiIds {
			params = append(params, id)
		}
	}

	result, err := Db.Exec(sql, params...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
