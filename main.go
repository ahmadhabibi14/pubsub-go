package main

import (
	"fmt"
	"log"
	"pub-sub-go/configs"
	"pub-sub-go/controller"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

const APP_NAME = `Notification System from Redis Pub/Sub`

func main() {
	configs.InitRedisClient()

	err := configs.RDS.Ping(configs.RDS_CTX).Err()
	if err != nil {
		log.Println(`failed to connect redis`)
		panic(err)
	}

	go publishNotification()

	engine := html.New(`./views`, `.html`)
	app := fiber.New(fiber.Config{
		AppName: APP_NAME,
		Views: engine,
	})
	app.Get(`/`, func(c *fiber.Ctx) error {
		return c.Render(`index`, fiber.Map{
			`title`: APP_NAME,
		})
	})
	app.Get(`/notification`, controller.Notification)

	app.Listen(`:8080`)
}

func publishNotification() {
	var i int
	for {
		i++

		msg := fmt.Sprintf("%d - The time is: '%v'", i, time.Now())
		if err := configs.RDS.Publish(configs.RDS_CTX, configs.CH_REDIS_PREFIX+`notification`, msg).Err(); err != nil {
			log.Println(`Error:`, err)
		}

		time.Sleep(2 * time.Second)
	}
}