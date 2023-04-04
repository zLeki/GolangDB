package types

import "database/sql"

var (
	BaseStruct     = Base{}
	TableSelection string
	TablePrimary   string
	Db             *sql.DB
	Page           = 0
)

type Base struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DataBaseName string
}
