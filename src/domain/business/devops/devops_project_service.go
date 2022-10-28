package devops

import (
	"gorm.io/gorm"
	"kubelilin/domain/dto"
)

type ProjectService struct {
	db *gorm.DB
}

func NewProjectService(db *gorm.DB) *ProjectService {
	return &ProjectService{db: db}
}

func (service *ProjectService) GetResourceMetrics(tenantId uint64, projectId uint64) (dto.DevOpsProjectResourceTotals, error) {
	var resourceList []dto.DevOpsProjectResource
	sql := `SELECT deploy.level,SUM(deploy.replicas) replicas,SUM(container.limit_cpu * deploy.replicas) sum_cpu,SUM(container.limit_memory * deploy.replicas) sum_memory 
FROM devops_projects dps 
INNER JOIN devops_projects_apps dpsapp on dpsapp.project_id = dps.id
INNER JOIN sgr_tenant_deployments deploy on deploy.app_id = dpsapp.application_id
INNER JOIN sgr_tenant_deployments_containers container on container.deploy_id = deploy.id
WHERE dps.tenant_id = ? AND dps.id = ?
GROUP BY deploy.level`
	err := service.db.Raw(sql, tenantId, projectId).Scan(&resourceList).Error
	return getTotalDetailsInfo(resourceList), err
}

func (service *ProjectService) GetResourceMetricsByTenantId(tenantId uint64) (dto.DevOpsProjectResourceTotals, error) {
	var resourceList []dto.DevOpsProjectResource
	sql := `SELECT deploy.level,SUM(deploy.replicas) replicas,SUM(container.limit_cpu * deploy.replicas) sum_cpu,SUM(container.limit_memory * deploy.replicas) sum_memory 
FROM sgr_tenant_deployments deploy 
INNER JOIN sgr_tenant_deployments_containers container on container.deploy_id = deploy.id
WHERE deploy.tenant_id = ?  
GROUP BY deploy.level`
	err := service.db.Raw(sql, tenantId).Scan(&resourceList).Error
	return getTotalDetailsInfo(resourceList), err
}

func getTotalDetailsInfo(items []dto.DevOpsProjectResource) dto.DevOpsProjectResourceTotals {
	resourceTotal := &dto.DevOpsProjectResourceTotals{
		TotalCpu:       0,
		TotalMemory:    0,
		DevMetrics:     dto.DevOpsProjectResource{Level: "dev", Replicas: 0, SumCpu: 0, SumMemory: 0},
		TestMetrics:    dto.DevOpsProjectResource{Level: "test", Replicas: 0, SumCpu: 0, SumMemory: 0},
		ReleaseMetrics: dto.DevOpsProjectResource{Level: "release", Replicas: 0, SumCpu: 0, SumMemory: 0},
		ProdMetrics:    dto.DevOpsProjectResource{Level: "prod", Replicas: 0, SumCpu: 0, SumMemory: 0},
	}
	for _, item := range items {
		resourceTotal.TotalMemory += item.SumMemory
		resourceTotal.TotalCpu += item.SumCpu

		switch item.Level {
		case "dev":
			resourceTotal.DevMetrics = item
			break
		case "test":
			resourceTotal.TestMetrics = item
			break
		case "release":
			resourceTotal.ReleaseMetrics = item
			break
		case "prod":
			resourceTotal.ProdMetrics = item
			break
		}
	}

	return *resourceTotal
}
