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
	UpdateApi(r entities.Api) (sql.Result, error)
	DeleteApi(r entities.Api) (sql.Result, error)
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

	apiIds := strings.Split(r.FormValue("apiId"), ",")
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

func (tm *apiModel) AddApi(r entities.Api) (sql.Result, error) {
	req := entities.Api{
		ApiName:        r.ApiName,
		ApiDescription: r.ApiDescription,
	}

	sql := `INSERT INTO Api(api_name, api_description) VALUES(?, ?)`

	result, err := Db.Exec(sql, req.ApiName, req.ApiDescription)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *apiModel) UpdateApi(r entities.Api) (sql.Result, error) {
	sql := `UPDATE todos SET status = "aaa" WHERE id = "aaa"`

	// afterStatus := "作業中"
	// if r.ApiStatus == "作業中" {
	// 	afterStatus = "完了"
	// }

	result, err := Db.Exec(sql)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *apiModel) DeleteApi(r entities.Api) (sql.Result, error) {
	sql := `DELETE FROM todos WHERE id = "aaa"`

	result, err := Db.Exec(sql)

	if err != nil {
		return result, err
	}

	return result, nil
}
