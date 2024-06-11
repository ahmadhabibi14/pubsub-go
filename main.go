package main

import "pub-sub-go/configs"

func main() {
	RDS := configs.NewRedisClient()
	_ = RDS
}