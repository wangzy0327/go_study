package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//golang 没有引用传递 都是 值传递 ，可以通过 传地址来修改

func eval(a,b int,op string) (int,error){
	switch op {
	case "+":
		return a + b,nil
	case "-":
		return a - b,nil
	case "*":
		return a * b,nil
	case "/":
		q,_ := div(a,b)
		return q,nil
	default:
		return 0,fmt.Errorf("unsupported operation: %s",op)
	}
}

func apply(op func(int,int) int,a,b int) int{
	p:= reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d,%d) \t",opName,a,b)
	return op(a,b)
}

func pow(a,b int) int{
	return int(math.Pow(float64(a),float64(b)))
}

func swap1(a,b int){
	a,b = b,a
}

func swap2(a,b *int){
	*a,*b = *b,*a
}

func div(a,b int)(q,r int){
	return a/b, a%b;
}


