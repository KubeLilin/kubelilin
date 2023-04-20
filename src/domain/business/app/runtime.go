package app

import (
	"gorm.io/gorm"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/database/models"
	"time"
)

type RuntimeService struct {
	db *gorm.DB
}

func NewRuntimeService(db *gorm.DB) *RuntimeService {
	return &RuntimeService{db: db}
}

func (service RuntimeService) SaveDaprComponent(req *requests.RuntimeReq) (models.ApplicationDaprCoponentsTemplete, error) {
	time := time.Now()
	if req.ID > 0 {
		model := models.ApplicationDaprCoponentsTemplete{
			Name:          req.ComponentType,
			ComponentType: req.ComponentType,
			Doc:           req.Doc,
			Template:      req.Template,
			UpdateTime:    &time,
		}
		err := service.db.Model(&models.ApplicationDaprCoponentsTemplete{}).
			Where("id = ?", req.ID).Updates(&model).Error
		return model, err
	} else {
		model := models.ApplicationDaprCoponentsTemplete{
			Name:          req.ComponentType,
			ComponentType: req.ComponentType,
			Doc:           req.Doc,
			Template:      req.Template,
			CreateTime:    &time,
			UpdateTime:    &time,
		}
		err := service.db.Create(&model).Error
		return model, err
	}
}

// GetDaprComponentTemplateByType get dapr component type by id
func (service RuntimeService) GetDaprComponentTemplateByType(typeName string) (string, error) {
	var componentType string
	sql := `SELECT template FROM application_dapr_coponents_templete WHERE component_type = ? LIMIT 1 `
	err := service.db.Raw(sql, typeName).First(&componentType).Error
	return componentType, err
}

func (service RuntimeService) GetDaprComponentList() ([]models.ApplicationDaprCoponentsTemplete, error) {
	var list []models.ApplicationDaprCoponentsTemplete
	err := service.db.Model(&models.ApplicationDaprCoponentsTemplete{}).Find(&list).Error
	return list, err
}

func (service RuntimeService) DeleteDaprComponent(id uint64) error {
	err := service.db.Delete(&models.ApplicationDaprCoponentsTemplete{}, id).Error
	return err
}

func (service RuntimeService) GetDaprComponentTypes() ([]string, error) {
	var componentTypes []string
	sql := `SELECT DISTINCT component_type componentType FROM application_dapr_coponents_templete 
ORDER BY component_type DESC`
	err := service.db.Raw(sql).Scan(&componentTypes).Error
	return componentTypes, err
}
