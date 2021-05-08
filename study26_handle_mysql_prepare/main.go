package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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
	db, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	return nil

}

type stu struct {
	ID   uint64 `uri:"id"`
	Name string `uri:"name"`
	AGE  uint32 `uri:"age"`
}

func querySingleRow(id uint64) stu {
	var s stu
	sqlStr := "select * from student where id = ?"
	stmt, err := db.Prepare(sqlStr)

	if err != nil {
		log.Println(err)
		return s
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(&s.Name, &s.AGE, &s.ID)
	if err != nil {
		log.Println(err)
		return s
	}

	return s
}

func main() {
	if err := initializeDatabase(); err != nil {
		panic(err)
	}

	fmt.Println("database succeed")
	var s stu
	r := gin.Default()
	r.GET("stu/:id", func(c *gin.Context) {
		if err := c.ShouldBindUri(&s); err != nil {
			c.JSON(200, gin.H{
				"Code": 200,
				"Msg":  err.Error(),
			})
			return
		}

		s = querySingleRow(s.ID)
		c.JSON(200, gin.H{
			"data": s,
		})
	})

	r.Run()

}
