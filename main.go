package main

import (
	"fmt"
	"leave-management/database"

	// "leave-management/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Backend is running")
}


func main() {
	database.Connect()

	http.HandleFunc("/",homeHandler)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	// router := routes.RegisterRoutes()

	fmt.Println("Server is running on port",port)
	// log.Fatal(http.ListenAndServe(":"+port,router))
}