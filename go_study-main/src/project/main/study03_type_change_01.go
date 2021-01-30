package main
import "fmt"

func main(){
	var i1 int = 100
	//var f1 float32 = i1  //cannot use i1 (type int) as type float32 in assign

	fmt.Println(i1)

	//需要强制转换
	var f1 float32 = float32(i1)
	fmt.Println(f1)

	fmt.Printf("%T \n",i1)

	//大数据类型转小数据类型，会产生数据溢出
	var i2 int64 = 888888888
	var i3 int8 = int8(i2)
	fmt.Println(i3)

	var i4 int32 = 10000
	// cannot use i4 + 100 (type int32) as type int64 in assignment
	//var i5 int64 = i4 + 100  需要保证等于2边数据类型一致
	var i5 int64 = int64(i4) + 100
	fmt.Println(i5)

	var i6 int8 = 127
	var i7 int8 = i6 + 100
	fmt.Println(i7)
}