package initialize

import (
	"container/global"
	"github.com/go-ini/ini"
	"log"
	"time"
)

func SetupConfig(configName string) {
	ConfigFile, err := ini.Load(configName)
	if err != nil {
		log.Fatalf("Fail to parse 'config/config.ini': %v", err)
	}

	err = ConfigFile.Section("server").MapTo(&global.CONFIG.Server)
	if err != nil {
		log.Fatalf("Config.MapTo Server err: %v", err)
	}

	global.CONFIG.Server.ReadTimeout = global.CONFIG.Server.ReadTimeout * time.Second
	global.CONFIG.Server.WriteTimeout = global.CONFIG.Server.WriteTimeout * time.Second

	err = ConfigFile.Section("mysql").MapTo(&global.CONFIG.Mysql)
	if err != nil {
		log.Fatalf("Cfg.MapTo Database err: %v", err)
	}
}