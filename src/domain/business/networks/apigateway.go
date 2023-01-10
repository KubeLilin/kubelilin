package networks

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/yoyofx/glinq"
	"gorm.io/gorm"
	"kubelilin/api/dto/requests"
	"kubelilin/api/dto/responses"
	"kubelilin/domain/database/models"
	"kubelilin/domain/dto"
)

type ApiGatewayService struct {
	db *gorm.DB
}

func NewApiGatewayService(db *gorm.DB) *ApiGatewayService {
	return &ApiGatewayService{db: db}
}

func (service *ApiGatewayService) GetAllGatewayList(clusterId uint64) ([]models.ApplicationAPIGateway, error) {
	var gatewayList []models.ApplicationAPIGateway
	sql := `SELECT gw.id,gw.name,gw.desc,gw.cluster_id,tc.name admin_uri,gw.export_ip,gw.vip FROM application_api_gateway gw
	INNER JOIN sgr_tenant_cluster tc on tc.id = gw.cluster_id `
	var sqlParams []interface{}
	if clusterId > 0 {
		sql += ` WHERE gw.cluster_id = ?`
		sqlParams = append(sqlParams, clusterId)
	}
	err := service.db.Raw(sql, sqlParams).Find(&gatewayList).Error
	return gatewayList, err
}

func (service *ApiGatewayService) GetById(id uint64) (models.ApplicationAPIGateway, error) {
	var gateway models.ApplicationAPIGateway
	err := service.db.Model(&models.ApplicationAPIGateway{}).First(&gateway, "id=?", id).Error
	return gateway, err
}

func (service *ApiGatewayService) GetByClusterId(clusterId uint64) (models.ApplicationAPIGateway, error) {
	var gateway models.ApplicationAPIGateway
	err := service.db.Model(&models.ApplicationAPIGateway{}).First(&gateway, "cluster_id=?", clusterId).Error
	return gateway, err
}

func (service *ApiGatewayService) CreateTeam(team models.ApplicationAPIGatewayTeams) error {
	var exitCount int64
	service.db.Model(&models.DevopsProjects{}).Where("tenant_id=? and gateway_id=? and name=?", team.TenantID, team.GatewayID, team.Name).Count(&exitCount)
	if exitCount > 0 {
		return errors.New("already have the same name gateway")
	}
	return service.db.Create(&team).Error
}

func (service *ApiGatewayService) CreateOrUpdateTeam(team models.ApplicationAPIGatewayTeams) error {
	if team.ID > 0 {
		return service.db.
			Model(&models.ApplicationAPIGatewayTeams{}).
			Where("id=?", team.ID).
			Updates(team).Error
	} else {
		return service.CreateTeam(team)
	}
}

func (service *ApiGatewayService) GetAllGatewayTeamList(gatewayId uint64, tenantId uint64) ([]models.ApplicationAPIGatewayTeams, error) {
	var gatewayList []models.ApplicationAPIGatewayTeams
	err := service.db.Model(&models.ApplicationAPIGatewayTeams{}).
		Where("gateway_id=? AND tenant_id=?", gatewayId, tenantId).Find(&gatewayList).Error
	return gatewayList, err
}

func (service *ApiGatewayService) GetRouterList(requestRouter *requests.GatewayRouterListRequest) ([]models.ApplicationAPIGatewayRouters, error) {
	var gatewayList []models.ApplicationAPIGatewayRouters
	query := service.db.Model(&models.ApplicationAPIGatewayRouters{}).
		Where("team_id=?", requestRouter.TeamId)

	if requestRouter.Name != "" {
		query = query.Where("name like ?", "%"+requestRouter.Name+"%")
	}

	if requestRouter.Host != "" {
		query = query.Where("host like ?", "%"+requestRouter.Host+"%")
	}
	if requestRouter.Desc != "" {
		query = query.Where("desc like ?", "%"+requestRouter.Desc+"%")
	}

	err := query.Find(&gatewayList).Error
	return gatewayList, err
}

func (service *ApiGatewayService) GetRouterListBy(applicationId uint64, deploymentId uint64) ([]models.ApplicationAPIGatewayRouters, error) {
	var gatewayList []models.ApplicationAPIGatewayRouters
	query := service.db.Model(&models.ApplicationAPIGatewayRouters{})
	if applicationId > 0 {
		query.Where("application_id=?", applicationId)
	}
	if deploymentId > 0 {
		query.Where("deployment_id=?", deploymentId)
	}
	err := query.Find(&gatewayList).Error
	return gatewayList, err
}

func (service *ApiGatewayService) GetRouterByDeployIdAndName(deploymentId uint64, routeName string) (models.ApplicationAPIGatewayRouters, error) {
	var gatewayRoute models.ApplicationAPIGatewayRouters
	query := service.db.Model(&models.ApplicationAPIGatewayRouters{})

	if deploymentId > 0 {
		query.Where("deployment_id=?", deploymentId)
	}

	if routeName != "" {
		query.Where("name=?", routeName)
	}

	err := query.First(&gatewayRoute).Error
	return gatewayRoute, err
}

func (service *ApiGatewayService) GetAppList(tenantId uint64) ([]responses.LabelValues, error) {
	var applist []models.SgrTenantApplication
	query := service.db.Raw("SELECT id,name FROM sgr_tenant_application WHERE tenant_id = ?", tenantId)
	err := query.Find(&applist).Error
	res := glinq.Map(glinq.From(applist), func(app models.SgrTenantApplication) responses.LabelValues {
		return responses.LabelValues{
			Label: app.Name,
			Value: app.ID,
		}
	})

	return res.ToSlice(), err
}

func (service *ApiGatewayService) GetDeploymentList(tenantId uint64, clusterId uint64, appId uint64) ([]responses.LabelValues, error) {
	var deploymentList []models.SgrTenantDeployments
	query := service.db.Raw(`SELECT id,name FROM sgr_tenant_deployments 
WHERE tenant_id = ? AND cluster_id = ? AND app_id = ?`, tenantId, clusterId, appId)
	err := query.Find(&deploymentList).Error
	res := glinq.Map(glinq.From(deploymentList), func(deploy models.SgrTenantDeployments) responses.LabelValues {
		return responses.LabelValues{
			Label: deploy.Name,
			Value: deploy.ID,
		}
	})
	return res.ToSlice(), err
}

func (service *ApiGatewayService) CreateOrEditRouter(request *requests.GatewayRouterRequest, deployment dto.DeploymentItemDto) (*models.ApplicationAPIGatewayRouters, error) {
	var router models.ApplicationAPIGatewayRouters
	_ = copier.Copy(&router, request)

	router.Label = "deployment"
	router.Websocket = 1
	router.Nodes = fmt.Sprintf("%s.%s:%d", deployment.ServiceName, deployment.NameSpace, deployment.ServicePort)
	router.Status = 1

	var err error
	if router.ID > 0 {
		// update the router
		err = service.db.Model(&models.ApplicationAPIGatewayRouters{}).
			Where("id=?", router.ID).
			Updates(&router).Error
	} else {
		// create a new router
		err = service.db.Create(&router).Error
	}
	return &router, err
}

func (service *ApiGatewayService) DeleteRouter(id uint64) error {
	return service.db.Delete(&models.ApplicationAPIGatewayRouters{}, "id=?", id).Error
}
