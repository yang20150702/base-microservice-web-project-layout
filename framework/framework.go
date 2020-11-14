package framework

import (
	"github.com/gin-gonic/gin"
)

// GlobalEngine 全局的路由engine
var GlobalEngine *gin.Engine

func init() {
	initLogger()

	GlobalEngine = gin.New()
	// 把gin的日志输出到logrus中
	gin.DefaultWriter = GetLogger().Writer()
	GlobalEngine.Use(gin.LoggerWithFormatter(logFormat))
	GlobalEngine.Use(gin.Recovery())
}
