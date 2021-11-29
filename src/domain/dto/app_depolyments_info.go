package dto

type DeploymentItemDto struct {
	ID          uint64 `gorm:"column:id;" json:"id,omitempty"`
	Name        string `gorm:"column:name" json:"name"`
	NickName    string `gorm:"column:nickname" json:"nickname"`
	NameSpace   string `gorm:"column:namespace" json:"namespace"`
	ClusterName string `gorm:"column:clusterName" json:"clusterName"`
	Status      string `gorm:"column:status" json:"status"`
	LastImage   string `gorm:"column:lastImage" json:"lastImage"`
	Running     uint64 `gorm:"column:running" json:"running"`
	Expected    uint64 `gorm:"column:expected" json:"expected"`
	ServiceIP   string `gorm:"column:serviceIP"  json:"serviceIP"`
	ServiceName string `gorm:"column:serviceName" json:"serviceName"`
}
