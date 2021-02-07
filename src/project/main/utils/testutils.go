package person

import "fmt"

var Name string
var Age int
var Gender string

func init() {
	fmt.Println("testutils init 函数被调用")
	Name = "韩梅梅"
	Age = 16
	Gender = "女"
}
