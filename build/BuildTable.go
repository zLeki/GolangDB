package build

import (
	"GopherDB/types"
	"database/sql"
	"log"
	"strings"
)

func BuildTable(db *sql.DB) ([]string, []string) {
	rows, err := db.Query(`SELECT * FROM "public"."` + strings.ToLower(types.TableSelection) + `" LIMIT 300 OFFSET 0;`)
	defer rows.Close()
	if err != nil {
		log.Fatalln("Error connecting to database", err)
	}
	columns, err := rows.Columns()
	if err != nil {
		log.Fatalln("Error connecting to database", err)

	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]

	}
	var row = []string{}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatalln("Error connecting to database", err)
		}
		var a string
		var value string
		for _, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			a += value + "|"
		}
		row = append(row, a)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln("Error connecting to database", err)
	}
	return columns, row
}
