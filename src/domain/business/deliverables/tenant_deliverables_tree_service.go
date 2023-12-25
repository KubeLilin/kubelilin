package deliverables

import (
	"gorm.io/gorm"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/database/models"
	"kubelilin/domain/dto"
)

type TenantDeliverablesTreeService struct {
	db *gorm.DB
}

func NewTenantDeliverablesTree(db *gorm.DB) *TenantDeliverablesTreeService {
	return &TenantDeliverablesTreeService{
		db: db,
	}
}

func (s *TenantDeliverablesTreeService) EditTree(req requests.EditTenantDeliverablesTreeReq) error {
	treeData := models.TenantDeliverablesTree{
		ProjectID: req.ProjectId,
		Name:      req.Name,
		ParentID:  req.ParentID,
	}
	tx := s.db.Save(treeData)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *TenantDeliverablesTreeService) queryTreeByProject(projectId uint64) []dto.DeliverablesTreeDTO {
	treeData := make([]dto.DeliverablesTreeDTO, 1)
	var dbData []models.TenantDeliverablesTree
	s.db.Model(models.TenantDeliverablesTree{}).Where("project_id=?", projectId).Scan(&dbData)
	if len(dbData) <= 0 {
		return treeData
	}
	// 查找根节点
	var treeRoot models.TenantDeliverablesTree
	for _, item := range dbData {
		if item.ParentID == 0 {
			treeRoot = item
			break
		}
	}
}

func (s *TenantDeliverablesTreeService) recursionTreeData(node models.TenantDeliverablesTree) {
	var dbData []models.TenantDeliverablesTree
	s.db.Model(models.TenantDeliverablesTree{}).Where("parent_id=?", node.ID).Scan(&dbData)
	if len(dbData) > 0 {

	}
}
