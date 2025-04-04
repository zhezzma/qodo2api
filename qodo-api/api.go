package qodo_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"qodo2api/common/config"
	logger "qodo2api/common/loggger"
	"qodo2api/cycletls"
	"strings"
)

const (
	baseURL      = "https://api.gen.qodo.ai"
	chatEndpoint = baseURL + "/v2/chats/chat"
)

func MakeStreamChatRequest(c *gin.Context, client cycletls.CycleTLS, jsonData []byte, cookie string) (<-chan cycletls.SSEResponse, error) {
	split := strings.Split(cookie, "=")
	tokenInfo, ok := config.QDTokenMap[split[0]]
	if !ok {
		return nil, fmt.Errorf("cookie not found in QDTokenMap")
	}

	options := cycletls.Options{
		Timeout: 10 * 60 * 60,
		Proxy:   config.ProxyUrl, // 在每个请求中设置代理
		Body:    string(jsonData),
		Method:  "POST",
		Headers: map[string]string{
			"User-Agent":      "axios/1.7.9",
			"Connection":      "close",
			"Host":            "api.gen.qodo.ai",
			"Accept":          "text/plain",
			"Accept-Encoding": "gzip, compress, deflate, br",
			"Content-Type":    "application/json",
			"Request-id":      uuid.New().String(),
			"Authorization":   `Bearer ` + tokenInfo.AccessToken,
		},
	}

	logger.Debug(c.Request.Context(), fmt.Sprintf("cookie: %v", cookie))

	sseChan, err := client.DoSSE(chatEndpoint, options, "POST")
	if err != nil {
		logger.Errorf(c, "Failed to make stream request: %v", err)
		return nil, fmt.Errorf("Failed to make stream request: %v", err)
	}
	return sseChan, nil
}
