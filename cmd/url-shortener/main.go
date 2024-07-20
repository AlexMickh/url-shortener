package main

import (
	"fmt"
	"log"

	"github.com/AlexMickh/url-shortener/iternal/config"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
