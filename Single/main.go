package main

import (
	"MyCrawler/Single/engine"
	"MyCrawler/Single/zhenai/parser"
)

func main() {
	//fmt.Println("hzl")
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
