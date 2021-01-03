package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:20055")
	if nil != err {
		log.Fatal("failed to connect server")
	}

	defer conn.Close()

	go func(c net.Conn) {
		data := make([]byte, 4096) // 4096 크기의 바이트 슬라이스 생성

		for {
			n, err := c.Read(data) // 서버에서 받은 데이터를 읽음
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("recv data : " + string(data[:n])) // 데이터 출력

			time.Sleep(1 * time.Second)
		}
	}(conn)

	_, _ = fmt.Scanln()
}
