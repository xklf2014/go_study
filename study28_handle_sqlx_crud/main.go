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

func querySingleRow(id uint64) stu {
	sqlstr := "select * from student where id = ?"
	var s stu
	if err := db.Get(&s, sqlstr, id); err != nil {
		fmt.Printf("query error %v \n", err)
		return s
	}

	return s
}

func queryMultiRow() []stu {
	sqlstr := "select * from student"
	var stus []stu
	if err := db.Select(&stus, sqlstr); err != nil {
		fmt.Printf("query error %v \n", err)
		return stus
	}

	return stus
}

func updateRow(age uint32, id uint64) int64 {
	sqlstr := "update student set age = ? where id = ?"
	res, err := db.Exec(sqlstr, age, id)
	if err != nil {
		fmt.Printf("update failed err: %v\n", err)
		return 0
	}

	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsAffected failed err: %v\n", err)
		return 0
	}

	return affected
}

func insertSingleRow(name string, age uint32) int64 {
	sqlstr := "insert into student(name,age) values (?,?)"
	res, err := db.Exec(sqlstr, name, age)
	if err != nil {
		fmt.Printf("insert failed err: %v\n", err)
		return 0
	}

	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("insert rowsAffected failed err: %v\n", err)
		return 0
	}
	return affected
}

func deleteSingleRow(id uint64) int64 {
	sqlstr := "delete from student where id = ?"
	row, err := db.Exec(sqlstr, id)
	if err != nil {
		fmt.Printf("delete failed err: %v\n", err)
		return 0
	}

	affected, err := row.RowsAffected()
	if err != nil {
		fmt.Printf("delete rowsAffected failed err: %v\n", err)
		return 0
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

	r.GET("stus", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": queryMultiRow(),
		})
	})

	r.PUT("stu/:id/:age", func(c *gin.Context) {
		if err := c.ShouldBindUri(&s); err != nil {
			fmt.Printf("ShouldBindUri error %v", err)
			return
		}
		row := updateRow(s.AGE, s.ID)
		c.JSON(200, gin.H{
			"data": row,
		})
	})

	r.POST("stu/:name/:age", func(c *gin.Context) {
		if err := c.ShouldBindUri(&s); err != nil {
			fmt.Printf("ShouldBindUri error %v", err)
			return
		}
		row := insertSingleRow(s.Name, s.AGE)
		c.JSON(200, gin.H{
			"data": row,
		})
	})
	r.DELETE("stu/:id", func(c *gin.Context) {
		if err := c.ShouldBindUri(&s); err != nil {
			fmt.Printf("ShouldBindUri error %v", err)
			return
		}
		row := deleteSingleRow(s.ID)
		c.JSON(200, gin.H{
			"data": row,
		})
	})

	r.Run()
}
