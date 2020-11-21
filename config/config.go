package config

import "time"

type Config struct {
	Server server
	Mysql  mysql
	Redis redis
}


type server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type mysql struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	MaxIdleConn int
	MaxOpenConn int
}

type redis struct {
	Host string
	Port int
	Password string
	Db int
}