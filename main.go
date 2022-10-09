package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/majidehamide/go-fiber2/database"
	"github.com/majidehamide/go-fiber2/router"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("No config detectted")
	}

	fmt.Println("Start server...")

	_, err = database.ConnectDB()

	if err != nil {
		fmt.Println("DB not connected")
	}

	router.Routes()

}
