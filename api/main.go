package main

import ("fmt"
 "github.com/gofiber/fiber/v2"
"log"
"os"
"github.com/joho/godotenv"
"github.com/michaelmounir12/url-shortener/routes"
)

func setupRoutes(app *fiber.App){
	app.Get("/:url",routes.ResolveURL)
	app.Post("/api/v1",routes.ShortenURL)
}

func main() {
	err:= godotenv.Load()
    if err!=nil{
		fmt.Print(err)
	}    
	app:=fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}