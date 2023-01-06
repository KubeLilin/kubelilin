package requests

import "github.com/yoyofx/yoyogo/web/mvc"

type Volume struct {
	VolumeType string `json:"volumeType"`
	VolumeName string `json:"volumeName"`
	Value      string `json:"value"`
}

type VolumeMount struct {
	Volume      string `json:"volume"`
	DesPath     string `json:"desPath"`
	SubPathType string `json:"subPathType"`
	MountType   string `json:"mountType"`
}

type DeploymentVolume struct {
	mvc.RequestBody
	DeploymentID uint64        `json:"deployId"`
	Volume       []Volume      `json:"volumes"`
	VolumeMounts []VolumeMount `json:"volumeMounts"`
}
