package main
import "fmt"

func main(){
	var name string
	var age int 
	var score float32
	var isVip bool
	fmt.Println("请输入学生姓名,年龄，成绩，是否是VIP，使用空格分隔")
	fmt.Scanf("%s %d %f %t",&name,&age,&score,&isVip)
	fmt.Printf("学生的姓名为%v,年龄%v,分数%f,VIP %t \n",name,age,score,isVip)
}

func type_stu(){
	var age int
	fmt.Println("请输入学生的年龄")
	fmt.Scanln(&age)

	var name string
	fmt.Println("请输入学生的姓名")
	fmt.Scanln(&name)

	var score float32
	fmt.Println("请输入学生的分数")
	fmt.Scanln(&score)

	var isVip bool
	fmt.Println("请输入该学生是否是VIP")
	fmt.Scanln(&isVip)


	fmt.Println(age)

	fmt.Printf("学生的姓名为%v,年龄%v,分数%f,VIP %t \n",name,age,score,isVip)
}