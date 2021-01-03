package main

import (
	"distribute-chat/server/cmd"
	dmnet "distribute-chat/server/net"
)

func main() {
	s := dmnet.NewServer()

	_ = s.Listen(":20055")
	s.Start()

	cmd.CommandHandler()
}