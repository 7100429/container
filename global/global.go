package global

import (
	"container/config"
	"github.com/jinzhu/gorm"
)

var (
	DB     *gorm.DB      // 数据库orm
	CONFIG config.Config // 配置文件
)
