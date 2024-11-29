package router

import (
	"gin-api-mono/internal/api/admin"
	"gin-api-mono/internal/pkg/core"
	"gin-api-mono/internal/repository/mysql"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func NewHTTPMux(logger *zap.Logger, db mysql.Repo) (core.Mux, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	if db == nil {
		return nil, errors.New("db required")
	}

	mux, err := core.New(logger,
		core.WithEnableCors(),
		core.WithEnableSwagger(),
		core.WithEnablePProf(),
	)

	if err != nil {
		panic(err)
	}

	// 示例 CURD，操作 mysql admin 表
	adminHandler := admin.New(logger, db)

	adminRouter := mux.Group("/api")
	{
		// 创建管理员
		adminRouter.POST("/admin", adminHandler.Create())

		// 管理员列表
		adminRouter.GET("/admins", adminHandler.List())

		// 单个管理员
		adminRouter.GET("/admin", adminHandler.First())

		// 删除管理员
		adminRouter.DELETE("/admin/:id", adminHandler.Delete())
	}

	return mux, nil
}
