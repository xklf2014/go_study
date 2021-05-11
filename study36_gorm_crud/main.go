package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gorm.io/gorm/schema"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string
	Age    uint
	Gender bool
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
			TablePrefix:   "api_", //设置表名前缀
			SingularTable: true,   //设置表名单复数,true为单数
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	//db.AutoMigrate(&Student{})

	/*	s := Student{
			Name:   "story3",
			Age:    18,
			Gender: false,
		}

		db.Create(&s)*/

	//var stus = []Student{{Name: "zhangsan"}, {Name: "lisi"}, {Name: "wangwu"}}
	//db.Create(&stus)

	//var stus1 = []Student{{Name: "zhangsan"}, {Name: "lisi"}, {Name: "wangwu"}}
	//db.CreateInBatches(&stus1, 5)

	/*	db.Model(&Student{}).Create([]map[string]interface{}{
		{"Name": "m_lisan", "Age": 18},
		{"Name": "m_lisi", "Age": 28},
	})*/

	/*	s := Student{
		Model: gorm.Model{
			ID: 1,
		},
	}*/

	//db.First(&s)
	//fmt.Println(s.ID, s.Name, s.Age)

	/*	s := Student{}
		res := db.Last(&s)

		fmt.Println(s.ID, s.Name, s.Age)*/

	/*s := Student{
		Model: gorm.Model{
			ID: 111,
		},
	}
	res := db.Last(&s)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		fmt.Println("no record was founded")
		return
	}
	fmt.Println(s.ID, s.Name, s.Age)*/

	/*	var stus []Student
		res := db.Find(&stus)
		fmt.Println(res.RowsAffected)*/

	//条件查询
	/*	var s Student
		db.Where("name = ? and age = ?", "story1", 18).First(&s)
		fmt.Println(s.ID, s.Name, s.Age)

		var s1 Student
		db.Where(&Student{Name: "story1", Age: 18}).First(&s1)
		fmt.Println(s1.ID, s1.Name, s1.Age)*/
	/*	var stus []Student
		db.Where([]int64{1, 3, 6}).Find(&stus)
		for _, v := range stus {
			fmt.Println(v.Name)

		}*/
	/*	var stus []Student
		db.Where("name = ?", "story1").Or("age = ?", 18).Find(&stus)
		for _, v := range stus {
			fmt.Println(v.Name)
		}*/

	//更新
	/*	var s Student
		db.First(&s)

		s.Name = "new"
		db.Save(&s)*/

	/*	s := Student{
			Model: gorm.Model{
				ID: 2,
			},
		}
		db.Model(&s).Update("name", "haha")*/

	//逻辑删除
	/*	s := Student{
			Model: gorm.Model{
				ID: 1,
			},
		}

		db.Delete(&s)*/
	var stus []Student
	db.Unscoped().Where("age=?", 18).Find(&stus)
	for _, v := range stus {
		fmt.Println(v.Name)
	}
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
