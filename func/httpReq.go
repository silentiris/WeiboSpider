package _func

import (
	"WeiboSpider/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func HttpReqCard() *model.CardList {
	params := url.Values{}
	Url, _ := url.Parse("https://m.weibo.cn/api/container/getIndex")
	params.Set("containerid", "1076031989660417")
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	resp, _ := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	cardlist := new(model.CardList)
	json.Unmarshal(body, &cardlist)
	return cardlist
}

func HttpReqLongText(list *model.CardList) string {
	params := url.Values{}
	Url, _ := url.Parse("https://m.weibo.cn/statuses/extend")
	params.Set("id", list.Data.Cards[0].Mblog.ID)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	resp, _ := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	longText := new(model.LongText)
	json.Unmarshal(body, &longText)
	longText.Data.LongTextContent = strings.ReplaceAll(longText.Data.LongTextContent, "<br /><br />", "\n")
	longText.Data.LongTextContent = strings.ReplaceAll(longText.Data.LongTextContent, "<br />", "\n")
	return longText.Data.LongTextContent
}

func HttpReqSync(result *model.Result) {
	bytesData, _ := json.Marshal(result)
	resp, _ := http.Post("http://127.0.0.1:8080/weibo", "application/json", bytes.NewReader(bytesData))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
