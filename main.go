package main

import (
	"fmt"
	"regexp"
)

func main() {
	// readAll := Fetch("http://www.zhenai.com/zhenghun")
	// getCityList(readAll)

}

func getCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]*)</a>`)

	matches := re.FindAllSubmatch(contents, -1)

	for _, value := range matches {
		fmt.Printf("City: %s, URL: %s", value[2], value[1])
		fmt.Println()
	}

	fmt.Println(len(matches))
}
