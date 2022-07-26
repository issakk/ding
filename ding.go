package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	httpPost()
}

type DingReqContent struct {
	Content string `json:"content"`
}

type DingReq struct {
	Msgtype string         `json:"msgtype"`
	Text    DingReqContent `json:"text"`
}

func httpPost() {
	// 所有URLs都被这个函数处理
	// http.HandleFunc使用了DefaultServeMux
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		body, _ := ioutil.ReadAll(r.Body)
		w.Write(body)
		fmt.Println(string(body))
		trans(r.URL.Path, string(body))
	})

	// Continue to process new requests until an error occurs
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}
	log.Fatal(http.ListenAndServe(port, nil))

	// handle error
}

func trans(accessToken string, rawContent string) {
	content := DingReq{
		Msgtype: "text",
		Text:    DingReqContent{Content: "content:" + rawContent},
	}
	a, err := json.Marshal(content)
	fmt.Printf("Map序列化后=%v\n", string(a))
	payload2 := strings.NewReader(string(a))
	preUrl := "https://oapi.dingtalk.com/robot/send?access_token="
	url := preUrl + strings.TrimLeft(accessToken, "/")
	resp, err := http.Post(url,
		"application/json",
		payload2)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}

	fmt.Println(string(body))
}
