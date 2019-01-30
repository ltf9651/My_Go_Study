package engine

import (
	"My_Go_Study/crawler/fetcher"
	"github.com/gpmgo/gopm/modules/log"
)

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	//log.Warn("Fetching %s", r.Url)
	if err != nil {
		log.Error("请求[%s]失败：%s", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parse.Parse(body,r.Url), nil
}
