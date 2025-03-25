package main

import (
	"deploy_demo/database"
	"deploy_demo/routes"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("starts connect on db..")
	if err := godotenv.Load(); err != nil {
		log.Printf("failed from .env file...%v", err)
	}
	database.InitDB()
	defer database.DB.Close()

	routes.SetHandlers()

}
