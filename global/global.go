package global

import (
	"go-web-template/pkg/model"
	"gorm.io/gorm"
)

var (
	System           model.SystemConf
	BlockAllRequests bool
	DB               *gorm.DB
)
