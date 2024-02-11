package main

import (
	"fmt"

	"github.com/SklifCybe/chat/pkg/logger"
)

func main() {
	fmt.Println("hello")

	log := logger.New("local")

	log.Info("hello world")
	// todo: create router (chi)
}
