package parser

import (
	"My_Go_Study/Crawler/engine"
	"My_Go_Study/Crawler/model"
	"regexp"
	"strconv"
)

var (
	ageRe    = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`)
	HeightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)cm</div>`)
	WeightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)kg</div>`)
)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = name
	age, _ := strconv.Atoi(extractString(contents, ageRe))
	profile.Age = age
	height, _ := strconv.Atoi(extractString(contents, HeightRe))
	profile.Height = height
	weight, _ := strconv.Atoi(extractString(contents, WeightRe))
	profile.Weight = weight

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result

	// match := ageRe.FindSubmatch(contents)
	// if match != nil {
	// 	age, err := strconv.Atoi(string(match[1]))
	// 	if err != nil {
	// 		profile.Age = age
	// 	}
	// }

	// match := marriageRe.FindSubmatch(contents)
	// if match != nil {
	// 	profile.Marriage = strconv.Atoi(string(match[1]))
	// }
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
