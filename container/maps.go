package main

import "fmt"


/**
map使用 哈希表 ，必须可以比较相等

除了 slice,map,function的内建类型都可以作为key

Struct类型不包含上述字段，也可以作为key

 */
func mapsFunc() {
	m:= map[string]string{
		"name":"ccmouse",
		"course":"golang",
		"site":"imooc",
	}
	m2 := make(map[string]int) //m2 == empty map

	var m3 map[string]int  // m3 == nil

	fmt.Println(m,m2,m3)

	fmt.Println("Traversing map")
	for k,v := range m{
		fmt.Println(k,v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Printf("%s\n",courseName)
	//如果键值不存在，打印空串
	//causeName := m["cause"]
	//fmt.Printf("%q\n",causeName)
	// ""

	courseName,ok := m["course"]
	fmt.Println(courseName,ok)
	//golang true
	if causeName,ok := m["cause"];ok{
		fmt.Println(causeName)
	}else{
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name,ok := m["name"]
	fmt.Printf("%q,%v \n",name,ok)
	//"ccmouse",true
	delete(m,"name")
	name,ok = m["name"]
	fmt.Printf("%q,%v \n",name,ok)
	//"",false
}

func lengthOfNonRepeatingSubStr(s string) int{
	record := make(map[byte]int)
	start := 0
	maxLen := 0
	for i,ch := range []byte(s){
		if ind,ok := record[ch];ok && ind >= start{
			start = ind + 1
		}
		if maxLen < (i - start + 1 ){
			maxLen = (i - start + 1)
		}
		record[ch] = i
	}
	return maxLen
}

func lengthOfNonRepeatingSubStrForIntel(s string) int{
	record := make(map[rune]int)
	start := 0
	maxLen := 0
	for i,ch := range []rune(s){
		if ind,ok := record[ch];ok && ind >= start{
			start = ind + 1
		}
		if maxLen < (i - start + 1 ){
			maxLen = (i - start + 1)
		}
		record[ch] = i
	}
	return maxLen
}