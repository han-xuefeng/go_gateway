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
	ID          int64    `json:"id" form:"id"`
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

type ServiceDeleteInput struct {
	ID int64 `json:"id" form:"id" comment:"服务ID" example:"56" validate:"required"` //服务ID
}

func (param *ServiceDeleteInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ServiceAddHTTPInput struct {
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名" example:"" validate:"required,valid_service_name"` //服务名
	ServiceDesc string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"" validate:"required,max=255,min=1"`     //服务描述

	RuleType       int    `json:"rule_type" form:"rule_type" comment:"接入类型" example:"" validate:"max=1,min=0"`                           //接入类型
	Rule           string `json:"rule" form:"rule" comment:"接入路径：域名或者前缀" example:"" validate:"required,valid_rule"`                      //域名或者前缀
	NeedHttps      int    `json:"need_https" form:"need_https" comment:"支持https" example:"" validate:"max=1,min=0"`                      //支持https
	NeedStripUri   int    `json:"need_strip_uri" form:"need_strip_uri" comment:"启用strip_uri" example:"" validate:"max=1,min=0"`          //启用strip_uri
	NeedWebsocket  int    `json:"need_websocket" form:"need_websocket" comment:"是否支持websocket" example:"" validate:"max=1,min=0"`        //是否支持websocket
	UrlRewrite     string `json:"url_rewrite" form:"url_rewrite" comment:"url重写功能" example:"" validate:"valid_url_rewrite"`              //url重写功能
	HeaderTransfor string `json:"header_transfor" form:"header_transfor" comment:"header转换" example:"" validate:"valid_header_transfor"` //header转换

	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限" example:"" validate:"max=1,min=0"`                     //关键词
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单ip" example:"" validate:""`                               //黑名单ip
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单ip" example:"" validate:""`                               //白名单ip
	ClientipFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端ip限流	" example:"" validate:"min=0"` //客户端ip限流
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" example:"" validate:"min=0"`          //服务端限流

	RoundType              int    `json:"round_type" form:"round_type" comment:"轮询方式" example:"" validate:"max=3,min=0"`                                //轮询方式
	IpList                 string `json:"ip_list" form:"ip_list" comment:"ip列表" example:"" validate:"required,valid_ipportlist"`                            //ip列表
	WeightList             string `json:"weight_list" form:"weight_list" comment:"权重列表" example:"" validate:"required,valid_weightlist"`               //权重列表
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" comment:"建立连接超时, 单位s" example:"" validate:"min=0"`   //建立连接超时, 单位s
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" comment:"获取header超时, 单位s" example:"" validate:"min=0"` //获取header超时, 单位s
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" comment:"链接最大空闲时间, 单位s" example:"" validate:"min=0"`       //链接最大空闲时间, 单位s
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" comment:"最大空闲链接数" example:"" validate:"min=0"`                     //最大空闲链接数
}

func (params *ServiceAddHTTPInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}