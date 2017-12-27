//Token管理
package gowechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
)

const (
	url_quarytoken = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET"
)

/*******************************************************************************************
*******************************   TYPE DEFINITIONS 	****************************************
*******************************************************************************************/

//全局token变量
var GAccessToken string

//微信服务器返回token的格式
type AccessTokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

//微信服务器返回错误信息的格式
type AccessTokenErrorResponse struct {
	Errcode float64
	Errmsg  string
}

/*******************************************************************************************
********************************   BASIC FUNCTIONS  ****************************************
*******************************************************************************************/

//生成获取Token请求的url
//@appid:微信后台的id号
//@appsecret:微信后台开发的secret号
func token_GetUrl(appid, appsecret string) string {
	get_token := url_quarytoken
	get_token = strings.Replace(get_token, "APPID", appid, -1)
	get_token = strings.Replace(get_token, "APPSECRET", appsecret, -1)
	return get_token
}

//执行获取Token请求
//@appid:微信后台的id号
//@appsecret:微信后台开发的secret号
func token_Get(appid, appsecret string) (*AccessTokenResponse, error) {

	//发送请求
	resp, err := http.Get(token_GetUrl(appid, appsecret))
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("获取AccessToken错误:发送请求", err)
		return nil, err
	}

	//接收到返回数据
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //此处可增加输入过滤
	if err != nil {
		fmt.Println("获取AccessToken错误:读取body", err)
		return nil, err
	}

	//解析返回数据
	if bytes.Contains(body, []byte("access_token")) {
		atr := AccessTokenResponse{}
		err = json.Unmarshal(body, &atr)
		if err != nil {
			fmt.Println("获取AccessToken错误:解析json错误", err)
			return nil, err
		}
		return &atr, nil
	} else {

		ater := AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		fmt.Printf("获取AccessToken错误:接收到错误返回 %+v\n", ater)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", ater.Errmsg)
	}
}

//刷新Token (刷新时间:1小时)
//@appid:微信后台的id号
//@appsecret:微信后台开发的secret号
func (we *Server) Token_Refresh() {
	//由于定时器为定时末尾触发，建议在此处先获取一次Token

	b1, err1 := token_Get(we.appid, we.appsecret)
	if err1 != nil {
		fmt.Println(time.Now().String(), " 获取Token失败")
	} else {
		GAccessToken = b1.AccessToken
		we.accesstoken = b1.AccessToken
	}

	t1 := time.NewTimer(time.Hour * 1)
	for {
		select {
		case <-t1.C:

			t1.Reset(time.Hour * 1)

			b, err := token_Get(we.appid, we.appsecret)
			if err != nil {
				logs.Debug("refresh local access token failly")
			} else {
				GAccessToken = b.AccessToken
				we.accesstoken = b.AccessToken
			}

			logs.Debug("fail to refresh local access token")
		}
	}
}
