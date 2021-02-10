package dao

import (
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"github.com/han-xuefeng/go_gateway/dto"
	"github.com/han-xuefeng/go_gateway/public"
	"time"
)

type ServiceInfo struct {
	Id          int       `json:"id" gorm:"primary_key" description:"自增主键"`
	LoadType    int       `json:"load_type" gorm:"column:load_type" description:"负载类型"`
	ServiceName string    `json:"salt" gorm:"column:salt" description:"服务名称"`
	ServiceDesc string    `json:"service_desc" gorm:"column:service_desc" description:"服务描述"`
	UpdatedAt   time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt   time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete    int       `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (t *ServiceInfo)TableName() string {
	return "gateway_service_info"
}

func (t *ServiceInfo)PageList(c *gin.Context,tx *gorm.DB, param *dto.ServiceListInput) ([]ServiceInfo, int64, error) {
	total := int64(0)
	list := []ServiceInfo{}
	query := tx.SetCtx(public.GetGinTraceContext(c))
	query = query.Table(t.TableName()).Where("is_delete=0")
	if param.Info != "" {
		query = query.Where("service_name like %?% or service_desc like %?%", param.Info, param.Info)
	}
	if err := query.Limit(param.PageSize).Offset((param.PageNo - 1) * param.PageSize).Find(&list).Error;err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	query.Limit(param.PageSize).Offset((param.PageNo - 1) * param.PageSize).Count(&total)
	return list, total ,nil
}
