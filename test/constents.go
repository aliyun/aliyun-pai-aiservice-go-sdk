package test

import "os"

var (
	//host  = "http://ai-service.ce8cc13b6421545749e7b4605f3d02607.cn-hangzhou.alicontainer.com"
	host  = "http://aiservice.c6e65da2319794458890a0c6320f260e3.cn-beijing.alicontainer.com"
	appId = os.Getenv("APPID")
	token = os.Getenv("TOKEN")
)
