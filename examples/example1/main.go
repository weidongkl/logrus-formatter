package main

import (
	"gitee.com/weidongkl/logrus-formatter/examples/example1/logger"
	"gitee.com/weidongkl/logrus-formatter/examples/example1/other"
)

func main() {
	logger.Log.Infoln("test messages")
	other.Other()
}
