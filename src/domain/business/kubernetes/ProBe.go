package kubernetes

import (
	"errors"
	"gorm.io/gorm"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/database/models"
	"kubelilin/utils"
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
			if err != nil {
				return err
			}
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
			err := tx.Save(readiness).Error
			if err != nil {
				return err
			}
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
			if err != nil {
				return err
			}
		} else {
			liveness.Port = proReq.LivenessPort
			liveness.Path = proReq.LivenessUrl
			liveness.Scheme = proReq.LivenessReqScheme
			liveness.PeriodSeconds = proReq.LivenessPeriodSeconds
			liveness.InitialDelaySeconds = proReq.LivenessInitialDelaySeconds
			if proReq.EnableLiveness {
				liveness.Enable = 1
			} else {
				liveness.Enable = 0
			}
			err := tx.Save(liveness).Error
			if err != nil {
				return err
			}
		}

		tx.Model(models.SgrTenantDeployments{}).Where("id=?", proReq.DpId).Updates(models.SgrTenantDeployments{TerminationGracePeriodSeconds: proReq.TerminationGracePeriodSeconds, MaxUnavailable: &proReq.MaxUnavailable, MaxSurge: &proReq.MaxSurge})

		mainContainer := &models.SgrTenantDeploymentsContainers{}
		tx.Model(models.SgrTenantDeploymentsContainers{}).Where("deploy_id=? and is_main=1", proReq.DpId).First(mainContainer)
		if mainContainer == nil {
			return errors.New("can't find the sole container of development ")
		}
		//Poststart: proReq.LifecyclePreStart, Podstop: proReq.LifecyclePreStop
		if proReq.EnableLifecycle {
			mainContainer.EnableLife = utils.UInt8Ptr(1)
		} else {
			mainContainer.EnableLife = utils.UInt8Ptr(0)
		}
		mainContainer.Poststart = proReq.LifecyclePreStart
		mainContainer.Podstop = proReq.LifecyclePreStop
		tx.Model(models.SgrTenantDeploymentsContainers{}).Where("deploy_id=? and is_main=1", proReq.DpId).Updates(mainContainer)

		tx.Commit()
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
	if dp.MaxUnavailable != nil && *dp.MaxUnavailable > 0 {
		res.MaxUnavailable = *dp.MaxUnavailable
	} else {
		res.MaxUnavailable = 25
	}
	if dp.MaxSurge != nil && *dp.MaxSurge > 0 {
		res.MaxSurge = *dp.MaxSurge
	} else {
		res.MaxSurge = 25
	}
	if dp.TerminationGracePeriodSeconds > 0 {
		res.TerminationGracePeriodSeconds = dp.TerminationGracePeriodSeconds
	} else {
		res.TerminationGracePeriodSeconds = 30
	}
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
	} else { // by default
		res.ReadinessReqScheme = "HTTP"
		res.ReadinessPort = dp.ServicePort
		res.ReadinessPeriodSeconds = 10
		res.ReadinessInitialDelaySeconds = 4
		res.ReadinessTimeoutSeconds = 3
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
	} else {
		res.LivenessReqScheme = "HTTP"
		res.LivenessPort = dp.ServicePort
		res.LivenessPeriodSeconds = 10
		res.LivenessInitialDelaySeconds = 4
		res.LivenessTimeoutSeconds = 3
	}
	return res, nil
}
