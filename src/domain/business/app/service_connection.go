package app

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"kubelilin/api/req"
	"kubelilin/api/res"
	"kubelilin/domain/database/models"
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

func (scs *ServiceConnectionService) CreateServiceConnection(req *req.ServiceConnectionReq) (*req.ServiceConnectionReq, error) {
	fmt.Println(req.ServiceType)
	if req.ServiceType != 1 && req.ServiceType != 2 {
		return nil, errors.New("请选择正确的连接类型")
	}
	var mainDatum = models.ServiceConnection{}
	var credentialDatum = models.ServiceConnectionCredentials{}
	var connectionDatum = models.ServiceConnectionDetails{}
	mainDatum.ServiceType = req.ServiceType
	mainDatum.Name = req.Name
	mainDatum.TenantID = req.TenantID
	if req.ServiceType == 1 {

		connectionDatum.Type = req.Type
		connectionDatum.Detail = req.Detail
	} else if req.ServiceType == 2 {
		credentialDatum.Type = req.Type
		credentialDatum.Detail = req.Detail
	}
	dbErr := scs.db.Transaction(func(tx *gorm.DB) error {
		mainErr := tx.Model(models.ServiceConnection{}).Create(&mainDatum)
		if mainErr.Error != nil {
			return mainErr.Error
		}
		if req.ServiceType == 1 {
			connectionDatum.MainID = mainDatum.ID
			connErr := tx.Model(models.ServiceConnectionDetails{}).Create(&connectionDatum)
			if connErr.Error != nil {
				return connErr.Error
			}
		}
		if req.ServiceType == 2 {
			credentialDatum.MainID = mainDatum.ID
			credentialErr := tx.Model(models.ServiceConnectionCredentials{}).Create(&credentialDatum)
			if credentialErr.Error != nil {
				return credentialErr.Error
			}
		}
		return nil
	})
	if dbErr != nil {
		return nil, dbErr
	}
	return req, nil
}

func (scs *ServiceConnectionService) UpdateServiceConnection(req *req.ServiceConnectionReq) (*req.ServiceConnectionReq, error) {
	if req.ID == 0 {
		return nil, errors.New("非法标识")
	}
	var mainDatum = models.ServiceConnection{}
	var credentialDatum = models.ServiceConnectionCredentials{}
	var connectionDatum = models.ServiceConnectionDetails{}
	mainRes := scs.db.Model(models.ServiceConnection{}).Where("id=?", req.ID).First(&mainDatum)
	if mainRes.Error != nil {
		return nil, mainRes.Error
	}
	mainDatum.ServiceType = req.ServiceType
	mainDatum.Name = req.Name
	if req.ServiceType == 1 {
		connRes := scs.db.Model(models.ServiceConnectionDetails{}).Where("main_id=?", req.ID).First(&connectionDatum)
		if connRes.Error != nil {
			return nil, connRes.Error
		}
	} else if req.ServiceType == 2 {
		credentialRes := scs.db.Model(models.ServiceConnectionCredentials{}).Where("main_id=?", req.ID).First(&credentialDatum)
		if credentialRes.Error != nil {
			return nil, credentialRes.Error
		}
	}
	dbErr := scs.db.Transaction(func(tx *gorm.DB) error {
		currtime := time.Now()
		mainDatum.UpdateTime = &currtime
		mainErr := tx.Model(models.ServiceConnection{}).Where("id=?", req.ID).Updates(&mainDatum)
		if mainErr.Error != nil {
			return mainErr.Error
		}
		if req.ServiceType == 1 {
			connectionDatum.UpdateTime = &currtime
			connErr := tx.Model(models.ServiceConnectionDetails{}).Where("id=?", connectionDatum.ID).Updates(connectionDatum)
			if connErr.Error != nil {
				return connErr.Error
			}
		}
		if req.ServiceType == 2 {
			credentialDatum.UpdateTime = &currtime
			credentialErr := tx.Model(models.ServiceConnectionCredentials{}).Where("id=?", credentialDatum.ID).Updates(credentialDatum)
			if credentialErr.Error != nil {
				return credentialErr.Error
			}
		}
		return nil
	})
	if dbErr != nil {
		return nil, dbErr
	}
	return req, nil
}

func (scs *ServiceConnectionService) QueryServiceConnections(req *req.ServiceConnectionPageReq) (*page.Page, error) {
	data := &[]res.ServiceConnectionRes{}
	var params []interface{}
	sql := strings.Builder{}
	sql.WriteString("select * from service_connection")
	sql.WriteString(" where tenant_id=? ")
	params = append(params, req.TenantID)
	if len(req.Name) > 0 {
		sql.WriteString(" and name=? ")
		params = append(params, req.Name)
	}
	err, pageRes := page.StartPage(scs.db, req.PageIndex, req.PageSize).DoScan(&data, sql.String(), params...)
	fmt.Println(pageRes.Data)
	return pageRes, err
}

func (scs *ServiceConnectionService) QueryServiceConnectionInfo(id int64) (*res.ServiceConnectionRes, error) {
	var datum res.ServiceConnectionRes
	var mainDatum models.ServiceConnection
	mainErr := scs.db.Model(&models.ServiceConnection{}).Where("id=?", id).First(&mainDatum)
	if mainErr.Error != nil {
		return nil, mainErr.Error
	}
	datum.ID = mainDatum.ID
	datum.Name = mainDatum.Name
	datum.ServiceType = mainDatum.ServiceType
	if mainDatum.ServiceType == 1 {
		var serviceConnectionDatum models.ServiceConnectionDetails
		err := scs.db.Model(&models.ServiceConnectionDetails{}).Where("main_id=?", id).First(&serviceConnectionDatum)
		if err.Error != nil {
			return nil, err.Error
		}
		datum.Type = serviceConnectionDatum.Type
		datum.Detail = serviceConnectionDatum.Detail
	} else {
		var credentialDatum models.ServiceConnectionCredentials
		err := scs.db.Model(&models.ServiceConnectionCredentials{}).Where("main_id=?", id).First(&credentialDatum)
		if err.Error != nil {
			return nil, err.Error
		}
		datum.Type = credentialDatum.Type
		datum.Detail = credentialDatum.Detail
	}
	return &datum, nil
}

func (svc *ServiceConnectionService) QueryRepoListByType(tenantId uint64, repoType string) ([]res.ServiceConnectionRes, error) {
	var sb strings.Builder
	data := make([]res.ServiceConnectionRes, 0)
	sb.WriteString("select t1.id,t1.name,t2.detail from service_connection as t1 ")
	sb.WriteString("inner join service_connection_details as t2 ON  t1.id=t2.main_id and t1.tenant_id=? and t2.type=?")
	serviceType := svc.switchServiceType(repoType)
	dbErr := svc.db.Raw(sb.String(), tenantId, serviceType).Scan(&data)
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
