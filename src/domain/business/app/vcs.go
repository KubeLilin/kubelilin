package app

import "time"

type VcsService interface {
	CreateOrganization(orgName string) (*VcsOrganization, error)
	CreateRepository(repoName string) (*VcsRepository, error)
	CreateRepositoryByOrg(orgName, repoName string) (*VcsRepository, error)
	CreateTenantRepository(tenantId uint64, repoName string) (*VcsRepository, error)
	GetGitBranches(gitAddr string, sourceType string, gitToken string) ([]string, error)
}

type VcsOrganization struct {
	OrgName string `json:"orgName"`
}

type VcsRepository struct {
	Name        string    `json:"name"`
	FullName    string    `json:"fullName"`
	Description string    `json:"description"`
	Size        int64     `json:"size"`
	CloneURL    string    `json:"clone_url"`
	Created     time.Time `json:"created_at"`
	Updated     time.Time `json:"updated_at"`
}
