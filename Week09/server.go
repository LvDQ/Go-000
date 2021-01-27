package main

import (
	"bufio"
	"context"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}

	ctx, _ := context.WithCancel(context.Background())

	log.Println("tcp listened!")
	var BusChan = make(chan string)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}

		go handleConn(ctx, conn, BusChan)
	}
}

func handleConn(ctx context.Context, conn net.Conn, BusChan chan string) {
	defer conn.Close()
	// 读写缓冲区
	reader := bufio.NewReader(conn)
	//writer := bufio.NewWriter(conn)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go ListenChanAndWriteBack(ctx, conn, BusChan)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("read error: %v\n", err)
			//退出时注销对Buschan的监听
			cancel()
			return
		}
		
		BusChan <- line
	}
}

func ListenChanAndWriteBack(ctx context.Context, conn net.Conn, ch <-chan string) {
	writer := bufio.NewWriter(conn)
	log.Println("start to writing back.")
	//题目有点模糊，暂时多goroutine抢占该channel来做纯粹功能实现，如若实现广播等IM功能
	//可用链表，hashmap等结构注册存储conn等信息，再做分发
	for{
		select{
		case msg:=<-ch:
			log.Println("Writing  Message: ", msg)
			writer.WriteString(msg)
			writer.Flush()
		case <-ctx.Done():
			log.Println("Done ")
			return
		}
	}

}
