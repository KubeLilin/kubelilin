package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yoyofx/glinq"
	"gorm.io/gorm"
	"kubelilin/api/req"
	"kubelilin/api/res"
	"kubelilin/domain/database/models"
	"kubelilin/domain/dto"
	"kubelilin/pkg/page"
	"strings"
	"time"
)

type ServiceConnectionService struct {
	db *gorm.DB
}

//1.github 2..gitlab 3.gogos 4.gitee
const (
	GITHUB = 1
	GITLAB = 2
	GOGS   = 3
	GITEE  = 4
)

func NewServiceConnectionService(db *gorm.DB) *ServiceConnectionService {
	return &ServiceConnectionService{
		db: db,
	}
}

func (svc *ServiceConnectionService) CreateServiceConnection(req *req.ServiceConnectionReq) (*req.ServiceConnectionReq, error) {
	if req.Type > 1 {
		var existsCount int64 = 0
		svc.db.Model(&models.ServiceConnection{}).Where("service_type=?", req.ServiceType).Count(&existsCount)
		if existsCount > 0 {
			return nil, errors.New("current service type exists")
		}
	}

	var mainDatum = models.ServiceConnection{}
	var connectionDatum = models.ServiceConnectionDetails{}
	mainDatum.ServiceType = req.ServiceType
	mainDatum.Name = req.Name
	mainDatum.TenantID = req.TenantID
	connectionDatum.Type = req.Type
	connectionDatum.Detail = req.Detail

	dbErr := svc.db.Transaction(func(tx *gorm.DB) error {
		mainErr := tx.Model(models.ServiceConnection{}).Create(&mainDatum)
		if mainErr.Error != nil {
			return mainErr.Error
		}
		connectionDatum.MainID = mainDatum.ID
		connErr := tx.Model(models.ServiceConnectionDetails{}).Create(&connectionDatum)
		if connErr.Error != nil {
			return connErr.Error
		}
		return nil
	})
	if dbErr != nil {
		return nil, dbErr
	}
	return req, nil
}

func (svc *ServiceConnectionService) UpdateServiceConnection(req *req.ServiceConnectionReq) (*req.ServiceConnectionReq, error) {
	if req.ID == 0 {
		return nil, errors.New("非法标识")
	}
	var mainDatum = models.ServiceConnection{}
	var connectionDatum = models.ServiceConnectionDetails{}
	mainRes := svc.db.Model(models.ServiceConnection{}).Where("id=?", req.ID).First(&mainDatum)
	if mainRes.Error != nil {
		return nil, mainRes.Error
	}
	mainDatum.ServiceType = req.ServiceType
	mainDatum.Name = req.Name
	connRes := svc.db.Model(models.ServiceConnectionDetails{}).Where("main_id=?", req.ID).First(&connectionDatum)
	if connRes.Error != nil {
		return nil, connRes.Error
	}

	dbErr := svc.db.Transaction(func(tx *gorm.DB) error {
		currtime := time.Now()
		mainDatum.UpdateTime = &currtime
		mainErr := tx.Model(models.ServiceConnection{}).Where("id=?", req.ID).Updates(&mainDatum)
		if mainErr.Error != nil {
			return mainErr.Error
		}
		connectionDatum.Detail = req.Detail
		connectionDatum.UpdateTime = &currtime
		connErr := tx.Model(models.ServiceConnectionDetails{}).Where("id=?", connectionDatum.ID).Updates(connectionDatum)
		if connErr.Error != nil {
			return connErr.Error
		}

		return nil
	})
	if dbErr != nil {
		return nil, dbErr
	}
	return req, nil
}

func (svc *ServiceConnectionService) QueryServiceConnections(req *req.ServiceConnectionPageReq) (*page.Page, error) {
	data := &[]res.ServiceConnectionRes{}
	var params []interface{}
	sql := strings.Builder{}
	sql.WriteString("select * from service_connection")
	sql.WriteString(" where 1=1 ")
	if len(req.Name) > 0 {
		sql.WriteString(" and name like ? ")
		params = append(params, "%"+req.Name+"%")
	}

	if req.ServiceType > 0 {
		sql.WriteString(" and service_type=? ")
		params = append(params, req.ServiceType)
	}

	err, pageRes := page.StartPage(svc.db, req.PageIndex, req.PageSize).DoScan(&data, sql.String(), params...)
	fmt.Println(pageRes.Data)
	return pageRes, err
}

