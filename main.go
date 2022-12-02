package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	name := os.Getenv("NAME")
	email := os.Getenv("EMAIL")

	fmt.Printf("Name : %s\n", name)
	fmt.Printf("Email : %s\n", email)
}
