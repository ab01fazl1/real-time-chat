package main

import (
	"log"

	"github.com/ab01fazl1/real-time-chat/cmd"
	"github.com/joho/godotenv"
)

func main() {

	// env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	
	cmd.Main()
}
