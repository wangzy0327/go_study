package main

import (
	"fmt"
	"reflect"
)

func sliceFunc() {

	// arr is array数组
	arr := [...]int{0,1,2,3,4,5,6,7}
	s:= arr[2:6]
	fmt.Printf("arr type : %s \n",reflect.TypeOf(arr))
	fmt.Printf("arr[2:6] = %v \n",s)
	fmt.Printf("s type : %s \n",reflect.TypeOf(s))
	s[0] = 100
	// slice 是 array 的一个view
	fmt.Printf("arr = %v \n",arr)
	fmt.Printf("s = %v \n",s)

	//s1 是 切片
	s1 := []int{0,1,2,3,4,5,6,7}
	fmt.Printf("s1 type : %v \n",reflect.TypeOf(s1))
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",s1,len(s1),cap(s1))
	arr[2] = 2
	s1 = arr[2:6] // [2,3,4,5]

	s2 := s1[3:5]  // [5,6]

	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",s1,len(s1),cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",s2,len(s2),cap(s2))

	s3 := append(s2,10)
	// s4 s5 no longer view of arr
	// 添加元素如果超越cap，系统会重新分配更大的底层数组
	// 由于值传递的关系 必须接收 append的返回值
	s4 := append(s3,11)
	s5 := append(s4,12)
	fmt.Printf("s3,s4,s5 = %v,%v,%v \n",s3,s4,s5)

}

func printSlice(s []int){
	fmt.Printf("%v , len = %d,cap = %d\n",s,len(s),cap(s))
}

func sliceOps() {
	fmt.Printf("Creating Slice ... \n")
	var s []int // zero value for slice is nil
	for i:= 0;i < 100;i++{
		printSlice(s)
		s = append(s,2*i+1)
	}
	/**
	len = 0,cap = 0
	len = 1,cap = 1
	len = 2,cap = 2
	len = 3,cap = 4
	len = 4,cap = 4
	len = 5,cap = 8
	len = 6,cap = 8
	len = 7,cap = 8
	len = 8,cap = 8
	len = 9,cap = 16
	len = 10,cap = 16
	len = 11,cap = 16
	len = 12,cap = 16
	len = 13,cap = 16
	len = 14,cap = 16
	len = 15,cap = 16
	len = 16,cap = 16
	len = 17,cap = 32
	len = 18,cap = 32
	len = 19,cap = 32
	len = 20,cap = 32
	len = 21,cap = 32
	len = 22,cap = 32
	*/
	fmt.Println(s)

	s1 := []int{2,4,6,8}
	printSlice(s1)

	//设置len = 16
	s2:= make([]int,16)
	//设置len = 16 , cap = 32
	s3:= make([]int,10,32)
	printSlice(s2)
	printSlice(s3)
	/**
	[2 4 6 8] , len = 4,cap = 4
	[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] , len = 16,cap = 16
	[0 0 0 0 0 0 0 0 0 0] , len = 10,cap = 32
	 */

	fmt.Printf("Copying Slice ... \n")
	copy(s2,s1)
	printSlice(s2)
	/**
	[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0] , len = 16,cap = 16
	 */

	fmt.Printf("Deleting eles from Slice ... \n")
	//删除第三个元素 append 后面采用... 表示 多个参数
	s2 = append(s2[:3],s2[4:]...)
	printSlice(s2)

	fmt.Printf("Popping from front ... \n")
	front := s2[0]
	s2 = s2[1:]

	fmt.Printf("Popping from back ... \n")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]
	fmt.Printf("%d,%d \n",front,tail)
	printSlice(s2)
}

