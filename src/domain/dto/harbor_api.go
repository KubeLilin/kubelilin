package dto

import "time"

type HarborProjectDTO struct {
	ChartCount         int       `json:"chart_count"`
	CreationTime       time.Time `json:"creation_time"`
	CurrentUserRoleId  int       `json:"current_user_role_id"`
	CurrentUserRoleIds []int     `json:"current_user_role_ids"`
	CveAllowlist       struct {
		CreationTime time.Time     `json:"creation_time"`
		Id           int           `json:"id"`
		Items        []interface{} `json:"items"`
		ProjectId    int           `json:"project_id"`
		UpdateTime   time.Time     `json:"update_time"`
	} `json:"cve_allowlist"`
	Metadata struct {
		Public string `json:"public"`
	} `json:"metadata"`
	Name       string    `json:"name"`
	OwnerId    int       `json:"owner_id"`
	OwnerName  string    `json:"owner_name"`
	ProjectId  int       `json:"project_id"`
	RepoCount  int       `json:"repo_count"`
	UpdateTime time.Time `json:"update_time"`
}

type CreateHarborProjectMetadata struct {
	Public string `json:"public"`
}

type CreateHarborProjectDTO struct {
	ProjectName  string                      `json:"project_name"`
	Metadata     CreateHarborProjectMetadata `json:"metadata"`
	StorageLimit int                         `json:"storage_limit"`
	RegistryId   interface{}                 `json:"registry_id"`
}
