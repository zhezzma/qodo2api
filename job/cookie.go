package job

import (
	"fmt"
	"github.com/deanxv/CycleTLS/cycletls"
	"qodo2api/common/config"
	logger "qodo2api/common/loggger"
	google_api "qodo2api/google-api"
	"time"
)

func UpdateCookieTokenTask() {
	client := cycletls.Init()
	defer safeClose(client)
	for {
		logger.SysLog("qodo2api Scheduled UpdateCookieCreditTask Task Job Start!")

		for _, cookie := range config.NewCookieManager().Cookies {
			tokenInfo, ok := config.QDTokenMap[cookie]
			if ok {
				request := google_api.RefreshTokenRequest{
					Key:          cookie,
					RefreshToken: tokenInfo.RefreshToken,
				}
				token, err := google_api.GetFirebaseToken(request)
				if err != nil {
					logger.SysError(fmt.Sprintf("GetFirebaseToken err: %v Req: %v", err, request))
				} else {
					config.QDTokenMap[cookie] = config.QDTokenInfo{
						ApiKey:       cookie,
						RefreshToken: token.RefreshToken,
						AccessToken:  token.AccessToken,
					}
				}
			}

		}

		logger.SysLog("qodo2api Scheduled UpdateCookieCreditTask Task Job End!")

		now := time.Now()
		remainder := now.Minute() % 10
		minutesToAdd := 10 - remainder
		if remainder == 0 {
			minutesToAdd = 10
		}
		next := now.Add(time.Duration(minutesToAdd) * time.Minute)
		next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), 0, 0, next.Location())
		time.Sleep(next.Sub(now))
	}
}
func safeClose(client cycletls.CycleTLS) {
	if client.ReqChan != nil {
		close(client.ReqChan)
	}
	if client.RespChan != nil {
		close(client.RespChan)
	}
}
