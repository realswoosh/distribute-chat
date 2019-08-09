package handler

import (
	"../../../proto-gen-go"
)

type MsgFunc func(msgBody []byte)

type MsgHandler struct {
	Dispatcher map[netmsg.Message_Type]MsgFunc
}
