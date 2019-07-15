package chat

import (
	dmnet "../net"
)

type Room struct {
	uuid string
	sessions []*dmnet.Session
}