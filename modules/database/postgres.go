package database

import (
	"database/sql"
	"fmt"

	"fast-api.io/models"
	"github.com/jmoiron/sqlx"
)

func PostgresExec(dsn string, query string) (sql.Result, error) {

	db, dbErr := sqlx.Connect("postgres", dsn)
	if dbErr != nil {
		return nil, dbErr
	}
	result, err := db.Exec(query)
	if err != nil {
		db.Close()
		return nil, err
	}

	db.Close()

	return result, nil
}

func PostgresQueryRow(dsn string, query string) (*int, error) {

	db, dbErr := sqlx.Connect("postgres", dsn)
	if dbErr != nil {
		return nil, dbErr
	}
	lastInsertId := 0
	err := db.QueryRow(query).Scan(&lastInsertId)
	if err != nil {
		db.Close()
		return nil, err
	}

	db.Close()

	return &lastInsertId, nil
}

func PostgresQuery(dsn string, query string) ([]map[string]interface{}, error) {
	db, dbErr := sqlx.Connect("postgres", dsn)
	if dbErr != nil {
		return nil, dbErr
	}
	rows, _ := db.Query(query)

	columns, _ := rows.Columns()
	var allMaps []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))
		for i, _ := range values {
			pointers[i] = &values[i]
		}
		rows.Scan(pointers...)
		resultMap := make(map[string]interface{})
		for i, val := range values {
			resultMap[columns[i]] = val
		}
		allMaps = append(allMaps, resultMap)
	}
	db.Close()
	return allMaps, nil
}

func GetDNSByWorkspace(workspace *models.Workspace) string {
	return fmt.Sprint(
		"host", "=", "localhost",
		" ",
		"user", "=", "user",
		" ",
		"password", "=", "abc@123",
		" ",
		"dbname", "=", "xano_clone",
		" ",
		"port", "=", "5432",
		" ",
		"sslmode", "=", "disable",
	)
}
