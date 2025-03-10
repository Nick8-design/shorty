package main

import (
	"log"
	"shorty/db"
	"shorty/handlels"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init(){
	err:=godotenv.Load()
	if err!=nil{
		log.Fatal("Error loading env")
	}
go	db.ConnectDb()
go	db.ConnectRedis()
}

func main() {
   

    app := fiber.New()
    app.Post("/shorten", handlels.ShortenURL) // Handle URL shortening
	app.Get("/:short",handlels.RedirectUrl)

    app.Listen(":8080") // Start server
}
