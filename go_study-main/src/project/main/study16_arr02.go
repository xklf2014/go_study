package main

import "fmt"

func main() {
	//二维数组
	var arr [2][3]int16
	fmt.Println(arr)

	fmt.Printf("arr地址%p \n", &arr)
	fmt.Printf("arr[0]地址%p \n", &arr[0])
	fmt.Printf("arr[0][0]地址%p \n", &arr[0][0])
	fmt.Printf("arr[1]地址%p \n", &arr[1])

	var arr1 [2][3]int16 = [2][3]int16{{1, 3, 6}, {2, 4, 7}}
	fmt.Println(arr1)

	fmt.Println("------------------------")

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Printf("arr[%d][%d]的值为%d \n", i, j, arr1[i][j])
		}
	}
	fmt.Println("------------------------")

	for k := range arr1 {
		for k1, v1 := range arr1[k] {
			fmt.Printf("arr[%d][%d]：%d\n", k, k1, v1)
		}
	}
}
