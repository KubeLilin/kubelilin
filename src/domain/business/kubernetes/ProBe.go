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
		tx.Model(models.SgrTenantDeployments{}).Update("termination_grace_period_seconds=?", proReq.TerminationGracePeriodSeconds).Where("id=?", proReq.DpId)
		tx.Model(models.SgrTenantDeploymentsContainers{}).Update("poststart=? ", proReq.LifecyclePreStart).Where("deploy_id=? and is_main=1", proReq.DpId)
		tx.Model(models.SgrTenantDeploymentsContainers{}).Update(" podstop=?", proReq.LifecyclePreStop).Where("deploy_id=? and is_main=1", proReq.DpId)
		return nil
	})
	return res

}
