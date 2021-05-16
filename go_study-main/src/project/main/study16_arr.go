package main

import "fmt"

func main() {

	//数组初始化 -- 第一种
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr1)

	//第二种
	var arr2 = [3]int{4, 5, 6}
	fmt.Println(arr2)

	//第三种
	var arr3 = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr3)

	//第四种
	var arr4 = [...]int{1: 100, 2: 200, 0: 0, 4: 400, 3: 300}
	fmt.Println(arr4)

	upd(&arr2)
	fmt.Println(arr2)
}

func upd(arr *[3]int) {
	(*arr)[0] = 10
}

func sum_and_avg1() {
	var scores [5]int
	scores[0] = 100
	scores[1] = 76
	scores[2] = 54
	scores[3] = 88
	scores[4] = 39
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}

	avg := sum / len(scores)
	fmt.Printf("总和为:%v,平均值为%v", sum, avg)
}

func sum_and_avg2() {
	var scores [5]int

	for i := 0; i < len(scores); i++ {
		fmt.Printf("请输入第%d个学生的分数", i+1)
		fmt.Scan(&scores[i])
	}

	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}

	avg := sum / len(scores)
	fmt.Printf("总和为:%v,平均值为%v \n", sum, avg)

	for i := 0; i < len(scores); i++ {
		fmt.Printf("第%d个学生的分数%d \n", i+1, scores[i])
	}

	fmt.Println("--------------------")

	for key, val := range scores {
		fmt.Printf("第%d个学生的分数%d \n", key+1, val)
	}
}
