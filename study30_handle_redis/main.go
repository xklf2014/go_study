package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

type MyRedisConnector struct {
	RdbAddr string
}

func readProperties() string {
	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var conn MyRedisConnector
	json.Unmarshal([]byte(byteValue), &conn)
	returnVal := conn.RdbAddr
	return returnVal
}

func initRedisClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     readProperties(),
		Password: "",
		DB:       0,
		PoolSize: 100,
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("redis connect failed by %v\n", err)
		return err
	}
	return nil
}

func main() {

	if err := initRedisClient(); err != nil {
		fmt.Printf("initRedisClient failed by %v\n", err)
		panic(err)
	}

	fmt.Println("redis connect succeed")

	/*	err := rdb.Set(ctx, "mykey", "hello go-redis", 0).Err()
		if err != nil {
			panic(err)
		}*/

	//返回错误信息
	result, err := rdb.Get(ctx, "mykey1").Result()
	if err == redis.Nil {
		fmt.Println("key doesn't exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}

	//不返回错误信息
	r1 := rdb.Get(ctx, "mykey").Val()
	fmt.Println("mykey:", r1)

	hGetAll()
}

func hGetAll() {
	v := rdb.HGetAll(ctx, "user").Val()
	fmt.Println(v)

	v2 := rdb.HGet(ctx, "user", "name").Val()
	fmt.Println(v2)
}
