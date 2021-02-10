package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/han-xuefeng/go_gateway/public"
	"time"
)

type AdminSessionInfo struct {
	ID        int       `json:"id"`
	UserName  string    `json:"username"`
	LoginTime time.Time `json:"login_time"`
}

type AdminLoginInput struct {
	Username string `form:"username" json:"username" comment:"用户名"  validate:"required" example:"admin"`
	Password string `form:"password" json:"password" comment:"密码"   validate:"required" example:"123456"`
}

func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type AdminLoginOutput struct {
	Token string `form:"token" json:"token" comment:"token"`
}
