package main

import (
	"gallery/database"
	"gallery/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env not found, err:", err)
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("!!! SERVER TERMINATED !!!")
			log.Println("Error:", r)
		}
	}()

	addr := os.Getenv("SRVR_PORT")
	if addr == "" {
		addr = "8082"
	}
	addr = ":" + addr

	os.Mkdir("gallery", 0644)

	log.Printf("Server: http://localhost%s\n", addr)
	database.InitConnect()
	mux := routes.Routes()

	// start server
	log.Println("Server successfully started.")
	panic(http.ListenAndServe(addr, mux))
}
