package main

import (
	"apipos/database"
	"apipos/routing"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8023,http://localhost,https://cms.jaygeegroupapp.com",
		AllowHeaders: "Origin,Content-type,Accept",
		AllowMethods: "GET,POST,PUT,OPTIONS",
	}))
	database.Conndbisoide()
	database.Connmaster()
	database.Conndbnahm()
	database.Conndbevent()
	routing.Routes(app)
	log.Fatal(app.Listen(":8022"))
}
