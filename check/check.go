package check

import (
	"qodo2api/common/config"
	logger "qodo2api/common/loggger"
)

func CheckEnvVariable() {
	logger.SysLog("environment variable checking...")

	if config.QDCookie == "" {
		logger.FatalLog("环境变量 QD_COOKIE 未设置")
	}

	logger.SysLog("environment variable check passed.")
}
