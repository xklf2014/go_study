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
	if err := db.QueryRow(sqlStr, id).Scan(&s.Name, &s.AGE, &s.ID); err != nil {
		log.Printf("scan failed err: %v \n", err)
		return s
	}

	return s
}

func queryMultiRows() []stu {
	sqlStr := "select * from student"
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Printf("scan failed err: %v \n", err)
		return nil
	}

	defer rows.Close()
	stus := make([]stu, 0)

	for rows.Next() {
		var s stu
		err := rows.Scan(&s.Name, &s.AGE, &s.ID)
		if err != nil {
			log.Println(err)
			return nil
		}
		stus = append(stus, s)
	}
	return stus
}

func updateRow(id uint64, name string) bool {
	sqlStr := "update student set name = ? where id = ?"
	res, err := db.Exec(sqlStr, name, id)
	if err != nil {
		log.Println(err)
		return false
	}

	rs, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return false
	}

	if rs != 1 {
		log.Println("更新失败")
		return false
	}
	return true
}

func deleteRow(id uint64) bool {
	sqlStr := "delete from student  where id = ?"
	res, err := db.Exec(sqlStr, id)
	if err != nil {
		log.Println(err)
		return false
	}

	rs, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return false
	}

	if rs != 1 {
		return false
	}
	return true

}

func insertRow(name string, age uint32) bool {
	sqlStr := "insert into student(name,age) values (?,?)"
	res, err := db.Exec(sqlStr, name, age)
	if err != nil {
		log.Println(err)
		return false
	}
	rs, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return false
	}

	if rs != 1 {
		return false
	}

	return true
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

	r.GET("stus", func(c *gin.Context) {
		s := queryMultiRows()
		c.JSON(200, gin.H{
			"data": s,
		})
	})

	r.GET("stu/:id/:name", func(c *gin.Context) {
		if err := c.ShouldBindUri(&s); err != nil {
			c.JSON(200, gin.H{
				"Code": 200,
				"Msg":  err.Error(),
			})
			return
		}

		s := updateRow(s.ID, s.Name)
		c.JSON(200, gin.H{
			"data": s,
		})
	})

	r.DELETE("stu/:id", func(c *gin.Context) {
		if err := c.ShouldBindUri(&s); err != nil {
			c.JSON(200, gin.H{
				"Code": 200,
				"Msg":  err.Error(),
			})
			return
		}

		if deleteRow(s.ID) {
			c.JSON(200, gin.H{
				"msg": "delete success",
			})
		}

	})

	r.POST("stu/:name/:age", func(c *gin.Context) {
		if err := c.ShouldBindUri(&s); err != nil {
			c.JSON(200, gin.H{
				"Code": 200,
				"Msg":  err.Error(),
			})
			return
		}

		if insertRow(s.Name, s.AGE) {
			c.JSON(200, gin.H{
				"msg": "insert success",
			})
		}

	})

	r.Run()

}
