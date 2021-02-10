package controller

import (
	"encoding/json"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/han-xuefeng/go_gateway/dao"
	"github.com/han-xuefeng/go_gateway/dto"
	"github.com/han-xuefeng/go_gateway/middleware"
	"github.com/han-xuefeng/go_gateway/public"
)

type AdminController struct {
}

func AdminRegister(group *gin.RouterGroup) {
	admin := &AdminController{}
	group.GET("/admin_info", admin.AdminInfo)
	group.POST("/change_pwd", admin.ChangePwd)
}

// AdminInfo godoc
// @Summary 管理员信息
// @Description 管理员信息
// @Tags 管理员接口
// @ID /admin/admin_info
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/admin_info [get]
func (admin *AdminController) AdminInfo(c *gin.Context) {
	//取session
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(sessInfo.(string)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	out := &dto.AdminInfoOutput{
		ID:           adminSessionInfo.ID,
		Name:         adminSessionInfo.UserName,
		LoginTime:    adminSessionInfo.LoginTime,
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Introduction: "I am a super administrator",
		Roles:        []string{"admin"},
	}
	middleware.ResponseSuccess(c, out)
}

// AdminInfo godoc
// @Summary 修改密码
// @Description 修改密码
// @Tags 管理员接口
// @ID /admin/change_pwd
// @Accept  json
// @Produce  json
// @Param body body dto.ChangePwdInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin/change_pwd [post]
func (admin *AdminController) ChangePwd(c *gin.Context) {
	params := &dto.ChangePwdInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//获取session
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(sessInfo.(string)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	//查询数据库是否存在用户
	adminDao := &dao.Admin{Id: adminSessionInfo.ID}
	adminDao, err = adminDao.Find(c, tx, adminDao)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	//用户存在更改密码
	saltPassword := public.GenSaltPassword(adminDao.Salt, params.Password)
	adminDao.Password = saltPassword
	if err = adminDao.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}
