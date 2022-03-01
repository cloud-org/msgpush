package msgpush

import (
	"testing"
)

func TestSlack_Send(t *testing.T) {
	s := NewSlack("your slack webhook url")
	s.Send("今天是个好日子")
}
