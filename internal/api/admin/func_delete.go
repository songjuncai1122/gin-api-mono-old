package admin

import (
	"net/http"

	"gin-api-mono/internal/code"
	"gin-api-mono/internal/pkg/core"
	"gin-api-mono/internal/repository/mysql/models"
)

type deleteRequest struct {
	Id int32 `uri:"id"` // 主键ID
}

type deleteResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Delete 删除管理员
// @Summary 删除管理员
// @Description 删除管理员
// @Tags API.admin
// @Accept json
// @Produce json
// @Param id path int32 true "主键ID"
// @Success 200 {object} deleteResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/{id} [delete]
func (h *handler) Delete() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(deleteRequest)
		res := new(deleteResponse)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		deleteWhere := new(models.Admin)
		deleteWhere.Id = req.Id

		dbResult := h.db.GetDbW().WithContext(ctx.RequestContext()).Delete(deleteWhere)

		if dbResult.Error != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminDeleteError,
				code.Text(code.AdminDeleteError)).WithError(dbResult.Error),
			)
			return
		}

		res.Id = req.Id
		ctx.Payload(res)
	}
}
