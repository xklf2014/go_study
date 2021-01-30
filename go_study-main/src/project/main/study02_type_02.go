package main
import "fmt"

func main(){
	//定义字符类型变量
	var c1 byte = 'a'
	fmt.Println(c1)
	var c2 byte = '6'
	fmt.Println(c2)
	var c3 byte = '('
	fmt.Println(c3)

	// \n换行符
	fmt.Println("aaa\nbbb")
	// \b 退格
	fmt.Println("aaa\bbbb")
	// \r 光标回到本行开头，后续内容替换原有内容
	fmt.Println("aaa\rbbb")
	// \t 制表符   8位
	fmt.Println("aaa\tbbb")
	fmt.Println("aaaa\tbbb")
	fmt.Println("aaaaa\tbbb")
	fmt.Println("aaaaaa\tbbb")
	fmt.Println("aaaaaaa\tbbb")
	fmt.Println("aaaaaaaa\tbbb")

	fmt.Println("hello \"go lang\"")

	var flag1 bool = true
	fmt.Println(flag1)
	var flag2 bool = false
	fmt.Println(flag2)

	fmt.Println(5>6)
	
	//字符串不可变，字符串一旦定义好，其中的字符不能改变
	var s1 string = "abc"
	fmt.Println(s1)
	fmt.Println(s1[2])
	//s1[0]="d" cannot assign to s1[0]

	//特殊字符可以使用反引号`
	var s2 string = `
		package main
		import "fmt"
		
		func main(){
			fmt.Println("hello go")
			fmt.Println("hello go")
			fmt.Println("hello go")
			fmt.Println("hello go")
		
			var age = 18
			age = age + 10
			fmt.Println("age plus 10 is:" )
			fmt.Println(age)
		
			fmt.Println("asaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
			"cccccccccccccccccccccccccccccccccccc")
		}
	`

	fmt.Println(s2)

	var s3 string = "abc" + "def"
	s3 += "ghi"
	fmt.Println(s3)

	//每行默认会补充一个换行符，所以字符串拼接如果未结束，
	//要把拼接符号留在当行结尾
	var s4 string = "abc" + "def" + "abc" + "def" + "abc" + "def" + 
	"abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + 
	"abc" + "def" + "abc" + "def" + "abc" + "def"

	fmt.Println(s4)
	
}