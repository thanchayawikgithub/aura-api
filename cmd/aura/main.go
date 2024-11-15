package main

import (
	"aura/internal/config"
	"fmt"
)

func init() {
	cfg := config.LoadConfig()

	fmt.Println(cfg)
}

func main() {

}
