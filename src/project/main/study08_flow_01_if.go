package main
import "fmt"

func main(){
	// 数量<30则库存不足
	fmt.Println("数量<30则库存不足，请输入当前库存数量")
	var count int
	fmt.Scanln(&count)
	
	if count < 30 {
		fmt.Println("对不起，库存不足")
	}else {
		count -= 30
		fmt.Println("库存剩余",count)
	}

	fmt.Println("请输入学生的英语成绩0-100")
	var score float32
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("perfect")
	}else if score >= 90 {
		fmt.Println("nice")
	}else if score >= 80 {
		fmt.Println("good")
	}else if score >= 60 {
		fmt.Println("pass")
	}else {
		fmt.Println("bad")
	}
}