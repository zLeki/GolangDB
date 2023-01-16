package types

import "database/sql"

var (
	BaseStruct     = Base{}
	TableSelection string
	TablePrimary   string
	Unit           string
	Db             *sql.DB
	Name           string
	Page           = 0
)

type Base struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DataBaseName string
}
