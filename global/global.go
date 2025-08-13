package global

import (
	"feishuReboot/pkg/model"
	"github.com/jinzhu/gorm"
)

var (
	System           model.SystemConf
	BlockAllRequests bool
	DB               *gorm.DB
)
