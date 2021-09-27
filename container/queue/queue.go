package queue

//type Queue []int

//支持任何类型
type Queue []interface{}

func(q *Queue) Push(v interface{}){
	*q = append(*q,v)
	//强制转为int 类型 v.(int)
	*q = append(*q,v.(int))
}

func(q *Queue) Pop() interface{}{
	head := (*q)[0]
	*q = (*q)[1:]
	//强制转为int类型
	//return head
	return head.(int)
}

func (q *Queue) IsEmpty() bool{
	return len(*q) == 0
}
