package controller

import (
	"fmt"
	"log"
	"pub-sub-go/configs"

	"github.com/gofiber/fiber/v2"
)

func HomePage(c *fiber.Ctx) error {
	message := fmt.Sprintf("User dengan IP %v mengakses website", c.IP())
	
	err := configs.RDS.Publish(
		configs.RDS_CTX,
		configs.REDIS_CHANNEL_NOTIFICATION,
		message,
	).Err()

	if err != nil {
		log.Println(`Error:`, err)
	}

	return c.Render(`index`, fiber.Map{})
}