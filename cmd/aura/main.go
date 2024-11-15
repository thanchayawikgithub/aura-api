package main

import (
	"aura/internal/config"
	"aura/internal/handler"
	"aura/internal/storage"
)

func init() {
	cfg := config.LoadConfig()

	storage := storage.New(&cfg.Database)
	handler := handler.New(storage, cfg)
}

func main() {

}
