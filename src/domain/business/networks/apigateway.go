package networks

import (
	"errors"
	"gorm.io/gorm"
	"kubelilin/domain/database/models"
)

type ApiGatewayService struct {
	db *gorm.DB
}

func NewApiGatewayService(db *gorm.DB) *ApiGatewayService {
	return &ApiGatewayService{db: db}
}

func (service *ApiGatewayService) GetAllGatewayList(clusterId uint64) ([]models.ApplicationAPIGateway, error) {
	var gatewayList []models.ApplicationAPIGateway
	sql := `SELECT gw.id,gw.name,gw.desc,gw.cluster_id,tc.name admin_uri FROM application_api_gateway gw
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
