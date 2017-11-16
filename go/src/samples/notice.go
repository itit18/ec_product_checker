//notice.go
//AWS SNS経由でnotice送信する例

package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
)

func main() {
	//credentialは環境変数にて指定すること
	sess := session.Must(session.NewSession())
	svc := sns.New(sess)

	//topicリスト取得
	list, err := svc.ListTopics(nil)
	if err != nil {
		panic("error")
	}
	for _, v := range list.Topics {
		log.Print(v)
	}

	//topic発行
	params := &sns.PublishInput{}
	params.SetTopicArn("arn:aws:sns:ap-northeast-1:706437443163:notice_switch")
	params.SetMessage("テスト通知")
	_, err = svc.Publish(params)
	if err != nil { // resp is now filled
		panic("error")
	}

	log.Print("ok")
}
