package kubernetes

import (
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
func (ds *ProBeService) CreateProBe(proReq *requests.ProbeRequest) {
	ds.db.Transaction(func(tx *gorm.DB) error {
		if proReq.EnableReadiness {
			probe := models.SgrDeploymentProbe{}
			probe.DpID = proReq.DpId
			probe.Type = proReq.ReadinessType
			probe.Port = proReq.ReadinessPort
			probe.Type = READINESS
			probe.Path = proReq.ReadinessUrl
			probe.ReqScheme = proReq.ReadinessReqScheme
			probe.PeriodSeconds = proReq.ReadinessPeriodSeconds
			probe.InitialDelaySeconds = proReq.ReadinessInitialDelaySeconds
			err := tx.Model(models.SgrDeploymentProbe{}).Save(&probe).Error
			return err
		}
		if proReq.EnableLiveness {
			probe := models.SgrDeploymentProbe{}
			probe.DpID = proReq.DpId
			probe.Type = proReq.LivenessType
			probe.Port = proReq.LivenessPort
			probe.Path = proReq.LivenessUrl
			probe.Type = LIVENESS
			probe.ReqScheme = proReq.LivenessReqScheme
			probe.PeriodSeconds = proReq.LivenessPeriodSeconds
			probe.InitialDelaySeconds = proReq.LivenessInitialDelaySeconds
			err := tx.Save(&probe).Error
			return err
		}
		tx.Model(models.SgrTenantDeployments{}).Update("termination_grace_period_seconds=?", proReq.TerminationGracePeriodSeconds).Where("id=?", proReq.DpId)
		tx.Model(models.SgrTenantDeploymentsContainers{}).Update("poststart=? ", proReq.LifecyclePreStart).Where("deploy_id=? and is_main=1", proReq.DpId)
		tx.Model(models.SgrTenantDeploymentsContainers{}).Update(" podstop=?", proReq.LifecyclePreStop).Where("deploy_id=? and is_main=1", proReq.DpId)
		return nil
	})
}
