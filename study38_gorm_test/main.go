package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm/schema"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ClassName string
	Students  []Student
}

type Student struct {
	gorm.Model
	ClassId     uint
	IDCard      IdCard
	Teachers    []Teacher `gorm:"many2many:Student_Teacher;"`
	StudentName string
}

type IdCard struct {
	gorm.Model
	Num       int
	Addr      string
	StudentId uint
}
type Teacher struct {
	gorm.Model
	TeacherName string
	Students    []Student `gorm:"many2many:Student_Teacher;"`
}

func main() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(readProperties()), &gorm.Config{
		Logger: newLogger,
		//
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, //设置表名单复数,true为单数
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	//db.AutoMigrate(&Class{}, &Teacher{}, &IdCard{}, &Student{})

	r := gin.Default()

	r.POST("class", func(c *gin.Context) {
		db.Create(&Class{ClassName: "GO课程"})
		c.JSON(200, gin.H{
			"msg": "success",
		})
	})

	r.POST("student", func(c *gin.Context) {
		var s Student
		c.ShouldBindJSON(&s)
		db.Create(&s)
		c.JSON(200, gin.H{
			"msg": "success",
		})
	})

	r.Run()
}

type MySqlConnector struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Url      string `json:"url"`
	Params   string `json:"params"`
}

var Db *sqlx.DB

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
