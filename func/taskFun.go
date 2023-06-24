package _func

import (
	"WeiboSpider/model"
	"fmt"
	"time"
)

func TaskFunc(flag string) func() {
	return func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered:", r)
			}
		}()
		fmt.Println(time.Now())
		cardList := HttpReqCard()
		println(flag)
		if flag != cardList.Data.Cards[0].Mblog.ID {
			flag = cardList.Data.Cards[0].Mblog.ID
			println("now flag:" + flag)
			result := new(model.Result)
			result.Content = HttpReqLongText(cardList)
			println(result.Content)
			var urls []string
			for _, v := range cardList.Data.Cards[0].Mblog.PicIds {
				urls = append(urls, "https://wx4.sinaimg.cn/large/"+v+".jpg")
			}
			println(urls)
			result.Urls = urls
			println(result.Urls[0])
			HttpReqSync(result)
		}
	}
}
