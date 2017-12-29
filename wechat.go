package gowechat

import (
	"errors"

	"github.com/astaxie/beego/logs"
)

// wechat server basic structure
type Server struct {
	appid       string
	appsecret   string
	token       string
	accesstoken string
}

// global variale
var (
	//user's token
	user_token string = ""
)

//	substantialize a wechat server
func New(appid, appsecret, token string) (*Server, error) {
	if appid == "" || appsecret == "" || token == "" {
		return nil, errors.New("Error,Invalid parameters")
	}
	user_token = token

	b, err := token_Get(appid, appsecret)
	if err != nil {
		logs.Debug("Failed to get access_token")
		return nil, err
	} else {
		GAccessToken = b.AccessToken
	}

	return &Server{
		appid:       appid,
		appsecret:   appsecret,
		token:       token,
		accesstoken: b.AccessToken}, nil
}

//	return the user token
func ServerToken() (string, error) {
	if user_token != "" {
		return user_token, nil
	}
	return "", errors.New("Error the user token is empty")
}
