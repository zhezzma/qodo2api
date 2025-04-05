package common

import "time"

var StartTime = time.Now().Unix() // unit: second
var Version = "v1.0.1"            // this hard coding will be replaced automatically when building, no need to manually change

type ModelInfo struct {
	Model     string
	MaxTokens int
}

// 创建映射表（假设用 model 名称作为 key）
var ModelRegistry = map[string]ModelInfo{
	"claude-3-7-sonnet": {"claude-3-7-sonnet", 100000},
	"claude-3-5-sonnet": {"claude-3-5-sonnet", 100000},
	"deepseek-r1":       {"deepseek-r1-full", 100000},
	"deepseek-r1-32b":   {"deepseek-r1", 100000},
	"gpt-4o":            {"gpt-4o", 100000},
	"o1":                {"o1", 100000},
	"o3-mini":           {"o3-mini", 100000},
	"o3-mini-high":      {"o3-mini-high", 100000},
	"gemini-2.5-pro":    {"gemini-2.5-pro", 100000},
	"gemini-2.0-flash":  {"gemini-2.0-flash", 100000},
}

// 通过 model 名称查询的方法
func GetModelInfo(modelName string) (ModelInfo, bool) {
	info, exists := ModelRegistry[modelName]
	return info, exists
}

func GetModelList() []string {
	var modelList []string
	for k := range ModelRegistry {
		modelList = append(modelList, k)
	}
	return modelList
}
