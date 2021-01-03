package handler

import (
	pb "distribute-chat/proto-gen-go/netmsg/pb"
)

type MsgFunc func(msgBody []byte)

type MsgHandler struct {
	Dispatcher map[pb.Message_Type]MsgFunc
}
