package parser

import (
	"My_Go_Study/Crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`"(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		// result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, name)
			},
		})
	}

	matche := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, ms := range matche {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(ms[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
