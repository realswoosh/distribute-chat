package net

var HeaderSize int32 = 8

type Packet struct {
	Msg string
	MessageBody []byte
}

func CreatePacket(msg string, body []byte) Packet {
	return Packet {
		Msg: msg,
		MessageBody: body,
	}
}