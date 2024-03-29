package deliverables

import (
	"gorm.io/gorm"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/business/harbor"
	"kubelilin/domain/database/models"
	"kubelilin/pkg/page"
	"strings"
	"time"
)

type TenantDeliverablesProjectService struct {
	db *gorm.DB
}

func NewTenantDeliverablesProjectService(db *gorm.DB) *TenantDeliverablesProjectService {
	return &TenantDeliverablesProjectService{
		db: db,
	}
}

// CreateTenantArtifactsProject 创建租户制品项目/**
func (svc *TenantDeliverablesProjectService) CreateTenantDeliverablesProject(reqData *requests.CreateTenantDeliverablesProjectReq) error {
	var now = time.Now()
	dbData := models.TenantDeliverablesProject{
		TenantID:            reqData.TenantId,
		ProjectName:         reqData.ProjectName,
		ServiceConnectionID: reqData.ServiceConnectionId,
		CreateTime:          &now,
	}
	// 获取habor连接
	var serviceConnectionDatum models.ServiceConnectionDetails
	svc.db.Model(&models.ServiceConnectionDetails{}).Where("main_id=?", reqData.ServiceConnectionId).First(&serviceConnectionDatum)
	// 调用harborapi创建项目
	err := harbor.CreateProject(reqData.ProjectName, serviceConnectionDatum)
	if err != nil {
		return err
	}
	// 查询harbor项目
	err, dtos := harbor.QueryProjectPage(reqData.ProjectName, serviceConnectionDatum)
	if err != nil {
		return err
	}
	for _, dto := range dtos {
		if dto.Name == reqData.ProjectName {
			dbData.HarborProjectID = uint64(dto.ProjectId)
		}
	}
	svc.db.Save(&dbData)
	return nil
}

// CreateTenantArtifactsProject 分页查询租户制品项目
func (svc *TenantDeliverablesProjectService) QueryTenantDeliverablesProject(req *requests.QueryTenantDeliverablesProjectReq) (err error, pageRes *page.Page) {
	sql := strings.Builder{}
	var res []models.TenantDeliverablesProject
	var sqlParams []interface{}
	sqlParams = append(sqlParams, req.TenantId)
	sql.WriteString("select * from tenant_deliverables_project where 1=1 and tenant_id=? ")
	if req.ProjectName != "" {
		sql.WriteString(" and project_name like '%")
		sql.WriteString(req.ProjectName)
		sql.WriteString("%'")
		sqlParams = append(sqlParams, req.ProjectName)
	}
	return page.StartPage(svc.db, req.CurrentPage, req.PageSize).DoScan(&res, sql.String(), sqlParams...)
}
