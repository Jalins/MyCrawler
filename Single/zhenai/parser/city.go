package parser

import (
	"MyCrawler/Single/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityRe)

	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, value := range matches {
		result.Items = append(result.Items, "User "+string(value[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(value[1]),
			ParserFunc: engine.ParserFunc,
		})
	}

	return result

}
