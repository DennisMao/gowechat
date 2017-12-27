//客服信息
package gowechat

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	service_add  = "https://api.weixin.qq.com/customservice/kfaccount/add?access_token="
	service_del  = "https://api.weixin.qq.com/customservice/kfaccount/del?access_token="
	service_get  = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token="
	service_send = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token="
)

//客服消息:客服账号添加
//@accesstoken：当前有效的token
//未设置传参,可直接修改kf_account、nickname和password。注意kf_account格式是 【自定义客服名@微信公众号账号】
func (we *Server) Service_Add() error {

	add := service_add + we.accesstoken
	//创建请求
	postReq, err := http.NewRequest("POST",
		add, //post链接
		strings.NewReader(`
				{
		     "kf_account" : "test@gh_e618fe5483dc",
		     "nickname" : "客服1",
		     "password" : "pswmd5"
		}
		`)) //post内容

	if err != nil {
		fmt.Println("增加客服:创建请求失败", err)
		return err
	}

	//增加header
	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	//执行请求
	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		fmt.Println("增加客服:创建请求失败", err)
		return err
	} else {
		//读取响应
		body, err := ioutil.ReadAll(resp.Body) //此处可增加输入过滤
		if err != nil {
			fmt.Println("增加客服:读取body失败", err)
			return err
		}

		fmt.Println("增加客服:创建成功", string(body))
	}
	defer resp.Body.Close()

	return nil
}

//客服消息:信息发送
//@openid:发送对象的微信id
//@content:需要发送的文字
func (we *Server) Service_Send(openid, content string) error {

	send := service_send + we.accesstoken
	reqbody := `
		{
 	   "touser":"OPENID",
  		  "msgtype":"text",
 		   "text":
  		  {
   	      "content":"THECONTENT"
  		  }
		}
		`
	reqbody = strings.Replace(reqbody, "OPENID", openid, -1)
	reqbody = strings.Replace(reqbody, "THECONTENT", content, -1)
	//创建请求
	postReq, err := http.NewRequest("POST",
		send, //post链接
		strings.NewReader(reqbody)) //post内容

	if err != nil {
		fmt.Println("客服消息发送:创建请求失败", err)
		return err
	}

	//增加header
	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	//执行请求
	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		fmt.Println("客服消息发送:创建请求失败", err)
		return err
	} else {
		//读取响应
		body, err := ioutil.ReadAll(resp.Body) //此处可增加输入过滤
		if err != nil {
			fmt.Println("客服消息发送:读取body失败", err)
			return err
		}

		fmt.Println("客服消息发送:创建成功", string(body))
	}
	defer resp.Body.Close()

	return nil
}

//客服消息:信息发送
//@openid:发送对象的微信id
//@content:需要发送的文字
func Service_Send_G(accesstoken, openid, content string) error {

	send := service_send + accesstoken
	reqbody := `
		{
 	   "touser":"OPENID",
  		  "msgtype":"text",
 		   "text":
  		  {
   	      "content":"THECONTENT"
  		  }
		}
		`
	reqbody = strings.Replace(reqbody, "OPENID", openid, -1)
	reqbody = strings.Replace(reqbody, "THECONTENT", content, -1)
	//创建请求
	postReq, err := http.NewRequest("POST",
		send, //post链接
		strings.NewReader(reqbody)) //post内容

	if err != nil {
		fmt.Println("客服消息发送:创建请求失败", err)
		return err
	}

	//增加header
	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	//执行请求
	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		fmt.Println("客服消息发送:创建请求失败", err)
		return err
	} else {
		//读取响应
		body, err := ioutil.ReadAll(resp.Body) //此处可增加输入过滤
		if err != nil {
			fmt.Println("客服消息发送:读取body失败", err)
			return err
		}

		fmt.Println("客服消息发送:创建成功", string(body))
	}
	defer resp.Body.Close()

	return nil
}
