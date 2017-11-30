package scraping

import (
	"fmt"
	"log"

	"github.com/sclevine/agouti"
)

//売り切れ判定
func IsSoldOut(domSelection *agouti.Selection, soldOutMessage string) bool {
	text, err := domSelection.Text()
	if err != nil {
		panic("Domからのテキスト取得に失敗しました")
	}
	value, err := domSelection.Attribute("value")
	if err != nil {
		panic("Domからのvalue取得に失敗しました")
	}
	salesMassageSlice := []string{text, value}
	salesMassage := ""
	for _, v := range salesMassageSlice {
		if len(v) > 0 {
			salesMassage = v
			break
		}
	}
	if len(salesMassage) == 0 {
		panic("販売状況テキストが取得できませんでした")
	}

	fmt.Println(salesMassage) //動作テスト用

	if salesMassage == soldOutMessage {
		return true
	}
	return false
}

//WebDriver経由で指定サイトのソースを取得
func StartChrome() (*agouti.WebDriver, *agouti.Page) {
	agoutiDriver := agouti.ChromeDriver()
	agoutiDriver.Start()
	//defer agoutiDriver.Stop()
	page, _ := agoutiDriver.NewPage(agouti.Desired(agouti.Capabilities{
		"chromeOptions": map[string][]string{
			"args": []string{
				"--headless",
			},
		},
	}),
	)
	return agoutiDriver, page
}

//指定URLからdom要素を取得
func FetchDom(page *agouti.Page, URL string, selecter string) *agouti.Selection {
	log.Print(URL)
	page.Navigate(URL)
	log.Print(page.Title())

	return page.Find(selecter)
}
