package router

import (
	"feishuReboot/router/system"
)

type Group struct {
	System system.RouterGroup
}

var GroupApp = new(Group)
