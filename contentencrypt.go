//	message encryption and decryption
package gowechat

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/astaxie/beego/logs"
)

/*******************************************************************************************
*******************************   TYPE DEFINITIONS 	****************************************
*******************************************************************************************/

//	Message pattern(safe mode)
type Content_Msg_Safe struct {
	ToUserName string `xml:"ToUserName"` //消息发送者微信openid
	Encrypt    string `xml:"Encrypt"`    //加密内容
}

type Content_Msg_Clear struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        string   `xml:"MsgId"`
}

//	Message pattern(content mode)

/*******************************************************************************************
********************************   BASIC FUNCTIONS  ****************************************
*******************************************************************************************/

//	Verify signature
// 	@token:	user token
//	@timestamp: timestamp of verification request
//	@nonce: nonce of verification request
//	@echostr: the echo should be responded when verification passes
//  @return:
//  response message returned to wechat server directly
func SignatureVerify(timestamp, nonce, remoteSignature, echostr string) string {

	// get latest token
	tk, err := ServerToken()
	if err != nil {
		logs.Debug("local accesstoken has not been configed ")
		return ""
	}

	// gen signature
	localSignature := signatureGenerate(tk, timestamp, nonce)
	if localSignature != remoteSignature {
		//logs.("signatureGen != signatureIn signatureGen=%s,signatureIn=%s\n", localSignature, remoteSignature)

		//	when verification failed,it should take the blank message like "" as response.
		return ""

	} else {
		//logs.Println("==== pass verification ======")

		//	when verification passed,it should take the echostr as response
		return echostr
	}
}

//	Generate local signature
// 	@token:	user token
//	@timestamp: timestamp of verification request
//	@nonce: nonce of verification request
func signatureGenerate(token, timestamp, nonce string) string {

	//step 1.  sort token、timestamp、nonce
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	//step 2.	joint tree params into a string word
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	//step 3.	product the string key with sha1
	return fmt.Sprintf("%x", s.Sum(nil))
}
