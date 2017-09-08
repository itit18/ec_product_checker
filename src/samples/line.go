//line.go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

type RequestJson struct {
	To       string          `json:"to"`
	Messages []MessageObject `json:"messages"`
}

type MessageObject struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func main() {
	//アクセストークンと送信先は環境変数で指定してください
	token := os.Getenv("LINE_TOKEN")
	to := os.Getenv("LINE_TO")

	url := "https://api.line.me/v2/bot/message/push"
	message := MessageObject{}
	requestJson := RequestJson{}
	message.Type = "text"
	message.Text = "test messsage"
	requestJson.Messages = append(requestJson.Messages, message)
	requestJson.To = to
	requestBody, err := json.Marshal(requestJson)
	if err != nil {
		log.Println("JSON Marshal error:", err)
		panic("error")
	}

	//httpリクエストの準備
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	dump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s", dump)

	dumpResp, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", dumpResp)

}
