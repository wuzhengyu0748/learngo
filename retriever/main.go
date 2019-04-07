package main

import (
	"fmt"
	"wuzhengyu/learngo/retriever/mock"
	real2 "wuzhengyu/learngo/retriever/real"
	"time"
)

const url  = "http://www.baidu.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) string {
	return poster.Post(url, map[string]string{
		"name": "rainy",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string{
	s.Post(url, map[string]string{
		"contents" : "another faked baidu.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake baidu"}
	inspect(&retriever)

	fmt.Println(download(r))

	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)

	// Type assertion
	realRetriever := r.(*real2.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))

	fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}