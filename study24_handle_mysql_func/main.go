package main

import (
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlConnector struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Url      string `json:"url"`
	Params   string `json:"params"`
}

var db *sql.DB

func readProperties() string {
	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var conn MySqlConnector
	json.Unmarshal([]byte(byteValue), &conn)
	returnVal := conn.UserName + ":" + conn.Password + "@" + conn.Url + "?" + conn.Params
	return returnVal
}

func initializeDatabase() (err error) {
	dataSourceName := readProperties()
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	return nil

}

func main() {
	if err := initializeDatabase(); err != nil {
		panic(err)
	}

	fmt.Println("database succeed")

}
