package msgpush

import (
	"fmt"
	"net/url"

	"github.com/imroc/req"
)

type PushDeer struct {
	Token string
}

func NewPushDeer(token string) *PushDeer {
	return &PushDeer{Token: token}
}

func (p *PushDeer) Send(content string) error {
	return p.SendMd(content)
}

func (p *PushDeer) SendMd(content string) error {
	// https://api2.pushdeer.com/message/push?pushkey=<key>&text=标题&desp=<markdown>&type=markdown
	// urlencode
	title := "msg from pushdeer"
	reqUrl := fmt.Sprintf("https://api2.pushdeer.com/message/push?pushkey=%s&text=%s&desp=%s&type=markdown",
		p.Token, url.QueryEscape(title), url.QueryEscape(content))
	//log.Println("url is", reqUrl)
	_, err := req.Post(reqUrl)
	return err
}

func (p *PushDeer) String() string {
	return "pushdeer"
}
