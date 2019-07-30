package net

type Packet struct {
	size int32
	msg string
	messageBody []byte
}

func CreatePacket(size int32, msg string, body []byte) Packet {
	return Packet {
		size: size,
		msg: msg,
		messageBody: body,
	}
}