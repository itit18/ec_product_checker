//webdriber.go

package main

import (
	"log"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
	// ブラウザはChromeを指定して起動
	// Chromeを利用することを宣言
	agoutiDriver := agouti.ChromeDriver()
	agoutiDriver.Start()
	defer agoutiDriver.Stop()
	//page, _ := agoutiDriver.NewPage()
	page, _ := agoutiDriver.NewPage(agouti.Desired(agouti.Capabilities{
		"chromeOptions": map[string][]string{
			"args": []string{
				"--headless",
			},
		},
	}),
	)

	// 自動操作
	page.Navigate("https://qiita.com/")
	log.Print(page.Title())
	//page.Screenshot("Screenshot01.png")

	page.FindByLink("もっと詳しく").Click()
	log.Print(page.Title())
	//page.Screenshot("Screenshot02.png")
	time.Sleep(3 * time.Second)

}
