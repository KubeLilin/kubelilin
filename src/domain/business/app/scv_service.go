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
	GIT_TOKEN = "vcs.git.token"
	GIT_URL   = "vcs.git.url"
)

type VcsService struct {
	db     *gorm.DB
	config abstractions.IConfiguration
}

func NewVcsService(db *gorm.DB, config abstractions.IConfiguration) *VcsService {
	return &VcsService{
		db:     db,
		config: config,
	}
}

func (vcs *VcsService) CreateGitOrganizationByTenant(tenantId uint64) (*gogs.Organization, error) {
	tenant := models.SgrTenant{}
	vcs.db.Model(models.SgrTenant{}).Where("id=?", tenantId).First(&tenant)
	if tenant.TCode == "" {
		return nil, errors.New("租户不存在")
	}
	//创建组织
	orgRes, err := vcs.CreateGitOrganization(tenant.TCode)
	if err != nil {
		if !strings.Contains(err.Error(), "exists") {
			return nil, err
		}
	}
	return orgRes, nil
}

func (vcs *VcsService) CreateGitOrganization(orgName string) (*gogs.Organization, error) {
	gitUrl := vcs.config.GetString(GIT_URL)
	gitToken := vcs.config.GetString(GIT_TOKEN)
	gogsClient := gogs.NewClient(gitUrl, gitToken)
	orgRes, err := gogsClient.CreateOrg(gogs.CreateOrgOption{
		UserName: orgName,
		FullName: orgName,
	})
	return orgRes, err
}

func (vcs *VcsService) CreateGitRepository(repoName, orgName string) (*gogs.Repository, error) {
	gitUrl := vcs.config.GetString(GIT_URL)
	gitToken := vcs.config.GetString(GIT_TOKEN)
	gogsClient := gogs.NewClient(gitUrl, gitToken)
	repoRes, err := gogsClient.CreateOrgRepo(orgName, gogs.CreateRepoOption{
		Name: repoName,
	})
	return repoRes, err
}

func (vcs *VcsService) InitGitRepository(tenantId uint64, appName string) (string, error) {
	tenant := models.SgrTenant{}
	dberr := vcs.db.Model(models.SgrTenant{}).Where("id=?", tenantId).First(&tenant)
	if dberr.Error != nil {
		return "", dberr.Error
	}
	sb := strings.Builder{}
	sb.WriteString(vcs.config.GetString(GIT_URL))
	sb.WriteString("/")
	sb.WriteString(tenant.TCode)
	sb.WriteString("/")
	sb.WriteString(appName)
	gitUrl := sb.String()
	return gitUrl, nil
}
