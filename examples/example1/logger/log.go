package logger

import (
	formatter "gitee.com/weidongkl/logrus-formatter"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.SetReportCaller(true)
	Log.SetFormatter(&formatter.Formatter{})
}
