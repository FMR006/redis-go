package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	mu := sync.RWMutex{}
	storage := make(map[string]string)
	expireAt := make(map[string]time.Time)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConn(conn, storage, expireAt, &mu)
	}
}
