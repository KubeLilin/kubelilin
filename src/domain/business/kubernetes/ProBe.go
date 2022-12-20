package kubernetes

import (
	"errors"
	"gorm.io/gorm"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/database/models"
)

type ProBeService struct {
	db *gorm.DB
}

func NewProBeService(_db *gorm.DB) *ProBeService {
	return &ProBeService{
		db: _db,
	}
}

// CreateProBe 创建声明周期探针/**
func (pbs *ProBeService) CreateProBe(proReq *requests.ProbeRequest) error {
	res := pbs.db.Transaction(func(tx *gorm.DB) error {
		//判断readniess是否已经配置
		readiness := &models.DeploymentContainerLifecycleCheck{}
		pbs.db.Model(models.DeploymentContainerLifecycleCheck{}).Where("deployment_id=? and type=? ", proReq.DpId, READINESS).First(readiness)
		if readiness == nil || readiness.ID <= 0 {
			readiness.DeploymentID = proReq.DpId
			readiness.Type = proReq.ReadinessType
			readiness.Port = proReq.ReadinessPort
			readiness.Type = READINESS
			readiness.Path = proReq.ReadinessUrl
			readiness.Scheme = proReq.ReadinessReqScheme
			readiness.PeriodSeconds = proReq.ReadinessPeriodSeconds
			readiness.InitialDelaySeconds = proReq.ReadinessInitialDelaySeconds
			readiness.TimeoutSeconds = proReq.ReadinessTimeoutSeconds
			if proReq.EnableReadiness {
				readiness.Enable = 1
			} else {
				readiness.Enable = 0
			}
			err := tx.Model(models.DeploymentContainerLifecycleCheck{}).Save(readiness).Error
			return err
		} else {
			readiness.Port = proReq.ReadinessPort
			readiness.Path = proReq.ReadinessUrl
			readiness.Scheme = proReq.ReadinessReqScheme
			readiness.PeriodSeconds = proReq.ReadinessPeriodSeconds
			readiness.InitialDelaySeconds = proReq.ReadinessInitialDelaySeconds
			if proReq.EnableReadiness {
				readiness.Enable = 1
			} else {
				readiness.Enable = 0
			}
			err := tx.Model(models.DeploymentContainerLifecycleCheck{}).Where("id=?", readiness.ID).Updates(&readiness).Error
			return err
		}
		// 判断liveness是否已经存在
		liveness := &models.DeploymentContainerLifecycleCheck{}
		pbs.db.Model(models.DeploymentContainerLifecycleCheck{}).Where("deployment_id=? and type=? ", proReq.DpId, LIVENESS).First(liveness)
		if liveness == nil || liveness.ID <= 0 {
			liveness.DeploymentID = proReq.DpId
			liveness.Type = LIVENESS
			liveness.Port = proReq.LivenessPort
			liveness.Path = proReq.LivenessUrl
			liveness.Type = LIVENESS
			liveness.Scheme = proReq.LivenessReqScheme
			liveness.PeriodSeconds = proReq.LivenessPeriodSeconds
			liveness.InitialDelaySeconds = proReq.LivenessInitialDelaySeconds
			liveness.TimeoutSeconds = proReq.LivenessTimeoutSeconds
			if proReq.EnableLiveness {
				liveness.Enable = 1
			} else {
				liveness.Enable = 0
			}
			err := tx.Model(models.DeploymentContainerLifecycleCheck{}).Save(liveness).Error
			return err
		} else {
			liveness.Port = proReq.LivenessPort
			liveness.Path = proReq.LivenessUrl
			liveness.Scheme = proReq.LivenessReqScheme
			liveness.PeriodSeconds = proReq.LivenessPeriodSeconds
			liveness.InitialDelaySeconds = proReq.LivenessInitialDelaySeconds
			if proReq.EnableLiveness {
				readiness.Enable = 1
			} else {
				readiness.Enable = 0
			}
			err := tx.Model(models.DeploymentContainerLifecycleCheck{}).Where("id=?", liveness.ID).Updates(&liveness).Error
			return err
		}
		mainContainer := &models.SgrTenantDeploymentsContainers{}
		tx.Model(models.SgrTenantDeploymentsContainers{}).Where("deploy_id=? and is_main=1", proReq.DpId).First(mainContainer)
		if mainContainer == nil {
			return errors.New("can't find the sole container of development ")
		}
		tx.Model(models.SgrTenantDeployments{}).Updates(models.SgrTenantDeployments{TerminationGracePeriodSeconds: proReq.TerminationGracePeriodSeconds, MaxUnavailable: &proReq.MaxUnavailable, MaxSurge: &proReq.MaxSurge}).Where("id=?", proReq.DpId)
		updateDatum := models.SgrTenantDeploymentsContainers{Poststart: proReq.LifecyclePreStart, Podstop: proReq.LifecyclePreStop}
		y := uint8(1)
		n := uint8(0)
		if proReq.EnableLifecycle {
			updateDatum.EnableLife = &y
		} else {
			updateDatum.EnableLife = &n
		}
		tx.Model(models.SgrTenantDeploymentsContainers{}).Updates(updateDatum).Where("deploy_id=? and is_main=1", proReq.DpId)
		return nil
	})
	return res

}

