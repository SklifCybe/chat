package main

import (
	"fmt"

	"github.com/SklifCybe/chat/internal/config"
	"github.com/SklifCybe/chat/internal/storage/postgres"
	"github.com/SklifCybe/chat/pkg/logger"
)

func main() {
	fmt.Println("hello")

	cfg := config.MustNew()

	log := logger.New("local")

	log.Info("hello world")

	st, _ := postgres.New(cfg.DatabaseUrl)
	
	defer st.Close()
}
