package engine

import (
	"MyCrawler/Single/fetcher"
	"fmt"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		// 维护一个队列，不断地取出队列中的第一个元素
		r := requests[0]        // 拿出第一个元素，即第一个Request
		requests = requests[1:] // 拿出第一个元素后剩下的元素， 模拟一个队列

		// 通过元素的Url来获取网页的信息
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Print("Fetch error"+"fetching url %s:%v", r.Url, err)
			continue
		}

		//	通过自身携带的函数ParserFunc来处理从网页中获取到的信息，返回一个个Request以及Items
		parseResult := r.ParserFunc(body)

		// 将这些一个个Request添加进队列中
		requests = append(requests, parseResult.Requests...)

		// 将这些item一个一个打印出来
		for _, item := range parseResult.Items {
			log.Printf("Got the item %v", item)
		}

		fmt.Println("=============================================")
		//for _, url := range requests {
		//	log.Printf("Got url %v", url)
		//}
	}
}