func (svc *ServiceConnectionService) QueryServiceConnectionInfo(id int64) (*res.ServiceConnectionRes, error) {
	var datum res.ServiceConnectionRes
	var mainDatum models.ServiceConnection
	mainErr := svc.db.Model(&models.ServiceConnection{}).Where("id=?", id).First(&mainDatum)
	if mainErr.Error != nil {
		return nil, mainErr.Error
	}
	datum.ID = mainDatum.ID
	datum.Name = mainDatum.Name
	datum.ServiceType = mainDatum.ServiceType
	var serviceConnectionDatum models.ServiceConnectionDetails
	err := svc.db.Model(&models.ServiceConnectionDetails{}).Where("main_id=?", id).First(&serviceConnectionDatum)
	if err.Error != nil {
		return nil, err.Error
	}
	datum.Type = serviceConnectionDatum.Type
	datum.Detail = serviceConnectionDatum.Detail

	return &datum, nil
}

func (svc *ServiceConnectionService) DeleteServiceConnectionInfo(id uint64) error {
	svc.db.Model(&models.ServiceConnectionDetails{}).Delete(&models.ServiceConnectionDetails{}, "main_id=?", id)
	return svc.db.Model(&models.ServiceConnection{}).Delete(&models.ServiceConnection{}, "id=?", id).Error
}

func (svc *ServiceConnectionService) QueryRepoListByType(tenantId uint64, repoType string) ([]res.ServiceConnectionRes, error) {
	var sb strings.Builder
	data := make([]res.ServiceConnectionRes, 0)
	sb.WriteString("select t1.id,t1.name,t2.detail from service_connection as t1 ")
	sb.WriteString("inner join service_connection_details as t2 ON  t1.id=t2.main_id and t2.type=?")
	serviceType := svc.switchServiceType(repoType)
	dbErr := svc.db.Raw(sb.String(), serviceType).Scan(&data)
	if dbErr.Error != nil {
		return nil, dbErr.Error
	}
	return data, nil
}

func (svc *ServiceConnectionService) switchServiceType(name string) int {

	switch name {
	case "github":
		return GITHUB
		break
	case "gitlab":
		return GITLAB
		break
	case "gogs":
		return GOGS
		break
	case "gitee":
		return GITEE
		break
	}
	return 0
}

func (svc *ServiceConnectionService) GetImageHub() (dto.ServiceConnectionInfo, error) {
	return queryServiceConnectionList(svc.db).Where(func(e dto.ServiceConnectionInfo) bool {
		return e.ServiceType == dto.SC_IMAGEHUB
	}).First()

}

func (svc *ServiceConnectionService) GetPipelineEngine() (dto.ServiceConnectionInfo, error) {
	return queryServiceConnectionList(svc.db).Where(func(e dto.ServiceConnectionInfo) bool {
		return e.ServiceType == dto.SC_PIPELINE
	}).First()
}

func (svc *ServiceConnectionService) GetSystemCallback() (dto.ServiceConnectionInfo, error) {
	return queryServiceConnectionList(svc.db).Where(func(e dto.ServiceConnectionInfo) bool {
		return e.ServiceType == dto.SC_CALLBACK
	}).First()
}

func (svc *ServiceConnectionService) GetTest() (dto.ServiceConnectionInfo, error) {
	return queryServiceConnectionList(svc.db).Where(func(e dto.ServiceConnectionInfo) bool {
		return e.ServiceType == 99999
	}).First()
}

func queryServiceConnectionList(db *gorm.DB) glinq.Queryable[dto.ServiceConnectionInfo] {
	var serviceConnectionList []dto.ServiceConnectionInfo
	sql := `SELECT scd.detail,sc.service_type,scd.type as value FROM service_connection sc
INNER JOIN service_connection_details scd on scd.main_id = sc.id WHERE scd.enable = 1`
	var list []dto.ServiceConnectionDTO
	err := db.Raw(sql).Scan(&list).Error
	if err == nil {
		serviceConnectionList = make([]dto.ServiceConnectionInfo, 0)
		for _, item := range list {
			var serviceConnectionInfo dto.ServiceConnectionInfo
			hasErr := json.Unmarshal([]byte(item.Detail), &serviceConnectionInfo)
			if hasErr == nil {
				serviceConnectionInfo.ServiceType = item.ServiceType
				serviceConnectionList = append(serviceConnectionList, serviceConnectionInfo)
			}
		}
	}
	return glinq.From(serviceConnectionList)
}
