//scraping.go

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/sclevine/agouti"
)

type configStruct struct {
	General generalConfig `toml:"general"`
	Sites []siteConfig `toml:"site"`
}

type generalConfig struct {
	TopicArn string
}

type siteConfig struct {
	Name            string
	URL             string
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
}

func isSoldOut(domSelection *agouti.Selection, soldOutMessage string) bool {
	text, err := domSelection.Text()
	if err != nil {
		panic("error")
	}
	value, err := domSelection.Attribute("value")
	if err != nil {
		panic("error")
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
		panic("error")
	}

	fmt.Println(salesMassage) //動作テスト用

	if salesMassage == soldOutMessage {
		return true
	}
	return false
}

//通知処理 / AWS SNSを利用
func sendNotice(message string) {
	log.Print("メッセージを送信")
	topicArn := config.General.TopicArn
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
func startChrome() (*agouti.WebDriver, *agouti.Page) {
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

func fetchDom(page *agouti.Page, URL string, selecter string) *agouti.Selection {
	log.Print(URL)
	page.Navigate(URL)
	log.Print(page.Title())

	return page.Find(selecter)
}

func main() {
	//設定値 / コンストラクタの設定と値の引き渡しのパターンがよくわからない…
	config := configStruct{}
	setConfig(&config)
	driver, page := startChrome()
	defer driver.Stop()

	for _, site := range config.Sites {
		domSelection := fetchDom(page, site.URL, site.Selecter)
		if isSoldOut(domSelection, site.SolodOutMessage) == false {
			log.Print("在庫切れでした")
		} else {
			//通知処理
			sendNotice(site.Name + " スイッチ在庫あり: " + site.URL)
		}
	}
}
