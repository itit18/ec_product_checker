package scraping

import (
	"fmt"
	"log"

	"github.com/sclevine/agouti"
)

//売り切れ判定
func IsSoldOut(domSelection *agouti.Selection, soldOutMessage string) (bool, error) {
	text, err := domSelection.Text()
	if err != nil {
		return true, fmt.Errorf("Domからのテキスト取得に失敗しました: %v", err)
	}
	value, err := domSelection.Attribute("value")
	if err != nil {
		return true, fmt.Errorf("Domからのvalue取得に失敗しました: %v", err)
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
		return true, fmt.Errorf("販売状況テキストが取得できませんでした: %v", err)
	}

	fmt.Println(salesMassage) //動作テスト用

	if salesMassage == soldOutMessage {
		return true, nil
	}
	return false, nil
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
