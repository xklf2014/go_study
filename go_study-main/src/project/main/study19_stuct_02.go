package main

import "fmt"

func main() {
	var s Student = Student{10}
	var p Person = Person{5}

	s = Student(p)
	fmt.Println(s)

	fmt.Println("---------------")
	var s1 Student = Student{19}
	var s2 Stu = Stu{10}
	s1 = Student(s2)
	fmt.Println(s1)
}

type Student struct {
	Age int
}

type Stu Student

type Person struct {
	Age int
}
