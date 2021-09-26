package main

import (
	"fmt"
	"math"
)

func main() {

	//function

	fmt.Println("=======================Functions ========================")
	//fmt.Println(eval(3,4,"x"))
	if result,err := eval(3,4,"x");err != nil{
		fmt.Println("Error:",err)
	}else{
		fmt.Println(result)
	}
	fmt.Println(apply(pow,3,4))
	fmt.Println(apply(func(a ,b int) int {
		return int(math.Pow(float64(a),float64(b)))
	},3,4))
	q,r := div(13,3)
	fmt.Println(q,r)

	fmt.Println("=======================Func params transfer========================")


	//func value transfer
	a,b := 3,5
	fmt.Printf("before swap : %d,%d \n",a,b)
	swap1(a,b)
	fmt.Printf("after swap value transfer: %d,%d \n",a,b)
	//before swap : 3,5
	//after swap value transfer : 3,5

	swap2(&a,&b)
	fmt.Printf("after swap reference transfer: %d,%d \n",a,b)
	//after swap reference transfer: 5,3


}
