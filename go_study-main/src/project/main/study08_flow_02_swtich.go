package main
import "fmt"

func main(){
	var score int
	fmt.Scanln(&score)

	switch score/10 {
		case 10:
			fmt.Println("perfect")
		case 9:
			fmt.Println("nice")
		case 8:
			fmt.Println("good")
		case 7,6:
			fmt.Println("pass")
			fallthrough //穿透到下一层
		case 5,4,3,2,1,0:
			fmt.Println("bad")
		default:
			fmt.Println("error")
	}


}