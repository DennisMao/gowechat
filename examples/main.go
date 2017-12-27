package main

import (
	"fmt"
	_ "teaer/routers"

	"github.com/astaxie/beego"
	"github.com/dennismao/gowechat"
)

const (
	Appid     = ""
	Appsecret = ""
	Token     = ""
)

func main() {

	//	Substantialize server
	myWechat, err := gowechat.New("", "", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	//	refresh token
	go myWechat.Token_Refresh()

	//	Initialization
	//	Create custom menu
	err = myWechat.CustomMenuCreate([]gowechat.CustomButton{
		wechat.CustomButton{
			Name: "主菜单",
			SubButton: []gowechat.Button{
				wechat.Button{
					Type: "views",
					Name: "百度搜索",
					Url:  "www.baidu.com",
				},
				gowechat.Button{
					Type: "click",
					Name: "点击事件",
					Key:  "V1000_button1",
				},
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	//	Run the http server
	beego.Run()
}
