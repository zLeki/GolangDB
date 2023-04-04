package types

import (
	"database/sql"
	"strconv"
)

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

func NewConnection(host string, port int, username string, password string, databaseName string) Base {
	portString := strconv.Itoa(port)
	return Base{
		Host:         host,
		Port:         portString,
		Username:     username,
		Password:     password,
		DataBaseName: databaseName,
	}
}
