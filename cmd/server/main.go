package main

import (
	"fmt"
	"os"

	"github.com/FMR006/redis-go/internal/server"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	s := server.NewServer("", nil)
	if s == nil {
		fmt.Println("Failed to create server")
		os.Exit(1)
	}

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server: ", err.Error())
		os.Exit(1)
	}
}
