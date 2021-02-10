package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/han-xuefeng/go_gateway/public"
	"time"
)

type AdminInfoOutput struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json:"roles"`
}

type ChangePwdInput struct {
	Password string `form:"password" json:"password" comment:"密码" validate:"required"`
}

func (param *ChangePwdInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
