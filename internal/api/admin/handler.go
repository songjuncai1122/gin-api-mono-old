package admin

import (
	"gin-api-mono/internal/pkg/core"
	"gin-api-mono/internal/repository/mysql"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 创建管理员
	// @Tags API.admin
	// @Router /api/admin [post]
	Create() core.HandlerFunc

	// List 管理员列表
	// @Tags API.admin
	// @Router /api/admins [get]
	List() core.HandlerFunc

	// First 单个管理员
	// @Tags API.admin
	// @Router /api/admin [get]
	First() core.HandlerFunc

	// Delete 删除管理员
	// @Tags API.admin
	// @Router /api/admin/{id} [delete]
	Delete() core.HandlerFunc
}

type handler struct {
	logger *zap.Logger
	db     mysql.Repo
}

func New(logger *zap.Logger, db mysql.Repo) Handler {
	return &handler{
		logger: logger,
		db:     db,
	}
}

func (h *handler) i() {}
