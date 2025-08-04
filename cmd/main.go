package main

import (
	"fmt"
	"wakeapi/internal/server"
)

func main() {
	err := server.Start()
	if err != nil {
		panic(fmt.Errorf("failed to start server: %v", err.Error()))
	}
}
