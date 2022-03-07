package msgpush

import "testing"

func TestFeiShu_Send(t *testing.T) {
	f := NewFeiShu("your feishu bot webhook url")
	f.Send("今天是个好日子")
}
