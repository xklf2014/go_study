package main

import "fmt"

func main() {

	num := add(5, 10)
	fmt.Println(num)

	num1 := minus(20, 10)
	fmt.Println(num1)

	//var rs1 int
	//var rs2 int
	rs1, rs2 := addAndMinus(10, 5)
	fmt.Printf("rs1=%d and rs2=%d \n", rs1, rs2)

	rs3, _ := addAndMinus(10, 5)
	fmt.Printf("rs3=%d \n", rs3)

}

func add(a int, b int) int {
	var rs int = a + b
	return rs
}

func minus(a int, b int) int {
	return a - b
}

func by(a int, b int) int {
	return a * b
}

func devide(a int, b int) int {
	if b != 0 {
		return a / b
	}
	return -1
}

func addAndMinus(a int, b int) (int, int) {
	var sums = a + b
	var minus = a - b
	return sums, minus
}
