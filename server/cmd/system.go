package cmd

import (
	dmnet "../net"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func CommandHandler(s *dmnet.Server) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ");
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("exit", text) == 0 {
			log.Print("exit server")
			s.Stop();
			break;
		}
	}
}