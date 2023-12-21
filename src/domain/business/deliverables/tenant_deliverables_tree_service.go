package deliverables

import "gorm.io/gorm"

type TenantDeliverablesTreeService struct {
	db *gorm.DB
}

func NewTenantDeliverablesTree(db *gorm.DB) *TenantDeliverablesTreeService {
	return &TenantDeliverablesTreeService{
		db: db,
	}
}
