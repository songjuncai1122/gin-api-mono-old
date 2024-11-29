package code

var zhCNText = map[int]string{
	ServerError:        "内部服务器错误",
	ParamBindError:     "参数信息错误",
	JWTAuthVerifyError: "JWT 授权验证错误",

	AdminCreateError: "创建管理员失败",
	AdminListError:   "获取管理员列表失败",
	AdminFirstError:  "获取单个管理员失败",
	AdminDeleteError: "删除管理员失败",
}
