package main

import (
	"context"
	"log"
	"pub-sub-go/configs"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var RDS *redis.Client
var CTX context.Context

func init() {
	CTX = context.Background()
	RDS = configs.NewRedisClient()
}

func main() {

	_, err := RDS.Ping(CTX).Result()
	if err != nil {
		log.Println(`failed to connect redis`)
		panic(err)
	}

	publish()

	pubSub := RDS.Subscribe(CTX, `channel-1`)
	defer pubSub.Close()
	for i := 0; i < 10; i++ {
		msg, err := pubSub.ReceiveMessage(CTX)
		if err != nil {
			log.Println(`pubSub.ReceiveMessage(CTX)`, err)
		} else {
			log.Println(`message:`, msg.Payload)
		}
	}
}

func publish() {
	for i := 0; i < 10; i++ {
		err := RDS.Publish(CTX, `channel-1`, `Hello ` + strconv.Itoa(i)).Err()
		if err != nil {
			log.Println(`RDS.Publish()`, err)
		}
	}
}