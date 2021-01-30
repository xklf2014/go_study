package main

import (
	"fmt"
	"project/main/dbutils"
)

func main() {
	fmt.Println("main 函数")
	dbutils.GetConn()
}
