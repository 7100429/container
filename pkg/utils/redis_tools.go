package utils

import (
	"context"
	"encoding/json"
	"gitee.com/chensyf/container/global"
)

func SetResponse(key string, value interface{}) error {
	conn := global.REDIS.GetConn()
	ctx := context.Background()
	marshalValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	conn.Set(ctx, key, marshalValue, 0)
	return nil
}

func GetResponse(key string, response interface{}) bool {
	conn := global.REDIS.GetConn()
	ctx := context.Background()

	if conn.Exists(ctx, key).Val() != 1 {
		return false
	}

	value, err := conn.Get(ctx, key).Bytes()
	if err != nil {
		return false
	}
	err = json.Unmarshal(value, &response)
	if err != nil {
		return false
	}
	return true
}