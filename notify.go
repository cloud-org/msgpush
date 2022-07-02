package msgpush

import "fmt"

// NotifyPush 通用接口
type NotifyPush interface {
	Send(string) error
	String() string
}

type NotifyPushImpl struct {
	Notifys []NotifyPush
}

func NewNotifyPushImpl() *NotifyPushImpl {
	return &NotifyPushImpl{Notifys: make([]NotifyPush, 0)}
}

func GetNotify(software, token string) (NotifyPush, error) {
	switch software {
	case "dingtalk": // 钉钉推送
		return NewDingTalk(token), nil
	case "wecom": // 企业微信推送
		return NewWeCom(token), nil
	case "slack": // slack 推送
		return NewSlack(token), nil
	case "pushdeer":
		return NewPushDeer(token), nil
	case "feishu":
		return NewFeiShu(token), nil
	default:
		return nil, fmt.Errorf("暂时不支持类型 %v", software)
	}
}

func (n *NotifyPushImpl) AddNotifyPush(software, token string) error {
	push, err := GetNotify(software, token)
	if err != nil {
		return err
	}

	n.Notifys = append(n.Notifys, push)
	return nil
}

func (n *NotifyPushImpl) Push(content string) {
	if n == nil {
		return
	}
	for i := 0; i < len(n.Notifys); i++ {
		if err := n.Notifys[i].Send(content); err != nil {
			continue
		}
	}
}
