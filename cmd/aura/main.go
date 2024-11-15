package main

import (
	"aura/internal/config"
	"aura/internal/storage"
)

func init() {
	cfg := config.LoadConfig()

	storage := storage.New(&cfg.Database)
	_ = storage
}

func main() {

}
