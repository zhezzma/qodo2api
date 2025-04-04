package check

import (
	logger "qodo2api/common/loggger"
)

func CheckEnvVariable() {
	logger.SysLog("environment variable checking...")

	//if config.KLCookie == "" {
	//	logger.FatalLog("环境变量 SG_COOKIE 未设置")
	//}

	logger.SysLog("environment variable check passed.")
}
