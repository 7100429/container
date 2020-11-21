package initialize

import (
	"gitee.com/chensyf/container/global"
	"gitee.com/chensyf/container/pkg/utils/lib"
)


func SetupRedis() {
	redisCfg := global.CONFIG.Redis
	global.REDIS = lib.NewMRedis(redisCfg.Host, redisCfg.Port, redisCfg.Db, redisCfg.Password)
}