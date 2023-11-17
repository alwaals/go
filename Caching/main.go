package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	//Redis - Connect to Redis server

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis address
		Password: "",       // Redis Password
		DB:       0,
	})

	//ctx := context.Background()

	json, err := json.Marshal(User{Username: "Saurav", Password: "*******"})
	if err != nil {
		fmt.Println(err)
	}

	//Redis Set - set data in Redis as a key pair

	err = client.Set("id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	//Redis Get - get data in Redis

	val, err := client.Get("id1234").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("User Details =============+>", val)
}
