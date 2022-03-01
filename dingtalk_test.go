package msgpush

import (
	"testing"
)

func TestDingTalk_Send(t *testing.T) {
	d := NewDingTalk("you dingtalk bot token")
	d.Send("今天是个好日子")
}
