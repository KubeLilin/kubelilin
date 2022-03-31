package app

import (
	"errors"
	"gorm.io/gorm"
	"kubelilin/api/req"
	"kubelilin/api/res"
	"kubelilin/domain/database/models"
	"time"
)

type ServiceConnectionService struct {
	db *gorm.DB
}

func (scs *ServiceConnectionService) CreateServiceConnection(req req.ServiceConnectionReq) (*req.ServiceConnectionReq, error) {
	if req.ServiceType != 1 && req.ServiceType != 2 {
		return nil, errors.New("请选择正确的连接类型")
	}
	var mainDatum = models.ServiceConnection{}
	var credentialDatum = models.ServiceConnectionCredentials{}
	var connectionDatum = models.ServiceConnectionDetails{}
	mainDatum.ServiceType = req.ServiceType
	mainDatum.Name = req.Name
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
			connErr := tx.Model(models.ServiceConnectionDetails{}).Create(connectionDatum)
			if connErr.Error != nil {
				return connErr.Error
			}
		}
		if req.ServiceType == 2 {
			credentialErr := tx.Model(models.ServiceConnectionCredentials{}).Create(credentialDatum)
			if credentialErr.Error != nil {
				return credentialErr.Error
			}
		}
		return nil
	})
	if dbErr != nil {
		return nil, dbErr
	}
	return &req, nil
}

func (scs *ServiceConnectionService) UpdateServiceConnection(req req.ServiceConnectionReq) (*req.ServiceConnectionReq, error) {
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
		mainErr := tx.Model(models.ServiceConnection{}).Updates(&mainDatum)
		if mainErr.Error != nil {
			return mainErr.Error
		}
		if req.ServiceType == 1 {
			connectionDatum.UpdateTime = &currtime
			connErr := tx.Model(models.ServiceConnectionDetails{}).Updates(connectionDatum)
			if connErr.Error != nil {
				return connErr.Error
			}
		}
		if req.ServiceType == 2 {
			credentialDatum.UpdateTime = &currtime
			credentialErr := tx.Model(models.ServiceConnectionCredentials{}).Updates(credentialDatum)
			if credentialErr.Error != nil {
				return credentialErr.Error
			}
		}
		return nil
	})
	if dbErr != nil {
		return nil, dbErr
	}
	return &req, nil
}

func (scs *ServiceConnectionService) QueryServiceConnections(req req.ServiceConnectionPageReq) ([]res.ServiceConnectionPageRes, error) {
	return nil, nil
}
