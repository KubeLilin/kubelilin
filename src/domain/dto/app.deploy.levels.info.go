package dto

type DeployLeveLCountInfo struct {
	Label string `json:"label" gorm:"column:label"`
	Value string `json:"value" gorm:"column:value"`
	Count int    `json:"count" gorm:"column:count"`
}

type TeamLeveLCountInfo struct {
	Label string `json:"label" gorm:"column:label"`
	Value string `json:"value" gorm:"column:value"`
	Count int    `json:"count" gorm:"column:count"`
}
