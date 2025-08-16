package router

import (
	"go-web-template/router/system"
)

type Group struct {
	System system.RouterGroup
}

var GroupApp = new(Group)
