package model

import (
	"src/model/entities"
	"database/sql"
	"net/http"
	"time"
)

type ApiResultModel interface {
	FetchApiResult(r *http.Request) ([]*entities.ApiResult, error)
	AddApiResult(r entities.ApiResult) (sql.Result, error)
	ChangeApiResult(r entities.ApiResult) (sql.Result, error)
	DeleteApiResult(r entities.ApiResult) (sql.Result, error)
}

type apiResultModel struct {
}

func CreateApiResultModel() ApiResultModel {
	return &apiResultModel{}
}

func (am *apiResultModel) FetchApiResult(r *http.Request) ([]*entities.ApiResult, error) {
  err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	sql := `SELECT apiName, apiStatus, apiDateTime FROM apiResultHistory WHERE apiName = ?`

	rows, err := Db.Query(sql, r.FormValue("apiName"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apiResults []*entities.ApiResult

	for rows.Next() {
		var apiName string
		var apiStatus string
		var apiDateTime time.Time
		// Time型のみキャストを行う必要がある
		var apiDateTimeRaw []uint8

		if err := rows.Scan(&apiName, &apiStatus, &apiDateTimeRaw); err != nil {
			return nil, err
		}

		apiDateTime, err := time.Parse("2006-01-02 15:04:05", string(apiDateTimeRaw))
    if err != nil {
      return nil, err
    }

		apiResults = append(apiResults, &entities.ApiResult{
			ApiName: apiName,
			ApiStatus: apiStatus,
			ApiDateTime: apiDateTime,
		})
	}

	return apiResults, nil
}

func (tm *apiResultModel) AddApiResult(r entities.ApiResult) (sql.Result, error) {
	req := entities.ApiResult{
		ApiName: r.ApiName,
		ApiStatus: r.ApiStatus,
		ApiDateTime: r.ApiDateTime,
	}

	sql := `INSERT INTO apiResultHistory(apiName, apiStatus, apiDateTime) VALUES(?, ?, ?)`

	result, err := Db.Exec(sql, req.ApiName, req.ApiStatus, req.ApiDateTime)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *apiResultModel) ChangeApiResult(r entities.ApiResult) (sql.Result, error) {
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

func (tm *apiResultModel) DeleteApiResult(r entities.ApiResult) (sql.Result, error) {
	sql := `DELETE FROM todos WHERE id = "aaa"`

	result, err := Db.Exec(sql)

	if err != nil {
		return result, err
	}

	return result, nil
}