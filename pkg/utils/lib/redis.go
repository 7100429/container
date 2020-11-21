package lib

import (
	"github.com/go-redis/redis/v8"
)

type MRedis struct {
	conn *redis.Client
	host string
	port int
	db int
	password string
}


func NewMRedis(host string, port int, db int, password string) *MRedis {
	return &MRedis{
		host: host,
		port: port,
		db: db,
		password: password,
	}
}


func (r *MRedis) GetConn() *redis.Client {
	r.conn = redis.NewClient(&redis.Options{
		Addr:     "10.10.10.111:6379",
		Password: "goodsang123", // no password set
		DB:       0,  // use default DB
	})
	return r.conn
}


func (r *MRedis) Close() {
	r.conn.Close()
}
