package main

import (
	"hack/client"
	"hack/iternal/app/apiserver"
	"log"
	//"github.com/redis/go-redis/v9"
)

func main() {
	s := apiserver.New(":8080")
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
	if err := client.Start(":8088"); err != nil {
		log.Fatal(err)
	}
	// client1 := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// 	Password:"",
	// 	DB: 0,
	// })

	select {}
}
