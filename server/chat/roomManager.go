package chat

import (
	"sync"
)

type RoomManager struct {
	Rooms []Room
}

var instance *RoomManager
var once sync.Once
var monitor sync.Mutex

func RoomManagerInstance() *RoomManager {
	once.Do(func() {
		instance = new (RoomManager)
	})
	return instance
}

func (manage* RoomManager)Add(newRoom Room) {
	monitor.Lock()
	defer monitor.Unlock()
	manage.Rooms = append(manage.Rooms, newRoom)
}

func (manage* RoomManager)Exists(title string) (bool, int) {
	for i, room := range manage.Rooms {
		if room.Title == title {
			return true, i
		}
	}

	return false, -1
}

func (manage* RoomManager)List() []RoomInfo {
	var roomInfos []RoomInfo
	monitor.Lock()
	for _, e := range manage.Rooms {
		ri := RoomInfo {
			e.Idx,
			e.Title,
			int32(len(e.Sessions)),
		}
		roomInfos = append(roomInfos, ri)
	}
	monitor.Unlock()
	return roomInfos
}