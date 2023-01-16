package main

import (
	"GopherDB/core"
	"GopherDB/input"
	"GopherDB/select"
	"GopherDB/types"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	"strings"
)

var page = 0

type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func init() {
	// SETUP
	var config *Config
	f, err := os.OpenFile("./helpers/config.json", os.O_RDONLY, 0644)
	if err != nil {
		os.Create("./helpers/config.json")
	}
	_ = json.NewDecoder(f).Decode(&config)
	if config.Host == "" || config.Port == 0 || config.Username == "" || config.Password == "" || config.Database == "" {
		types.Page--
		var items = []string{"PostgresDB", "MongoDB", "MySQL", "SQLite"}
		DropDown.DropDown(items, "Setup")
		var questions = []string{"Host", "Port", "Username", "Password", "DatabaseName"}
		input.Input(ParseQuestions(questions))
		config.Host = types.BaseStruct.Host
		config.Port, _ = strconv.Atoi(types.BaseStruct.Port)
		config.Username = types.BaseStruct.Username
		config.Password = types.BaseStruct.Password
		config.Database = types.BaseStruct.DataBaseName
		fmt.Println(config.Host, types.BaseStruct.Host)
		f, _ := os.OpenFile("./helpers/config.json", os.O_WRONLY, 0644)
		var b []byte
		b, err = json.Marshal(config)
		if err != nil {
			panic(err)
		}
		_, err = f.Write(b)
		if err != nil {
			panic(err)
		}
		fmt.Println("Setup Complete\nYou can now use GopherDB")
		select {}

	} else {
		types.BaseStruct = types.Base{Host: config.Host, Port: strconv.Itoa(config.Port), Username: config.Username, Password: config.Password, DataBaseName: config.Database}
	}
}
func main() {
	_, _ = ParseDB()

	core.PrevPage()

}

type info struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}

func ParseDB() (*sql.DB, error) {
	var port, _ = strconv.Atoi(types.BaseStruct.Port)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		types.BaseStruct.Host, port, types.BaseStruct.Username, types.BaseStruct.Password, types.BaseStruct.DataBaseName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	types.Db = db
	return db, nil
}
func ParseQuestions(items []string) []input.Questions {
	var questions []input.Questions
	for i := range items {
		if strings.Contains(items[i], "Password") {
			questions = append(questions, input.Questions{Placeholder: items[i], CharLimit: 128, Password: true})
		} else {
			questions = append(questions, input.Questions{Placeholder: items[i], CharLimit: 128, Password: false})
		}

	}
	return questions
}
