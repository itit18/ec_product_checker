//webdriber.go

package main

import (
	"fmt"
	"github.com/fedesog/webdriver"
)

func main() {
	path := "/usr/local/Cellar/chromedriver/2.31/bin/chromedriver"
	chromeDriver := webdriver.NewChromeDriver(path)
	err := chromeDriver.Start()
	if err != nil {
		fmt.Println(err)
	}
	desired := webdriver.Capabilities{"Platform": "Mac"}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		fmt.Println("webdriber error!")
	}

	err = session.Url("http://golang.org")
	if err != nil {
		fmt.Println("get url error!")
	}
	session.Delete()
	chromeDriver.Stop()

}
