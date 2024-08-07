package controller

import (
	"bufio"
	"fmt"
	"log"
	"pub-sub-go/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func Notification(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		pubSub := configs.RDS.Subscribe(configs.RDS_CTX, configs.REDIS_CHANNEL_NOTIFICATION)
		defer pubSub.Close()

		for {
			msg, err := pubSub.ReceiveMessage(configs.RDS_CTX)
			if err != nil {
				log.Println(`pubSub.ReceiveMessage(CTX)`, err)
			} else {
				log.Println(`message:`, msg.Payload)
			}	
				
			fmt.Fprintf(w, "data: %s\n\n", msg.Payload)

			err = w.Flush()
			if err != nil {
				fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)
				break
			}
		}
	}))

	return nil
}