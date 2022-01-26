package app

import (
	"errors"
	"github.com/gogs/go-gogs-client"
	"github.com/yoyofx/yoyogo/abstractions"
	"gorm.io/gorm"
	"sgr/domain/database/models"
)

const (
	GIT_TOKEN = "vcs.git.token"
	GIT_URL   = "vcs.git.url"
)

type GogsVcsService struct {
	db     *gorm.DB
	config abstractions.IConfiguration
}

func NewVcsService(db *gorm.DB, config abstractions.IConfiguration) *GogsVcsService {
	return &GogsVcsService{
		db:     db,
		config: config,
	}
}

func (vcs *GogsVcsService) CreateTenantRepository(tenantId uint64, repoName string) (*VcsRepository, error) {
	tenant := models.SgrTenant{}
	vcs.db.Model(models.SgrTenant{}).Where("id=?", tenantId).First(&tenant)
	if tenant.TCode == "" {
		return nil, errors.New("租户不存在")
	}
	//创建组织
	orgRes, err := vcs.CreateOrganization(tenant.TCode)
	if err != nil {
		return nil, err
	}
	repoRes, repoErr := vcs.CreateRepositoryByOrg(orgRes.OrgName, repoName)
	return repoRes, repoErr
}

func (vcs *GogsVcsService) CreateOrganization(orgName string) (*VcsOrganization, error) {
	gitUrl := vcs.config.GetString(GIT_URL)
	gitToken := vcs.config.GetString(GIT_TOKEN)
	gogsClient := gogs.NewClient(gitUrl, gitToken)
	existOrg, existErr := gogsClient.GetOrg(orgName)
	if existOrg != nil {
		return &VcsOrganization{
			OrgName: existOrg.UserName,
		}, nil
	}
	if existErr != nil {
		return nil, existErr
	}
	orgRes, err := gogsClient.CreateOrg(gogs.CreateOrgOption{
		UserName: orgName,
		FullName: orgName,
	})
	if err != nil {
		return &VcsOrganization{
			OrgName: orgRes.UserName,
		}, nil
	} else {
		return nil, err
	}
}

func (vcs *GogsVcsService) CreateRepositoryByOrg(orgName, repoName string) (*VcsRepository, error) {
	gitUrl := vcs.config.GetString(GIT_URL)
	gitToken := vcs.config.GetString(GIT_TOKEN)
	gogsClient := gogs.NewClient(gitUrl, gitToken)
	repoRes, err := gogsClient.CreateOrgRepo(orgName, gogs.CreateRepoOption{
		Name: repoName,
	})
	if err != nil {
		return &VcsRepository{
				Name:        repoRes.Name,
				FullName:    repoRes.FullName,
				Description: repoRes.Description,
				Size:        repoRes.Size,
				CloneURL:    repoRes.CloneURL,
				Updated:     repoRes.Updated,
				Created:     repoRes.Created,
			},
			nil
	}
	return nil, err
}

func (vcs *GogsVcsService) CreateRepository(repoName string) (*VcsRepository, error) {
	gitUrl := vcs.config.GetString(GIT_URL)
	gitToken := vcs.config.GetString(GIT_TOKEN)
	gogsClient := gogs.NewClient(gitUrl, gitToken)
	repoRes, err := gogsClient.CreateRepo(gogs.CreateRepoOption{
		Name: repoName,
	})
	if err != nil {
		return &VcsRepository{
				Name:        repoRes.Name,
				FullName:    repoRes.FullName,
				Description: repoRes.Description,
				Size:        repoRes.Size,
				CloneURL:    repoRes.CloneURL,
				Updated:     repoRes.Updated,
				Created:     repoRes.Created,
			},
			nil
	}
	return nil, err
}
