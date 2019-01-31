package main

import (
	"My_Go_Study/Crawler_Version_1/engine"
	"My_Go_Study/Crawler_Version_1/scheduler"
	"My_Go_Study/Crawler_Version_1/zhenai/parser"
	"fmt"
	"regexp"
)

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s", m[2], m[1])
		fmt.Println()
	}
	fmt.Printf("len:%d\n", len(matches))
}

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QuenedScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	// e.Run(engine.Request{
	// 	Url:        "https://www.zhenai.com/zhenghun/shanghai",
	// 	ParserFunc: parser.ParseCity,
	// })
}
