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
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	if err != nil {
		log.Fatalln("Error connecting to database", err)
	}
	var tableNames []string
	for rows.Next() {
		var table string
		_ = rows.Scan(&table)
		// Deprecated: The rule Title uses for word boundaries does not handle Unicode punctuation properly. Use golang.org/x/text/cases instead.
		table = strings.Title(strings.Replace(table, "public.", "", -1))
		tableNames = append(tableNames, table)
	}
	return tableNames
}
