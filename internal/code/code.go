package code

import (
	_ "embed"

	"gin-api-mono/configs"
)

//go:embed code.go
var ByteCodeFile []byte

// Failure 错误时返回结构
type Failure struct {
	Code    int    `json:"code"`    // 业务码
	Message string `json:"message"` // 描述信息
}

const (
	ServerError        = 10101
	ParamBindError     = 10102
	JWTAuthVerifyError = 10103

	AdminCreateError = 20201
	AdminListError   = 20202
	AdminFirstError  = 20203
	AdminDeleteError = 20204
)

func Text(code int) string {
	lang := configs.Get().Language.Local

	if lang == configs.ZhCN {
		return zhCNText[code]
	}

	if lang == configs.EnUS {
		return enUSText[code]
	}

	return zhCNText[code]
}
