// @title KILO-AI-2API
// @version 1.0.0
// @description KILO-AI-2API
// @BasePath
package main

import (
	"fmt"
	"os"
	"qodo2api/check"
	"qodo2api/common"
	"qodo2api/common/config"
	logger "qodo2api/common/loggger"
	"qodo2api/job"
	"qodo2api/middleware"
	"qodo2api/model"
	"qodo2api/router"
	"strconv"

	"github.com/gin-gonic/gin"
)

//var buildFS embed.FS

func main() {
	logger.SetupLogger()
	logger.SysLog(fmt.Sprintf("qodo2api %s starting...", common.Version))

	check.CheckEnvVariable()

	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	var err error

	model.InitTokenEncoders()
	_, err = config.InitQDCookies()
	if err != nil {
		logger.FatalLog(err)
	}

	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(middleware.RequestId())
	middleware.SetUpLogger(server)

	// 设置API路由
	router.SetApiRouter(server)
	// 设置前端路由
	//router.SetWebRouter(server, buildFS)

	var port = os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(*common.Port)
	}

	if config.DebugEnabled {
		logger.SysLog("running in DEBUG mode.")
	}

	logger.SysLog("qodo2api start success. enjoy it! ^_^\n")
	go job.UpdateCookieTokenTask()

	err = server.Run(":" + port)

	if err != nil {
		logger.FatalLog("failed to start HTTP server: " + err.Error())
	}
}
