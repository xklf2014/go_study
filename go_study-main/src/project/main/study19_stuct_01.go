package main

import "fmt"

type teacher struct {
	Name       string
	Age        int
	Experience int
	Edu        string
}

func main() {

	t := teacher{
		Name:       "张三",
		Age:        27,
		Experience: 5,
		Edu:        "哈佛",
	}

	fmt.Println(t)

	var t1 teacher
	t1.Name = "李四"
	t1.Age = 25
	t1.Edu = "剑桥"
	t1.Experience = 3

	fmt.Println(t1)

	var t2 *teacher = new(teacher)
	(*t2).Name = "王五"
	(*t2).Age = 18
	(*t2).Edu = "华盛顿"
	t2.Experience = 1 // t2.Experience 编译时会转换为(*t2).Experience

	fmt.Println(t2)

	var t3 *teacher = &teacher{}
	t3.Name = "赵六"
	(*t3).Age = 40
	t3.Experience = 20
	t3.Edu = "堪萨斯"

	fmt.Println(t3)
}
