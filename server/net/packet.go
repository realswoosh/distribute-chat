package net

var HeaderSize int32 = 8

type Packet struct {
	Size uint32
	Msg string
	MessageBody []byte
}

func CreatePacket(msg string, body []byte) Packet {
	return Packet {
		Size: uint32(len(body)),
		Msg: msg,
		MessageBody: body,
	}
}