package net

func byteToInt32(data []byte) int32{
	var n int32

	n |= int32(data[0])
	n |= int32(data[1]) << 8
	n |= int32(data[2]) << 16
	n |= int32(data[3]) << 24

	return n
}
