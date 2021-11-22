package app

import "gorm.io/gorm"

type ApplicationService struct {
	db *gorm.DB
}

func NewApplicationService(db *gorm.DB) *ApplicationService {
	return &ApplicationService{db: db}
}

func (*ApplicationService) CreateApp() {

}
