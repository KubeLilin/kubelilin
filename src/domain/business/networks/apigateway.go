package networks

import "gorm.io/gorm"

type ApiGatewayService struct {
	db *gorm.DB
}

func NewApiGatewayService(db *gorm.DB) *ApiGatewayService {
	return &ApiGatewayService{db: db}
}
