package main

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"golang.org/x/net/context"
	"time"
)

// TODO：单元测试

func main() {
	err := initClient()
	if err != nil {
		fmt.Println(err)
	}
}

// 初始化连接
func initClient() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "root", // no password set
		DB:       0,      // use default DB
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
