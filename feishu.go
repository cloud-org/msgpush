package msgpush

import (
	"github.com/imroc/req"
)

type FeiShu struct {
	WebHookUrl string
}

func NewFeiShu(webHookUrl string) *FeiShu {
	return &FeiShu{WebHookUrl: webHookUrl}
}

type FeiShuSendText struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

func (f *FeiShu) Send(content string) error {
	return f.SendText(content)
}

func (f *FeiShu) SendText(content string) error {
	_, err := req.Post(f.WebHookUrl, req.BodyJSON(&FeiShuSendText{
		MsgType: "text",
		Content: struct {
			Text string `json:"text"`
		}{
			Text: content,
		},
	}))
	return err
}

func (f *FeiShu) String() string {
	return "feishu"
}
