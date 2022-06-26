<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [msgpush](#msgpush)
  - [usage](#usage)
  - [acknowledgement](#acknowledgement)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### msgpush

通用推送消息 simple sdk

#### usage

```sh
go get -v github.com/cloud-org/msgpush
```

```go
type NotifyPush struct {
	Notifys []msgpush.NotifyPush
}

func GetNotify(software, token string) (msgpush.NotifyPush, error) {
	switch software {
	case "dingtalk": // 钉钉推送
		return msgpush.NewDingTalk(token), nil
	case "wecom": // 企业微信推送
		return msgpush.NewWeCom(token), nil
	case "slack": // slack 推送
		return msgpush.NewSlack(token), nil
	case "pushdeer":
		return msgpush.NewPushDeer(token), nil
	case "feishu":
		return msgpush.NewFeiShu(token), nil
	default:
		return nil, fmt.Errorf("暂时不支持类型 %v", software)
	}
}

func (n *NotifyPush) Push(content string) {
	if n == nil {
		logx.Infof("notify push not init")
		return
	}
	for i := 0; i < len(n.Notifys); i++ {
		if err := n.Notifys[i].Send(content); err != nil {
			logx.Infof("notify err: %v", err)
			continue
		}
	}
}

```

#### acknowledgement

- github.com/imroc/req