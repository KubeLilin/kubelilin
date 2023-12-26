package deliverables

import (
	"gorm.io/gorm"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/database/models"
	"kubelilin/domain/dto"
	"strconv"
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
	treeData := make([]dto.DeliverablesTreeDTO, 0)
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
	treeData = s.recursionTreeData(treeRoot)
	return treeData
}

func (s *TenantDeliverablesTreeService) recursionTreeData(node models.TenantDeliverablesTree) []dto.DeliverablesTreeDTO {
	var dbData []models.TenantDeliverablesTree
	var resData []dto.DeliverablesTreeDTO
	treeNode := dto.DeliverablesTreeDTO{Key: strconv.FormatUint(node.ID, 10), Title: node.Name, Children: make([]dto.DeliverablesTreeDTO, 0)}
	resData = append(resData, treeNode)
	s.db.Model(models.TenantDeliverablesTree{}).Where("parent_id=?", node.ID).Scan(&dbData)
	if len(dbData) > 0 {
		for _, nodeItem := range dbData {
			childrenData := s.recursionTreeData(nodeItem)
			treeNode.Children = append(treeNode.Children, childrenData...)
		}
	}
	return resData

}
