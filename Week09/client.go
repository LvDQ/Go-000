package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	clientReader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)


	go func() {
		for {
			// Waiting for the server response
			serverResponse, err := serverReader.ReadString('\n')

			if err != nil {
				log.Fatalf("server error: %v\n", err)
				return
			}
			log.Println(strings.TrimSpace(serverResponse))
		}
	}()

	for {
		// Waiting for the client request
		clientRequest, err := clientReader.ReadString('\n')

		if err == io.EOF {
			log.Println("client closed the connection")
			return
		}

		clientRequest = strings.TrimSpace(clientRequest)
		if _, err = conn.Write([]byte(clientRequest + "\n")); err != nil {
			log.Printf("failed to send the client request: %v\n", err)
		}
	}
}
