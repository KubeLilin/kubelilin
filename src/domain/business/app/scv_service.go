package app

import (
	"errors"
	"github.com/gogs/go-gogs-client"
	"github.com/yoyofx/yoyogo/abstractions"
	"gorm.io/gorm"
	"sgr/domain/database/models"
	"strings"
)

const (
	GIT_TOKEN = "scv.git.token"
	GIT_URL   = "scv.git.url"
)

type ScvService struct {
	db     *gorm.DB
	config abstractions.IConfiguration
}

func (scv *ScvService) CreateGitOrganizationByTenant(tenantId uint64) (*gogs.Organization, error) {
	tenant := models.SgrTenant{}
	scv.db.Model(models.SgrTenant{}).Where("id=?", tenantId).First(&tenant)
	if tenant.TCode == "" {
		return nil, errors.New("租户不存在")
	}
	//创建组织
	orgRes, err := scv.CreateGitOrganization(tenant.TCode)
	if err != nil {
		if !strings.Contains(err.Error(), "exists") {
			return nil, err
		}
	}
	return orgRes, nil
}

func (scv *ScvService) CreateGitOrganization(orgName string) (*gogs.Organization, error) {
	gitUrl := scv.config.GetString(GIT_URL)
	gitToken := scv.config.GetString(GIT_TOKEN)
	gogsClient := gogs.NewClient(gitUrl, gitToken)
	orgRes, err := gogsClient.CreateOrg(gogs.CreateOrgOption{
		UserName: orgName,
		FullName: orgName,
	})

	return orgRes, err
}

func (scv *ScvService) CreateGitRepository(repoName, orgName string) (*gogs.Repository, error) {
	gitUrl := scv.config.GetString(GIT_URL)
	gitToken := scv.config.GetString(GIT_TOKEN)
	gogsClient := gogs.NewClient(gitUrl, gitToken)
	repoRes, err := gogsClient.CreateOrgRepo(orgName, gogs.CreateRepoOption{
		Name: repoName,
	})
	return repoRes, err
}
