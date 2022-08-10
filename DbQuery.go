package utils

import (
	"database/sql"
	"errors"
)

// DbQuery executes a SQL query and returns array of key value maps
func DbQuery(db *sql.DB, query string) (result map[int]map[string]string, err error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()
	cols, _ := rows.Columns()               //columns
	vals := make([][]byte, len(cols))       //array
	scans := make([]interface{}, len(cols)) //arrat
	for k := range vals {
		scans[k] = &vals[k]
	}
	i := 0

	result = make(map[int]map[string]string)
	for rows.Next() {
		rows.Scan(scans...)            //scan
		row := make(map[string]string) //array
		for k, v := range vals {
			key := cols[k]
			// DEBUG: fmt.Println(string(v))
			row[key] = string(v)

		}
		result[i] = row
		i++
	}
	rows.Close()

	return result, nil
}
