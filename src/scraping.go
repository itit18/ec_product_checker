//scraping.go

package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/fedesog/webdriver"
	"log"
	"os"
	"strings"
)

type configStruct struct {
	ChromeDriver string
	SiteName     string
	SiteUrl      string
	Selecter     string
}

func setConfig(config *configStruct) {
	config.SiteName = "オムニ7"
	config.SiteUrl = "http://7net.omni7.jp/detail/2110595636"
	config.Selecter = `#cart_whole > div.box01.boxInteractive.mod-productDetails3Column_colCartBtnWrap.u-marginBottom05 > ul > li > p > input`
	config.ChromeDriver = "/usr/local/Cellar/chromedriver/2.31/bin/chromedriver"
}

func isSoldOut(domObject *goquery.Selection) bool {
	text, exist_err := domObject.Attr("value")
	if exist_err == false {
		panic("要素が見つかりません")
	}
	fmt.Println(text)

	if text == "在庫切れ" {
		return true
	}
	return false
}

//通知処理 / AWS SNSを利用
func sendNotice(message string) {
	log.Print("メッセージを送信")
	topicArn := "arn:aws:sns:ap-northeast-1:706437443163:notice_switch"
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sns.New(sess)

	params := &sns.PublishInput{}
	params.SetTopicArn(topicArn)
	params.SetMessage(message)
	_, err = svc.Publish(params)
	if err != nil { // resp is now filled
		panic("error")
	}
}

//WebDriver経由で指定サイトのソースを取得
func fetchHtml(config *configStruct, url string) string {
	chromeDriver := webdriver.NewChromeDriver(config.ChromeDriver)
	err := chromeDriver.Start()
	if err != nil {
		panic(err)
	}
	defer chromeDriver.Stop()

	desired := webdriver.Capabilities{"Platform": "Mac"}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		panic(err)
	}
	session.Url(config.SiteUrl)
	source, err := session.Source()
	if err != nil {
		panic(err)
	}
	defer session.Delete()
	//ページの読み込み待ち
	//time.Sleep(3 * time.Second)

	return source
}

func main() {
	//設定値 / コンストラクタの設定と値の引き渡しのパターンがよくわからない…
	config := configStruct{}
	setConfig(&config)

	source := fetchHtml(&config, config.SiteUrl)
	r := strings.NewReader(source)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		panic(err)
	}

	domObject := doc.Find(config.Selecter).First()
	if isSoldOut(domObject) {
		log.Print("在庫切れでした")
		return // mainでreturnすると正常終了のステータスコードが帰らないので注意
	}

	//通知処理
	sendNotice(config.SiteName + " スイッチ在庫あり: " + config.SiteUrl)
}
