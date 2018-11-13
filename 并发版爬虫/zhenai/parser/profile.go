
package parser

import (
	"regexp"
	"groot/spider/engine"
	"strconv"
	"groot/spider/model"
	// "fmt"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08>([\d]+)岁</div>`)
var basiceRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08>([^<]+)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	

	profile := model.Profile{}

	profile.Name = name
	// age是int类型 strconv.Atoi把dtring转成int
	age, err := strconv.Atoi(extractString(contents, ageRe))

	if err != nil {
		profile.Age = age
	}
	// Marriage时string类型，直接赋值
	profile.Marriage = extractString(contents, basiceRe)
	profile.Xinzuo = extractString(contents, basiceRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string{
	// 找到第一个匹配项 FindSubmatch [][]byte ，
	// 返回nil没有match
	match := re.FindSubmatch(contents)
	// for _, m := range match {
	// 	fmt.Println("haha",string(m))
	// }
	if len(match) >= 2 {

		return string(match[1])
	} else {
		return ""
	}
}









