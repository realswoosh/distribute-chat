package chat

type RoomInfo struct {
	Idx int32
	Title string
	ParticipateCount int32
}

type ISession interface {

}
type Room struct {
	Idx int32
	Title string
	uuid string
	Sessions []ISession
}