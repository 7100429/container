package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.10.10.111:6379",
		Password: "goodsang123", // no password set
		DB:       0,  // use default DB
	})

	defer rdb.Close()
	ctx := context.Background()
	// String 字符串
	fmt.Println("String")
	err := rdb.Set(ctx, "chensy", "goodsang123", 0).Err()
	if err != nil {
		panic(err)
	}

	valBool := rdb.Exists(ctx, "chensy").Val()
	fmt.Println(valBool)

	val, err := rdb.Get(ctx, "chensy").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("set 'chensy' get: " + val)

	rdb.Del(ctx, "chensy")


	// Hash 哈希
	fmt.Println("Hash")
	err = rdb.HSet(ctx, "student", "chensy", "22", "chensy2", "21").Err()
	if err != nil {
		panic(err)
	}

	result, err := rdb.HGetAll(ctx, "student").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	result_str, err := rdb.HGet(ctx, "student", "chensy").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result_str)

	rdb.Del(ctx, "student")

	// List列表
	fmt.Println("List")
	rdb.LPush(ctx, "chensy",  "20")
	rdb.LPush(ctx, "chensy", "21")
	rdb.LPush(ctx, "chensy",  "22")

	resultList := rdb.LRange(ctx, "chensy", 0, 10).Val()
	fmt.Println(resultList)

	rdb.Del(ctx, "chensy")


	// Set集合
	fmt.Println("Set")

	rdb.SAdd(ctx, "chensy", "cjp")
	rdb.SAdd(ctx, "chensy", "txp")
	rdb.SAdd(ctx, "chensy", "txp")
	rdb.SAdd(ctx, "chensy", "zxl")
	rdb.SAdd(ctx, "chensy", "zxl")

	fmt.Println(rdb.SMembers(ctx, "chensy"))

	rdb.Del(ctx, "chensy")

	// Sorted Set 有序集合
	fmt.Println("Sorted Set")
	rdb.ZAdd(ctx, "chensy", &redis.Z{Score: 1, Member: "cjp"})
	rdb.ZAdd(ctx, "chensy", &redis.Z{Score: 2, Member: "zxl"})
	rdb.ZAdd(ctx, "chensy", &redis.Z{Score: 3, Member: "zzq"})
	rdb.ZAdd(ctx, "chensy", &redis.Z{Score: 4, Member: "txp"})
	rdb.ZAdd(ctx, "chensy", &redis.Z{Score: 5, Member: "zzz"})

	fmt.Println(rdb.ZRange(ctx, "chensy", 0, 10))
}
