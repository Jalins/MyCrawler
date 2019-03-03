package parser

import (
	"MyCrawler/Single/engine"
	"regexp"
)

var citylistRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]*)</a>`

func ParseCityList(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(citylistRe)

	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	limit := 2

	for _, value := range matches {
		result.Items = append(result.Items, "City "+string(value[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(value[1]),
			ParserFunc: ParseCity,
		})

		limit--
		if limit == 0 {
			break
		}

	}

	return result
}
