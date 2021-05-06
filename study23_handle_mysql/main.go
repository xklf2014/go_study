package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dataSource := "root:123456@tcp(localhost:3306)/gin_demo"
	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	//尝试连接
	if err := db.Ping(); err != nil {
		fmt.Println("database connect failed")
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("database connected")
}
