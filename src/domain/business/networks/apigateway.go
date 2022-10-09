package networks

import (
	"gorm.io/gorm"
	"kubelilin/domain/database/models"
)

type ApiGatewayService struct {
	db *gorm.DB
}

func NewApiGatewayService(db *gorm.DB) *ApiGatewayService {
	return &ApiGatewayService{db: db}
}

func (service *ApiGatewayService) GetById(id uint64) (models.ApplicationAPIGateway, error) {
	var gateway models.ApplicationAPIGateway
	err := service.db.Model(&models.ApplicationAPIGateway{}).First(&gateway, "id=?", id).Error
	return gateway, err
}
