//模板消息
package gowechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/////////////////  通用数据结构 /////////////////////
const (
	tmpmsg_send = `https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=`
)

//模板消息发送请求格式
type ReqWechatTplMsg struct {
	Touser      string                     `json:"touser"`      //接收人微信OpenId
	TempateId   string                     `json:"template_id"` //模板ID  在微信公众平台后台生成
	Url         string                     `json:"url"`         //消息内的跳转页面
	Miniprogram ReqWechatTplMiniprogramMsg `json:"-"`           //微信小程序信息 当前未使用
	Data        interface{}                `json:"data"`        //需要传入模板的数据
}

//模板消息响应格式
type RespWechatTplMsg struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

//模板消息-微信小程序信息
type ReqWechatTplMiniprogramMsg struct {
	AppId    string `json:"appid"`    //微信小程序id
	PagePath string `json:"pagepath"` //所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
}

//模板数据-官方默认
type ReqWechatTplData_Default struct {
	Value string `json:"value"` //数据值
	Color string `json:"color"` //颜色值 #173177 (蓝色)
}

//模板-官方默认
type ReqWechatTpl_Default struct {
	First    ReqWechatTplData_Default `json:"first"`    //首标题
	KeyNote1 ReqWechatTplData_Default `json:"keynote1"` //数据块1
	KeyNote2 ReqWechatTplData_Default `json:"keynote2"` //数据块2
	KeyNote3 ReqWechatTplData_Default `json:"keynote3"` //数据块3
	Remark   ReqWechatTplData_Default `json:"remark"`   //备注 可用于拓展
}

//模板-告警2.2.1
type ReqWechatTpl_Alarm struct {
	First    ReqWechatTplData_Default `json:"first"`       //首标题
	KeyNote1 ReqWechatTplData_Default `json:"performance"` //数据块1
	KeyNote2 ReqWechatTplData_Default `json:"time"`        //数据块2
	Remark   ReqWechatTplData_Default `json:"remark"`      //备注 可用于拓展
}

///////////////// 自定义模块数据结构 //////////////////
//模板消息:信息发送
//@openid:发送对象的微信id
//@time: 发送时间
//@content:需要发送的文字
func TplMessage_Send_G(accesstoken, openid, time, content string) error {

	send := tmpmsg_send + accesstoken
	//整理内容

	var tmpbody ReqWechatTplMsg
	tmpbody.Data = ReqWechatTpl_Alarm{
		First: ReqWechatTplData_Default{Value: "标题!!",
			Color: "#173177"},
		KeyNote1: ReqWechatTplData_Default{Value: content,
			Color: "#173177"},
		KeyNote2: ReqWechatTplData_Default{Value: time,
			Color: "#173177"},
		Remark: ReqWechatTplData_Default{Value: "备注",
			Color: "#173177"}}
	tmpbody.TempateId = "nUZcMXi2WOrTByGeYNs1EkNRP_mavGRqGFocJQyP5cU"
	tmpbody.Touser = openid
	b, _ := json.Marshal(tmpbody)
	//创建请求
	resp, err := http.Post(send, //post链接
		"application/json; encoding=utf-8",
		bytes.NewReader(b)) //post内容
	if err != nil {
		fmt.Println("模板消息发送:创建请求失败", err)
		return err
	} else {
		//读取响应
		body, err := ioutil.ReadAll(resp.Body) //此处可增加输入过滤
		if err != nil {
			fmt.Println("模板消息发送:读取body失败", err)
			return err
		}

		fmt.Println("模板消息发送:创建成功", string(body))
	}
	defer resp.Body.Close()

	return nil
}
