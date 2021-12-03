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

type DeploymentRequest struct {
	mvc.RequestBody
	*models.SgrTenantDeployments
}

type ContainersRequest struct {
	mvc.RequestBody
	Containers []*models.SgrTenantDeploymentsContainers `json:"containers"`
}
