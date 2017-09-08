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
	"path/filepath"
	"strings"
	"time"
)

type configStruct struct {
	General generalConfig `toml:"general"`
	Sites   []siteConfig  `toml:"site"`
}

type generalConfig struct {
	ChromeDriver string
	SiteName     string
	SiteUrl      string
	Selecter     string
}

type siteConfig struct {
	Name            string
	Url             string
	Selecter        string
	SolodOutMessage string
}

func setConfig(config *configStruct) {
	appPath, _ := os.Executable()
	configPath := [3]string{}
	configPath[0] = "."
	configPath[1] = os.Getenv("GOPATH") + "/src/config"
	configPath[2] = filepath.Dir(appPath)

	var err error
	//いくつかのfilepath候補からconfigを探す
	for _, path := range configPath {
		log.Print(path)
		_, err = toml.DecodeFile(path+"/config.toml", &config)
		if err == nil {
			break
		}
	}
	//configが1つも見つからなければ強制終了
	if err != nil {
		panic(err)
	}

	log.Print(config.General) //動作確認用
}

func isSoldOut(domObject *goquery.Selection, soldOutMessage string) bool {
	text := ""
	text = domObject.Text()
	if text == "" { //売り切れ表示がinput要素のときがあるので
		text, _ = domObject.Attr("value")
	}
	if text == "" {
		panic("要素が見つかりません")
	}
	fmt.Println(text) //動作テスト用

	if text == soldOutMessage {
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
	_, err := svc.Publish(params)
	if err != nil { // resp is now filled
		panic("error")
	}
}

//WebDriver経由で指定サイトのソースを取得
func fetchHtml(config *configStruct, url string) string {
	chromeDriver := webdriver.NewChromeDriver(config.General.ChromeDriver)
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
	session.Url(url)
	source, err := session.Source()
	if err != nil {
		panic(err)
	}
	defer session.Delete()
	//ページの読み込み待ち
	time.Sleep(1 * time.Second)

	return source
}

func fetchDom(source string, selecter string) *goquery.Selection {
	r := strings.NewReader(source)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		panic(err)
	}
	return doc.Find(selecter).First()
}

func main() {
	//設定値 / コンストラクタの設定と値の引き渡しのパターンがよくわからない…
	config := configStruct{}
	setConfig(&config)

	for _, site := range config.Sites {
		source := fetchHtml(&config, site.Url)

		domObject := fetchDom(source, site.Selecter)
		if isSoldOut(domObject, site.SolodOutMessage) == false {
			log.Print("在庫切れでした")
		} else {
			//通知処理
			sendNotice(site.Name + " スイッチ在庫あり: " + site.Url)
		}
	}
}
