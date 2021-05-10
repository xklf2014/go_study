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

type User struct {
	UserId uint   `gorm:"primarykey"`
	Name   string `gorm:"column:user_name;type:varchar(255);index:idx_user_name;unique"`
}

//自定义表名
/*func (User) TableName() string {
	return "my_user"
}
*/
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
	db.AutoMigrate(&User{})

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
