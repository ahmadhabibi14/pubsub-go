package main

import (
	"log"
	"pub-sub-go/configs"
	"pub-sub-go/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	configs.InitRedisClient()

	err := configs.RDS.Ping(configs.RDS_CTX).Err()
	if err != nil {
		log.Println(`failed to connect redis`)
		panic(err)
	}

	engine := html.New(`./views`, `.html`)
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get(`/`, controller.HomePage)
	app.Get(`/notification`, controller.Notification)

	app.Listen(`:8080`)
}