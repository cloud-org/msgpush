package msgpush

import (
	"fmt"

	"github.com/imroc/req"
)

type slackTextContent struct {
	Text string `json:"text"`
}

type Slack struct {
	ReqUrl string
	Pre    string
}

func NewSlack(token string) *Slack {
	return &Slack{ReqUrl: token}
}

func (s *Slack) Send(content string) error {
	return s.sendText(content)
}

func (s *Slack) SendMrDown(title, content string) error {
	msg := fmt.Sprintf(`{
	"blocks": [
		{
			"type": "header",
			"text": {
				"type": "plain_text",
				"text": "%s",
				"emoji": true
			}
		},
		{
			"type": "section",
			"text": {
				"type": "plain_text",
				"text": "%s",
				"emoji": true
			}
		}
	]
}`, title, content)
	_, err := req.Post(s.ReqUrl, req.BodyJSON(msg))
	return err
}

func (s *Slack) sendText(content string) error {
	_, err := req.Post(s.ReqUrl, req.BodyJSON(&slackTextContent{Text: content}))
	return err
}

func (s *Slack) String() string {
	return "slack"
}
