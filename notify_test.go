package msgpush

import "testing"

func TestNotifyPushImpl_Push(t *testing.T) {
	np := NewNotifyPushImpl()
	token := "slack token"
	t.Log(np.AddNotifyPush("slack", token))
	np.Push("msgpush test")
}
