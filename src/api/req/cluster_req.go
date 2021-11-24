package req

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"sgr/domain/database/models"
)

type ImportClusterReq struct {
	mvc.RequestBody
	NickName string `form:"nickName"`
	TenantId uint64 `form:"tenantId"`
}

type DeploymentReq struct {
	mvc.RequestBody

	Deployments models.SgrTenantDeployments             `json:"deployment"`
	Containers  []models.SgrTenantDeploymentsContainers `json:"containers"`
}
