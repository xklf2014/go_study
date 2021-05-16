package main

import "fmt"

func main() {

	//方式一
	var a map[int]string
	a = make(map[int]string, 10)
	a[20090101] = "zhangsan"
	a[20090102] = "lisi"
	a[20090103] = "wangwu"

	fmt.Println(a)

	fmt.Println("----------------")
	//方式二
	b := make(map[int]string)
	b[20090102] = "lisi"
	b[20090103] = "wangwu"

	fmt.Println(b)
	fmt.Println("----------------")
	//方式三
	c := map[int]string{
		20090101: "张三",
		20090102: "李四",
		20090103: "王五",
	}

	fmt.Println(c)
	fmt.Println("----------------")
	delete(c, 20090101)

	fmt.Println(c)

	rs, flag := c[20090101]
	if flag {
		fmt.Println(rs)
	}

	fmt.Println(len(c))

	for k, v := range c {
		fmt.Println(k, v)
	}

	d := make(map[string]map[int]string)
	d["class1"] = make(map[int]string, 3)
	d["class1"][20200101] = "zhangsan"
	d["class1"][20200102] = "lisi"
	d["class1"][20200103] = "wangwu"

	d["class2"] = make(map[int]string, 3)
	d["class2"][20200101] = "小黑"
	d["class2"][20200102] = "小丽"
	d["class2"][20200103] = "小武"

	for k, v := range d {
		for v1, v2 := range v {
			fmt.Println(k, v1, v2)
		}
	}
}