// GetProBe 获取探针信息 /**
func (pbs *ProBeService) GetProBe(dpId uint64) (*requests.ProbeRequest, error) {
	res := &requests.ProbeRequest{}
	dp := models.SgrTenantDeployments{}
	pbs.db.Model(models.SgrTenantDeployments{}).Where("id=?", dpId).First(&dp)
	if dp.ID <= 0 {
		return nil, errors.New("没有招到对应的deployment")
	}
	if dp.MaxUnavailable != nil {
		res.MaxUnavailable = *dp.MaxUnavailable
	}
	if dp.MaxSurge != nil {
		res.MaxSurge = *dp.MaxSurge
	}
	res.TerminationGracePeriodSeconds = dp.TerminationGracePeriodSeconds
	mainContainer := &models.SgrTenantDeploymentsContainers{}
	pbs.db.Model(models.SgrTenantDeploymentsContainers{}).Where("deploy_id=? and is_main=1", dpId).First(mainContainer)
	if mainContainer != nil {
		res.LifecyclePreStart = mainContainer.Poststart
		res.LifecyclePreStop = mainContainer.Podstop
		if mainContainer.EnableLife != nil {
			if *mainContainer.EnableLife == 1 {
				res.EnableLifecycle = true
			} else {
				res.EnableLifecycle = false
			}
		}
	}
	readiness := models.DeploymentContainerLifecycleCheck{}
	pbs.db.Model(models.DeploymentContainerLifecycleCheck{}).Where("deployment_id=? and  type=?", dpId, READINESS).First(&readiness)
	if readiness.ID > 0 {
		res.DpId = readiness.DeploymentID
		res.ReadinessPort = readiness.Port
		res.ReadinessUrl = readiness.Path
		res.ReadinessReqScheme = readiness.Scheme
		res.ReadinessPeriodSeconds = readiness.PeriodSeconds
		res.ReadinessInitialDelaySeconds = readiness.InitialDelaySeconds
		res.EnableReadiness = readiness.Enable == 1
		res.ReadinessTimeoutSeconds = readiness.TimeoutSeconds
	}
	liveness := models.DeploymentContainerLifecycleCheck{}
	pbs.db.Model(models.DeploymentContainerLifecycleCheck{}).Where("deployment_id=? and  type=?", dpId, LIVENESS).First(&liveness)
	if liveness.ID > 0 {
		res.DpId = liveness.DeploymentID
		res.LivenessPort = liveness.Port
		res.LivenessUrl = liveness.Path
		res.LivenessReqScheme = liveness.Scheme
		res.LivenessPeriodSeconds = liveness.PeriodSeconds
		res.LivenessInitialDelaySeconds = liveness.InitialDelaySeconds
		res.EnableLiveness = liveness.Enable == 1
		res.LivenessTimeoutSeconds = liveness.TimeoutSeconds
	}
	return res, nil
}
