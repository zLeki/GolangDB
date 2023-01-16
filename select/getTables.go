package DropDown

import (
	"database/sql"
	"log"
	"strings"
)

func GetTableNames(db *sql.DB) []string {
	rows, err := db.Query(`
		SELECT
			table_schema || '.' || table_name
		FROM
			information_schema.tables
		WHERE
			table_type = 'BASE TABLE'
		AND
			table_schema NOT IN ('pg_catalog', 'information_schema');
	`)
	defer rows.Close()
	if err != nil {
		log.Fatalln("Error connecting to database", err)
	}
	var tableNames []string
	for rows.Next() {
		var table string
		rows.Scan(&table)
		table = strings.Title(strings.Replace(table, "public.", "", -1))
		tableNames = append(tableNames, table)
	}
	return tableNames
}
