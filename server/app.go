package main

import (
	"./cmd"
	dmnet "./net"
)

func main() {
	s := dmnet.NewServer()

	s.Listen(":10055")
	s.Start()

	cmd.CommandHandler(s)
}