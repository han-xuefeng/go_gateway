package controller

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/han-xuefeng/go_gateway/dao"
	"github.com/han-xuefeng/go_gateway/dto"
	"github.com/han-xuefeng/go_gateway/middleware"
)

type ServiceController struct {
}

func ServiceRegister(group *gin.RouterGroup) {
	serviceController := &ServiceController{}
	group.GET("service_list", serviceController.ServiceList)
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
	outList := []dto.ServiceListItemOutput{}
	for _,listItem := range list {
		outListItem := dto.ServiceListItemOutput{
			Id: listItem.Id,
			ServiceName: listItem.ServiceName,
			ServiceDesc: listItem.ServiceName,
		}
		outList = append(outList, outListItem)
	}
	out.List = outList
	out.Total = total
	middleware.ResponseSuccess(c, out)
	return
}
