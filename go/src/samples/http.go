//http.go

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://itit-r.tumblr.com/rss"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("http error!")
	}
	defer res.Body.Close()

	byteArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("IO read error!")
	}
	fmt.Println(string(byteArray))

}
