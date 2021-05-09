package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySqlConnector struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Url      string `json:"url"`
	Params   string `json:"params"`
}

var db *sqlx.DB

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
	db, err = sqlx.Connect("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		return
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return nil

}

type stu struct {
	ID   uint64 `uri:"id"`
	Name string `uri:"name"`
	AGE  uint32 `uri:"age"`
}

func selectNameQuery(age uint32) []stu {
	tmpStu := stu{
		AGE: age,
	}
	sqlStr := "select * from student where age = :age"
	rows, err := db.NamedQuery(sqlStr, tmpStu)
	if err != nil {
		fmt.Printf("NamedQuery failed by %v\n", err)
	}
	defer rows.Close()
	stus := make([]stu, 0)

	for rows.Next() {
		var s stu
		err = rows.StructScan(&s)
		if err := rows.StructScan(&s); err != nil {
			fmt.Printf("StructScan failed by %v\n", err)
			continue
		}
		stus = append(stus, s)
	}

	return stus
}

func batchInsert() int64 {
	stus := []stu{
		{Name: "zhangyi", AGE: 18},
		{Name: "zhanger", AGE: 18},
		{Name: "zhangsi", AGE: 18},
	}

	sqlStr := "insert into student (name,age) values (:name,:age)"
	rows, err := db.NamedExec(sqlStr, stus)
	if err != nil {
		fmt.Printf("NamedExec failed by %v\n", err)
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		fmt.Printf("RowsAffected failed by %v\n", err)
	}

	return affected

}

func main() {
	if err := initializeDatabase(); err != nil {
		panic(err)
	}

	fmt.Println("connect succeed")

	r := gin.Default()
	var s stu
	r.GET("stus/:age", func(c *gin.Context) {

		if err := c.ShouldBindUri(&s); err != nil {
			fmt.Printf("ShouldBindUri failed by %v\n", err)
			return
		}

		c.JSON(200, gin.H{
			"data": selectNameQuery(s.AGE),
		})
	})

	r.POST("stus", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"rows": batchInsert(),
		})
	})

	r.Run()
}
