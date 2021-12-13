package xrouter

type RouterTag uint8

func (r RouterTag) String() string {
	if r > defineStart && r < defineEnd {
		return RTag[r]
	}
	return ""
}

const (
	defineStart RouterTag = iota
	// INDENT 字面量
	INDENT //indent
	// PATH 参数在path中
	PATH // path
	defineEnd
)

var RTag = [...]string{
	INDENT: "indent",
	PATH:   "path",
}
