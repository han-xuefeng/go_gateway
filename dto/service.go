package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/han-xuefeng/go_gateway/public"
)

type ServiceListInput struct {
	Info     string `form:"info" json:"info" comment:"关键词" example:"" validate:""`
	PageNo   int    `form:"page_no" json:"page_no" comment:"页数" example:"1" validate:"required"`
	PageSize int    `form:"page_size" json:"page_size" comment:"每页条数" example:"20" validate:"required"`
}

func (param *ServiceListInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ServiceListItemOutput struct {
	Id          int    `json:"id" form:"id"`
	LoadType    int    `json:"load_type" `
	ServiceName string `json:"salt" `
	ServiceDesc string `json:"service_desc" `
	ServiceAddr string `json:"service_addr" `
	Qpd         int64  `json:"qpd" `
	Qps         int64  `json:"qps" `
	TotalNode   int    `json:"total_node"`
}

type ServiceListOutput struct {
	Total int64                   `json:"total" form:"total" comment:"总数" example:"" validate:""` //总数
	List  []ServiceListItemOutput `json:"list" form:"list" comment:"列表" example:"" validate:""`   //列表
}
