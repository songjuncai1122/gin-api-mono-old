package admin

import (
	"net/http"

	"gin-api-mono/internal/code"
	"gin-api-mono/internal/pkg/core"
	"gin-api-mono/internal/pkg/timeutil"
	"gin-api-mono/internal/repository/mysql/models"

	"gorm.io/gorm"
)

type firstRequest struct {
	Id int32 `form:"id"` // 主键ID
}

type firstResponse struct {
	Id        int32  `json:"id"`         // 主键ID
	Username  string `json:"username"`   // 用户名
	Mobile    string `json:"mobile"`     // 手机号
	CreatedAt string `json:"created_at"` // 创建时间
}

// First 单个管理员
// @Summary 单个管理员
// @Description 单个管理员
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id query int32 false "主键ID"
// @Success 200 {object} firstResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin [get]
func (h *handler) First() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(firstRequest)
		res := new(firstResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		queryWhere := new(models.Admin)
		if req.Id != 0 {
			queryWhere.Id = req.Id
		}

		var resultData *models.Admin
		dbRequest := h.db.GetDbR().WithContext(ctx.RequestContext()).Where(queryWhere).First(&resultData)

		if dbRequest.Error != nil && dbRequest.Error != gorm.ErrRecordNotFound {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminFirstError,
				code.Text(code.AdminFirstError)).WithError(dbRequest.Error),
			)
			return
		}

		res.Id = resultData.Id
		res.Username = resultData.Username
		res.Mobile = resultData.Mobile
		res.CreatedAt = resultData.CreatedAt.Format(timeutil.CSTLayout)

		ctx.Payload(res)
	}
}
