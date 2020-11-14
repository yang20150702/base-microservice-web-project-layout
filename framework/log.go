package framework

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func GetLogger() *logrus.Logger {
	return log
}

func initLogger() {
	log = logrus.New()
	f, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	log.Out = f
	log.Level = logrus.InfoLevel
}

func logFormat(params gin.LogFormatterParams) string {
	return fmt.Sprintf(
		"%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		params.ClientIP,
		params.TimeStamp.Format(time.RFC1123),
		params.Method,
		params.Path,
		params.Request.Proto,
		params.StatusCode,
		params.Latency,
		params.Request.UserAgent(),
		params.ErrorMessage,
	)
}
