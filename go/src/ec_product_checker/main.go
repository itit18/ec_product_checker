//scraping.go

package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"

	"ec_product_checker/model/scraping"
)

type configStruct struct {
	General generalConfig `toml:"general"`
	Sites   []siteConfig  `toml:"site"`
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

//通知処理 / AWS SNSを利用
func sendNotice(message string, topicArn string) {
	log.Print("メッセージを送信")
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

func main() {
	//設定値 / コンストラクタの設定と値の引き渡しのパターンがよくわからない…
	config := configStruct{}
	setConfig(&config)
	driver, page := scraping.StartChrome()
	defer driver.Stop()

	for _, site := range config.Sites {
		domSelection := scraping.FetchDom(page, site.URL, site.Selecter)
		if scraping.IsSoldOut(domSelection, site.SolodOutMessage) == false {
			log.Print("在庫切れでした")
		} else {
			//通知処理
			sendNotice(site.Name+" スイッチ在庫あり: "+site.URL, config.General.TopicArn)
		}
	}
}
