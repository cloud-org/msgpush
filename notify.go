package msgpush

// NotifyPush 通用接口
type NotifyPush interface {
	Send(string) error
	String() string
}
