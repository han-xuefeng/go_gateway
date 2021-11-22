package controller

import (
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/han-xuefeng/go_gateway/dao"
	"github.com/han-xuefeng/go_gateway/dto"
	"github.com/han-xuefeng/go_gateway/middleware"
	"github.com/han-xuefeng/go_gateway/public"
)

type ServiceController struct {
}

func ServiceRegister(group *gin.RouterGroup) {
	serviceController := &ServiceController{}
	group.GET("service_list", serviceController.ServiceList)
	group.GET("service_delete", serviceController.ServiceDelete)
}

// ServiceList godoc
// @Summary 服务列表
// @Description 服务列表
// @Tags 服务管理
// @ID /service/service_list
// @Accept  json
// @Produce  json
// @Param info query string false "关键词"
// @Param page_size query int true "每页个数"
// @Param page_no query int true "当前页数"
// @Success 200 {object} middleware.Response{data=dto.ServiceListOutput} "success"
// @Router /service/service_list [get]
func (service *ServiceController) ServiceList(c *gin.Context) {
	param := &dto.ServiceListInput{}
	if err := param.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	serviceInfo := &dao.ServiceInfo{}
	list, total, err := serviceInfo.PageList(c, tx, param)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	out := &dto.ServiceListOutput{}
	var outList []dto.ServiceListItemOutput
	for _,listItem := range list {
		serviceDetail,err := listItem.ServiceDetail(c, tx, &listItem)
		if err != nil {
			middleware.ResponseError(c, 2003, err)
			return
		}
		serviceAddr := "unknow"
		clusterIp := lib.GetStringConf("base.cluster.cluster_ip")
		clusterPort := lib.GetStringConf("base.cluster.cluster_port")
		clusterSslPort := lib.GetStringConf("base.cluster.cluster_ssl_port")

		if serviceDetail.Info.LoadType == public.LoadTypeHttp && serviceDetail.HttpRule.RuleType == public.HttpRuleTypePreFixUrl && serviceDetail.HttpRule.NeedHttps == 0{
			serviceAddr = clusterIp + ":" + clusterPort + serviceDetail.HttpRule.Rule
		}
		if serviceDetail.Info.LoadType == public.LoadTypeHttp && serviceDetail.HttpRule.RuleType == public.HttpRuleTypePreFixUrl && serviceDetail.HttpRule.NeedHttps == 1{
			serviceAddr = clusterIp + ":" + clusterSslPort + serviceDetail.HttpRule.Rule
		}
		if serviceDetail.Info.LoadType == public.LoadTypeHttp && serviceDetail.HttpRule.RuleType == public.HttpRuleTypeDomain{
			serviceAddr = serviceDetail.HttpRule.Rule
		}

		if serviceDetail.Info.LoadType == public.LoadTypeTcp{
			serviceAddr = clusterIp + clusterPort
		}

		if serviceDetail.Info.LoadType == public.LoadTypeGrpc{
			serviceAddr = fmt.Sprintf("%s:%d", clusterIp,serviceDetail.GRPCRule.Port)
		}
		ipList := serviceDetail.LoadBalanceRule.GetIPListByModel()
		outListItem := dto.ServiceListItemOutput{
			ID: listItem.ID,
			ServiceName: listItem.ServiceName,
			ServiceDesc: listItem.ServiceDesc,
			ServiceAddr: serviceAddr,
			TotalNode: len(ipList),
		}
		outList = append(outList, outListItem)
	}
	out.List = outList
	out.Total = total
	middleware.ResponseSuccess(c, out)
	return
}

// ServiceDelete godoc
// @Summary 服务删除
// @Description 服务删除
// @Tags 服务管理
// @ID /service/service_delete
// @Accept  json
// @Produce  json
// @Param id query string true "服务ID"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_delete [get]
func (service *ServiceController) ServiceDelete(c *gin.Context)  {
	params := &dto.ServiceDeleteInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	serviceInfo := &dao.ServiceInfo{ID: params.ID}
	serviceInfo, err = serviceInfo.Find(c, tx, serviceInfo)
	serviceInfo.IsDelete = 1
	if err := serviceInfo.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")

}
