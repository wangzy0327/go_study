package main

import "fmt"

func main() {
	fmt.Println("======================sliceFunc=========================")

	//slice

	sliceFunc()
	fmt.Println("======================sliceOps=========================")
	sliceOps()

	fmt.Println("======================mapsFunc=========================")
	mapsFunc()

	fmt.Println("======================maps Example Case=========================")
	fmt.Println("=====================lengthOfNonRepeatingSubStr==========================")
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("我是没有我的"))
	fmt.Println(lengthOfNonRepeatingSubStrForIntel("八百标兵奔北坡,炮兵并排北边跑"))
}
