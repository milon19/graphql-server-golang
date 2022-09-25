package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"simple-graphql-server/internal/app/adapter"
	"simple-graphql-server/internal/app/adapter/controller"
)

func init() {
	initEnv()
}

func initEnv() {
	log.Printf("Loading environment settings.")
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No local env file. Using global OS environment variables, ", err)
	}

}

func main() {
	service := controller.InitializeService()
	adapter.Routes(service)

	fmt.Println("Server is running...")
	err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
