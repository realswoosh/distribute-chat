package net

import "sync"

type SessionManager struct {
	sessions []*Session
	mutex *sync.Mutex
}

func CreateSessionManager() *SessionManager{
	return &SessionManager{
		mutex: &sync.Mutex{},
	}
}

func (manager* SessionManager) TotalSessionCount() int {
	return len(manager.sessions) + 1;
}

func (manager* SessionManager) Remove(session *Session) {
	manager.mutex.Lock();
	defer manager.mutex.Unlock()

	// 전체를 다 순환하는데..
	// 이건 UUID를 주던가 해서 수정을 해야 될거 같음. (maybe map??)
	for i, check := range manager.sessions {
		if check == session {
			manager.sessions = append(manager.sessions[:i], manager.sessions[i+1:]...)
		}
	}
}