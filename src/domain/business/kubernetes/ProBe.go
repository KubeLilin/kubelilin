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
func (pbs *ProBeService) CreateProBe(proReq *requests.ProbeRequest) {
	// 判断配置是否已经存在
	exitsCount := int64(0)
	pbs.db.Model(models.DeploymentContainerLifecycleCheck{}).Where("deploymentId=?", proReq.DpId).Count(&exitsCount)
	if exitsCount > 0 {
		pbs.db.Transaction(func(tx *gorm.DB) error {
			readiness := &models.DeploymentContainerLifecycleCheck{}
			pbs.db.Model(models.DeploymentContainerLifecycleCheck{}).Where("deploymentId=? and type=? ", proReq.DpId, READINESS).First(readiness)
			if readiness != nil {
				readiness.Port = proReq.ReadinessPort
				readiness.Path = proReq.ReadinessUrl
				readiness.Scheme = proReq.ReadinessReqScheme
				readiness.PeriodSeconds = proReq.ReadinessPeriodSeconds
				readiness.InitialDelaySeconds = proReq.ReadinessInitialDelaySeconds
				err := tx.Model(models.DeploymentContainerLifecycleCheck{}).Updates(&readiness).Error
				return err
			}
			liveness := &models.DeploymentContainerLifecycleCheck{}
			pbs.db.Model(models.DeploymentContainerLifecycleCheck{}).Where("deploymentId=? and type=? ", proReq.DpId, LIVENESS).First(liveness)
			if liveness != nil {
				liveness.Port = proReq.LivenessPort
				liveness.Path = proReq.LivenessUrl
				liveness.Scheme = proReq.LivenessReqScheme
				liveness.PeriodSeconds = proReq.LivenessPeriodSeconds
				liveness.InitialDelaySeconds = proReq.LivenessInitialDelaySeconds
				err := tx.Model(models.DeploymentContainerLifecycleCheck{}).Updates(&liveness).Error
				return err
			}
			return nil
		})
	} else {
		pbs.db.Transaction(func(tx *gorm.DB) error {
			mainContainer := &models.SgrTenantDeploymentsContainers{}
			tx.Model(models.SgrTenantDeploymentsContainers{}).Where("deploy_id=? and is_main=1", proReq.DpId).First(mainContainer)
			if mainContainer == nil {
				return errors.New("can't find the sole container of development ")
			}

			if proReq.EnableReadiness {
				probe := models.DeploymentContainerLifecycleCheck{}
				probe.DeploymentID = proReq.DpId
				probe.Type = proReq.ReadinessType
				probe.Port = proReq.ReadinessPort
				probe.Type = READINESS
				probe.Path = proReq.ReadinessUrl
				probe.Scheme = proReq.ReadinessReqScheme
				probe.PeriodSeconds = proReq.ReadinessPeriodSeconds
				probe.InitialDelaySeconds = proReq.ReadinessInitialDelaySeconds
				err := tx.Model(models.DeploymentContainerLifecycleCheck{}).Save(&probe).Error
				return err
			}
			if proReq.EnableLiveness {
				probe := models.DeploymentContainerLifecycleCheck{}
				probe.DeploymentID = proReq.DpId
				probe.Type = LIVENESS
				probe.Port = proReq.LivenessPort
				probe.Path = proReq.LivenessUrl
				probe.Type = LIVENESS
				probe.Scheme = proReq.LivenessReqScheme
				probe.PeriodSeconds = proReq.LivenessPeriodSeconds
				probe.InitialDelaySeconds = proReq.LivenessInitialDelaySeconds
				err := tx.Model(models.DeploymentContainerLifecycleCheck{}).Save(&probe).Error
				return err
			}
			tx.Model(models.SgrTenantDeployments{}).Update("termination_grace_period_seconds=?", proReq.TerminationGracePeriodSeconds).Where("id=?", proReq.DpId)
			tx.Model(models.SgrTenantDeploymentsContainers{}).Update("poststart=? ", proReq.LifecyclePreStart).Where("deploy_id=? and is_main=1", proReq.DpId)
			tx.Model(models.SgrTenantDeploymentsContainers{}).Update(" podstop=?", proReq.LifecyclePreStop).Where("deploy_id=? and is_main=1", proReq.DpId)
			return nil
		})
	}
}
