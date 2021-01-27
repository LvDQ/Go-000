package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}

	log.Println("tcp listened!")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}
		// 开始goroutine监听连接
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	// 读写缓冲区
	rd := bufio.NewReader(conn)
	wr := bufio.NewWriter(conn)

	for {
		log.Println("start to receive messages.")
		line, err := rd.ReadString('\n')
		if err != nil {
			log.Printf("read error: %v\n", err)
			return
		}

		log.Println("start to wrighting back.")

		wr.WriteString("hello ")
		wr.WriteString(line)
		wr.Flush() // 一次性syscall
	}
}
