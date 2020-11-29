package database

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func IniStorage() *redis.Client {
	addr := fmt.Sprintf("%s:%s", os.Getenv("stg.Host"), os.Getenv("stg.Port"))
	fmt.Println(addr)
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}
