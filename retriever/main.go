package main

import (
	"fmt"
	"go-demo/retriever/mock"
	"go-demo/retriever/real"
	"time"
)

type Retriever interface {
	//func
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string{
	return r.Get(url)
}

func post(poster Poster){
	poster.Post(url,
		map[string]string{
		"name":"ccmouse",
		"course":"golang",
		})
}

//这里组合了Retriever 和 Poster两个接口
type RetrieverPoster interface{
	Retriever
	Poster
}

func session(s RetrieverPoster)string{
	s.Get(url)
	s.Post(url, map[string]string{
		"contents":"another faked imooc.com",
	})
	return s.Get(url)
}

func main(){
	var r Retriever
	retriever := &mock.Retriever{"this is a fake imoooc.com"}
	r = retriever
	inspect(r)
	//r = real.Retriever{}
	r = &real.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}
	inspect(r)

	//Type assertion
	if mockRetriever,ok := r.(*mock.Retriever);ok{
		fmt.Println(mockRetriever.Contents)
	}else{
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(retriever))


	/**
	mock.Retriever {this is a fake imoooc.com}
	real.Retriever { 0s}
	 */
	//fmt.Println(download(r))
}

func inspect(r Retriever){
	fmt.Printf("%T %v\n",r,r)
	switch v:= r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:",v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:",v.UserAgent)
	}
}
