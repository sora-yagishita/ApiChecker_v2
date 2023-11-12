package model

import (
	"database/sql"
	"net/http"
	"time"
)

type ApiResultModel interface {
	FetchApiResult(r *http.Request) ([]*ApiResult, error)
	AddApiResult(r *http.Request) (sql.Result, error)
	ChangeApiResult(r *http.Request) (sql.Result, error)
	DeleteApiResult(r *http.Request) (sql.Result, error)
}

type apiResultModel struct {
}

type ApiResult struct {
	ApiName     string    `json:"apiName"`
	ApiStatus   string    `json:"apiStatus"`
	ApiDateTime time.Time `json:"apiDateTime"`
}

func CreateApiResultModel() ApiResultModel {
	return &apiResultModel{}
}

func (tm *apiResultModel) FetchApiResult(r *http.Request) ([]*ApiResult, error) {
  err := r.ParseForm()

	if err != nil {
		return nil, nil
	}

	sql := `SELECT apiName, apiStatus, apiDateTime FROM apiResultHistory WHERE apiName = ?`

	rows, err := Db.Query(sql, r.FormValue("apiName"))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var apiResults []*ApiResult

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

		apiResults = append(apiResults, &ApiResult{
			ApiName: apiName,
			ApiStatus: apiStatus,
			ApiDateTime: apiDateTime,
		})
	}

	return apiResults, nil
}

func (tm *apiResultModel) AddApiResult(r *http.Request) (sql.Result, error) {
	// t := time.Now()
	// entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	// id := ulid.MustNew(ulid.Timestamp(t), entropy)

	// req := entities.Todo{
	// 	ApiName: r.ApiName,
	// 	ApiStatus: r.ApiStatus,
	// 	ApiDateTime: r.ApiDateTime,
	// }

	sql := `INSERT INTO todos(id, name, status) VALUES("aaa", "bbb", "vvv")`

	result, err := Db.Exec(sql)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *apiResultModel) ChangeApiResult(r *http.Request) (sql.Result, error) {
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

func (tm *apiResultModel) DeleteApiResult(r *http.Request) (sql.Result, error) {
	sql := `DELETE FROM todos WHERE id = "aaa"`

	result, err := Db.Exec(sql)

	if err != nil {
		return result, err
	}

	return result, nil
}