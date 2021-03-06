package msgpush

import (
	"github.com/imroc/req"
)

const baseUrl = "https://oapi.dingtalk.com/robot/send?access_token="

type SendTextContent struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

type DingTalk struct {
	Token  string
	ReqUrl string
}

func NewDingTalk(token string) *DingTalk {
	return &DingTalk{Token: token, ReqUrl: baseUrl + token}
}

func (d *DingTalk) Send(content string) error {
	return d.SendText(content)
}

func (d *DingTalk) SendText(content string) error {
	_, err := req.Post(d.ReqUrl, req.BodyJSON(&SendTextContent{
		Msgtype: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
	}))
	return err
}

func (d *DingTalk) String() string {
	return "dingtalk"
}
