package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
Fetch函数用于获取网页的信息
 */
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	readall, err := ioutil.ReadAll(resp.Body)

	return readall, nil

}
