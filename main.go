package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type RequestPayload struct {
	Msgtype string `json:"msgtype"`
}

type TextMessage struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type TextMessagePayload struct {
	RequestPayload
	Text TextMessage `json:"text"`
}

type APIResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

const defaultWechatWorkHookURL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"

func main() {
	url := os.Getenv("PLUGIN_HOOK_URL")
	key := os.Getenv("PLUGIN_KEY")
	content := os.Getenv("PLUGIN_CONTENT")
	mentionedList := strings.Split(os.Getenv("PLUGIN_MENTIONED_LIST"), ",")
	mentionedMobileList := strings.Split(os.Getenv("PLUGIN_MENTIONED_MOBILE_LIST"), ",")

	if url == "" {
		if key == "" {
			log.Fatalln(errors.New("Error: Hook url is missed."))
		} else {
			url = fmt.Sprintf("%v?key=%v", defaultWechatWorkHookURL, key)
		}
	}

	if content == "" {
		log.Fatalln(errors.New("Error: Content is missed."))
	}

	text := TextMessage{
		Content:             content,
		MentionedList:       mentionedList,
		MentionedMobileList: mentionedMobileList,
	}

	data := TextMessagePayload{}

	data.Msgtype = "text"
	data.Text = text

	payload, err := json.Marshal(data)

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(payload))
	defer resp.Body.Close()

	log.Printf("Request  %v", string(payload))

	if err != nil {
		log.Fatalln(err)
	}

	if resp != nil {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Response %v", string(body))

		var response APIResponse

		err = json.Unmarshal(body, &response)

		if err != nil {
			log.Fatalln(err)
		}

		if response.ErrCode == 0 {
			log.Print("Succeed")
		} else {
			log.Fatalf("Error: %v.", response.ErrMsg)
		}
	}
}
