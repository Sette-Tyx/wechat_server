package process

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// WXCheckSignature 微信接入校验
// func WXCheckSignature(c *gin.Context) {
// 	signature := c.Query("signature")
// 	timestamp := c.Query("timestamp")
// 	nonce := c.Query("nonce")
// 	echostr := c.Query("echostr")

// 	ok := util.CheckSignature(signature, timestamp, nonce, Token)
// 	if !ok {
// 		log.Println("[微信接入] - 微信公众号接入校验失败!")
// 		return
// 	}

// 	log.Println("[微信接入] - 微信公众号接入校验成功!")
// 	_, _ = c.Writer.WriteString(echostr)
// }

// WXTextMsg 微信文本消息结构体
type WXTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	MsgId        int64
}

// WXMsgReceive 微信消息接收
func WXMsgReceive(c *gin.Context) {
	var textMsg WXTextMsg
	err := c.ShouldBindXML(&textMsg)
	if err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		return
	}
	// body, _ := ioutil.ReadAll(c.Request)
	// if body != nil {
	// 	fmt.Printf("请求body内容为:%s\n", body)
	// }
	fmt.Printf("请求body内容为:%s\n", c.Request.RequestURI)

	log.Printf("[消息接收] - 收到消息, 消息类型为: %s, 消息内容为: %s\n", textMsg.MsgType, textMsg.Content)

	classification(c, textMsg)
	// WXMsgReply(c, textMsg.ToUserName, textMsg.FromUserName)
}

func classification(c *gin.Context, textMsg WXTextMsg) {
	if textMsg.Content == "北京" {
		repTextMsg := WXRepTextMsg{
			ToUserName:   textMsg.FromUserName,
			FromUserName: textMsg.ToUserName,
			CreateTime:   time.Now().Unix(),
			MsgType:      "text",
			// Content:      fmt.Sprintf("[消息回复] - %s", "北京天气很好"),
			Content: "hello",
		}

		msg, err := xml.Marshal(&repTextMsg)
		if err != nil {
			log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
			return
		}
		_, _ = c.Writer.Write(msg)
	}
}

// WXRepTextMsg 微信回复文本消息结构体
type WXRepTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}

// WXMsgReply 微信消息回复
func WXMsgReply(c *gin.Context, fromUser, toUser string) {
	repTextMsg := WXRepTextMsg{
		ToUserName:   toUser,
		FromUserName: fromUser,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      fmt.Sprintf("[消息回复] - %s", time.Now().Format("2006-01-02 15:04:05")),
	}

	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
}
