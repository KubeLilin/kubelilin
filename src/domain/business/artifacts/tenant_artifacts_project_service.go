package artifacts

import (
	"gorm.io/gorm"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/database/models"
	"kubelilin/pkg/page"
	"strings"
	"time"
)

type TenantArtifactsProjectService struct {
	db *gorm.DB
}

func NewTenantArtifactsProjectService(db *gorm.DB) *TenantArtifactsProjectService {
	return &TenantArtifactsProjectService{
		db: db,
	}
}

// CreateTenantArtifactsProject 创建租户制品项目/**
func (svc *TenantArtifactsProjectService) CreateTenantArtifactsProject(reqData requests.CreateTenantArtifactsProjectReq) {
	var now = time.Now()
	dbData := models.TenantArtifactsProject{
		TenantID:        reqData.TenantId,
		HarborProjectID: reqData.HarborProjectId,
		ProjectName:     reqData.ProjectName,
		CreateTime:      &now,
	}
	svc.db.Save(dbData)
}

// CreateTenantArtifactsProject 分页查询租户制品项目/**
func (svc *TenantArtifactsProjectService) QueryTenantArtifactsProject(req requests.QueryTenantArtifactsProjectReq) (err error, pageRes *page.Page) {
	sql := strings.Builder{}
	var res []models.TenantArtifactsProject
	var sqlParams []interface{}
	sqlParams = append(sqlParams, req.TenantId)
	sql.WriteString("select * from tenant_artifacts_project where 1=1 and tenant_id=? ")
	if req.ProjectName != "" {
		sql.WriteString(" and project_name like '%?%'")
		sqlParams = append(sqlParams, req.ProjectName)
	}
	return page.StartPage(svc.db, req.PageIndex, req.PageSize).DoScan(res, sql.String(), sqlParams...)
}
