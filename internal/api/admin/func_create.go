package admin

import (
	"net/http"

	"gin-api-mono/internal/code"
	"gin-api-mono/internal/pkg/core"
	"gin-api-mono/internal/pkg/validation"
	"gin-api-mono/internal/repository/mysql/models"
)

type createRequest struct {
	Username string `form:"username" binding:"required"` // 用户名
	Mobile   string `form:"mobile" binding:"required"`   // 手机号
}

type createResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Create 创建管理员
// @Summary 创建管理员
// @Description 创建管理员
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "用户名"
// @Param mobile formData string true "手机号"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin [post]
func (h *handler) Create() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}

		createData := new(models.Admin)
		createData.Username = req.Username
		createData.Mobile = req.Mobile

		dbResult := h.db.GetDbW().WithContext(ctx.RequestContext()).Create(createData)

		if dbResult.Error != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminCreateError,
				code.Text(code.AdminCreateError)).WithError(dbResult.Error),
			)
			return
		}

		res.Id = createData.Id
		ctx.Payload(res)
	}
}
