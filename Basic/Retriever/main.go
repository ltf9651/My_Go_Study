package main

import (
	"My_Go_Learning/Basic/Retriever/mockRetriever"
	"My_Go_Learning/Basic/Retriever/realRetriever"
	"fmt"
	"time"
)

const url = "http://www.baidu.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(r Poster) {
	r.Post(url, map[string]string{
		"name": "guest",
		"lang": "go",
	})
}

func session(s RetrieverPoster) {
	s.Post(url, map[string]string{
		"name": "guest",
		"lang": "go",
	})
	s.Get(url)
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mockRetriever.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *realRetriever.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

func main() {
	var r Retriever
	r = mockRetriever.Retriever{"this is fake baidu.com"}
	fmt.Println(download(r))
	fmt.Printf("%T, %v\n", r, r)

	r = &realRetriever.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T, %v\n", r, r)

	// fmt.Println(download(r))

	inspect(r)

	realRetriever := r.(*realRetriever.Retriever)
	fmt.Println(realRetriever.TimeOut)

	// mockRetriever := r.(mockRetriever.Retriever)
	// fmt.Println(mockRetriever.Contents)
	if mockRetriever, ok := r.(mockRetriever.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("this is not a mockRetriever")
	}
}
