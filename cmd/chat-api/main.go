package main

import (
	"fmt"

	"github.com/SklifCybe/chat/pkg/logger"
)

func main() {
	fmt.Println("hello")

	log := logger.New("local")

	log.Info("hello world")

	// todo: add docker
	// todo: init postgresql
	// todo: create router (chi)
}
