package admin

import (
	"net/http"

	"gin-api-mono/internal/code"
	"gin-api-mono/internal/pkg/core"
	"gin-api-mono/internal/pkg/timeutil"
	"gin-api-mono/internal/repository/mysql/models"
)

type listRequest struct {
	Id int32 `form:"id"` // 主键ID
}

type listData struct {
	Id        int32  `json:"id"`         // 主键ID
	Username  string `json:"username"`   // 用户名
	Mobile    string `json:"mobile"`     // 手机号
	CreatedAt string `json:"created_at"` // 创建时间
}

type listResponse struct {
	List []listData `json:"list"`
}

// List 管理员列表
// @Summary 管理员列表
// @Description 管理员列表
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id query int32 false "主键ID"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/admins [get]
func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(listRequest)
		res := new(listResponse)
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

		var resultData []*models.Admin
		dbRequest := h.db.GetDbR().WithContext(ctx.RequestContext()).Where(queryWhere).Find(&resultData)

		if dbRequest.Error != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminListError,
				code.Text(code.AdminListError)).WithError(dbRequest.Error),
			)
			return
		}

		res.List = make([]listData, len(resultData))

		for k, v := range resultData {
			res.List[k] = listData{
				Id:        v.Id,
				Username:  v.Username,
				Mobile:    v.Mobile,
				CreatedAt: v.CreatedAt.Format(timeutil.CSTLayout),
			}
		}

		ctx.Payload(res)
	}
}
