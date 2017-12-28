//自定义菜单
package gowechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dennismao/gowechat/utils"
)

const (
	url_create = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=" //post 添加自定义菜单，送出json
	url_query  = "https://api.weixin.qq.com/cgi-bin/menu/get?access_token="    //get 查询当前自定义菜单，返回json
	url_delete = "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=" //get 删除全部菜单

)

/*******************************************************************************************
*******************************   TYPE DEFINITIONS 	****************************************
*******************************************************************************************/

//	custon menu requset structure
type CustomMenu struct {
	Url     string
	Content []CustomButton
}

//	button menu
//	@name
type CustomButton struct {
	Type      string   `json:"type;omitempty"`
	Name      string   `json:"name;omitempty"`
	Key       string   `json:"key;omitempty"`
	Url       string   `json:"url;omitempty"`
	SubButton []Button `json:"sub_button;omitempty"`
}

//	button basic structure
//	with compat on click mode and views mode
type Button struct {
	Type string `json:"type;omitempty"`
	Name string `json:"name;omitempty"`
	Key  string `json:"key;omitempty"`
	Url  string `json:"url;omitempty"`
}

/*******************************************************************************************
********************************   BASIC FUNCTIONS  ****************************************
*******************************************************************************************/

//
//	url format:  https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
func MenuCreate(accesstoken string, menuContent []CustomButton) (error, *CustomMenu) {
	var menu CustomMenu
	menu.Url = url_create
	menu.Url += accesstoken
	menu.Content = menuContent
	return nil, &menu
}

//	create cunstom menu
func (we *Server) CustomMenuCreate(menuContent []CustomButton) error {
	err, customMenuRequest := MenuCreate(we.accesstoken, menuContent)
	if err != nil {
		return err
	}

	err = utils.PostJson(customMenuRequest.Url, customMenuRequest.Content)
	if err != nil {
		return err
	}

	return nil
}

//查询当前自定义菜单
//@accesstoken: 传入当前有效的token
func CustomMenuQuery(accesstoken string) error {
	//整理url
	quary := url_query
	quary += accesstoken

	//发送请求
	resp, err := http.Get(quary)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("获取菜单错误:发送请求", err)
		return err
	}

	//接收到返回数据
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //此处可增加输入过滤
	if err != nil {
		fmt.Println("获取菜单错误:读取body", err)
		return err
	}

	//解析返回数据
	if bytes.Contains(body, []byte("menu")) {
		fmt.Println("获取菜单成功:", string(body))
		return nil
	} else {

		ater := AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		fmt.Printf("获取菜单错误:接收到错误返回 %+v\n", ater)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", ater.Errmsg)
	}
}

//删除当前自定义菜单
//@accesstoken: 传入当前有效的token
func Menu_Delete(accesstoken string) error {
	//整理url
	del := url_delete
	del += accesstoken

	//发送请求
	resp, err := http.Get(del)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("删除菜单错误:发送请求", err)
		return err
	}

	//接收到返回数据
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //此处可增加输入过滤
	if err != nil {
		fmt.Println("删除菜单错误:读取body失败", err)
		return err
	}

	//解析返回数据
	if bytes.Contains(body, []byte("menu")) {
		fmt.Println("删除菜单成功:", string(body))
		return nil
	} else {

		ater := AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		fmt.Printf("删除菜单错误:接收到错误返回 %+v\n", ater)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", ater.Errmsg)
	}
}
