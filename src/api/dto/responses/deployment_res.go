package responses

import "time"

type DeploymentReleaseRecordRes struct {
	ID             uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	AppID          uint64     `gorm:"column:app_id;type:bigint unsigned;not null" json:"appId"`
	DeploymentID   uint64     `gorm:"column:deployment_id;type:bigint unsigned;not null" json:"deploymentId"`
	Level          string     `gorm:"column:level;"`
	DeploymentName string     `gorm:"column:deployment_name;"`
	ApplyImage     string     `gorm:"column:apply_image;type:varchar(255);not null" json:"applyImage"`
	OpsType        string     `gorm:"column:ops_type;type:char(20);not null" json:"opsType"`
	Operator       *uint64    `gorm:"column:operator;type:bigint unsigned" json:"operator"`
	OperatorName   string     `gorm:"column:operator_name;" json:"operatorName"`
	CreationTime   *time.Time `gorm:"column:creation_time;type:datetime" json:"creationTime"`
}

type DeploymentProbeInfo struct {
	DpId                          uint64 `json:"dpId"`
	EnableLifecycle               bool   `json:"enableLifecycle"`
	EnableReadiness               bool   `json:"enableReadiness"`
	EnableLiveness                bool   `json:"enableLiveness"`
	TerminationGracePeriodSeconds uint64 `json:"terminationGracePeriodSeconds"`
	LifecyclePreStop              string `json:"lifecyclePreStop"`
	LifecyclePreStart             string `json:"lifecyclePreStart"`
	ReadinessType                 string `json:"readinessType"`
	ReadinessPort                 uint   `json:"readinessPort"`
	ReadinessUrl                  string `json:"readinessUrl"`
	ReadinessReqScheme            string `json:"readinessReqScheme"`
	ReadinessInitialDelaySeconds  uint   `json:"readinessInitialDelaySeconds"`
	ReadinessTimeoutSeconds       uint64 `json:"readinessTimeoutSeconds"`
	ReadinessPeriodSeconds        uint   `json:"readinessPeriodSeconds"`
	LivenessType                  string `json:"livenessType"`
	LivenessPort                  uint   `json:"livenessPort"`
	LivenessUrl                   string `json:"livenessUrl"`
	LivenessReqScheme             string `json:"livenessReqScheme"`
	LivenessInitialDelaySeconds   uint   `json:"livenessInitialDelaySeconds"`
	LivenessTimeoutSeconds        uint64 `json:"livenessTimeoutSeconds"`
	LivenessPeriodSeconds         uint   `json:"livenessPeriodSeconds"`
}
