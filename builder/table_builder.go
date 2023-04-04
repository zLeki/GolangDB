package builder

import (
	"GopherDB/types"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func BuildTable(db *sql.DB) ([]string, []string) {
	tableName := fmt.Sprintf("public.%s", strings.ToLower(types.TableSelection))
	const (
		maxRows = 300
		offset  = 0
	)
	query := fmt.Sprintf("SELECT * FROM %s LIMIT $1 OFFSET $2;", tableName)
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, query, maxRows, offset)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)
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
	var row = make([]string, 0)
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
