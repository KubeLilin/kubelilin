package app

//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"github.com/drone/go-scm/scm"
//	"github.com/gogs/go-gogs-client"
//	"github.com/yoyofx/yoyogo/abstractions"
//	"gorm.io/gorm"
//	"kubelilin/domain/database/models"
//	"regexp"
//)
//
//const (
//	GIT_TOKEN = "vcs.git.token"
//	GIT_URL   = "vcs.git.url"
//)
//
//type GogsVcsService struct {
//	db     *gorm.DB
//	config abstractions.IConfiguration
//}
//
//func NewVcsService(db *gorm.DB, config abstractions.IConfiguration) *GogsVcsService {
//	return &GogsVcsService{
//		db:     db,
//		config: config,
//	}
//}
//
//func (vcs *GogsVcsService) CreateTenantRepository(tenantId uint64, repoName string) (*VcsRepository, error) {
//	tenant := models.SgrTenant{}
//	vcs.db.Model(models.SgrTenant{}).Where("id=?", tenantId).First(&tenant)
//	if tenant.TCode == "" {
//		return nil, errors.New("租户不存在")
//	}
//	//创建组织
//	orgRes, err := vcs.CreateOrganization(tenant.TCode)
//	if err != nil {
//		return nil, err
//	}
//	repoRes, repoErr := vcs.CreateRepositoryByOrg(orgRes.OrgName, repoName)
//	return repoRes, repoErr
//}
//
//func (vcs *GogsVcsService) CreateOrganization(orgName string) (*VcsOrganization, error) {
//	gitUrl := vcs.config.GetString(GIT_URL)
//	gitToken := vcs.config.GetString(GIT_TOKEN)
//	gogsClient := gogs.NewClient(gitUrl, gitToken)
//	existOrg, existErr := gogsClient.GetOrg(orgName)
//	if existOrg != nil {
//		return &VcsOrganization{
//			OrgName: existOrg.UserName,
//		}, nil
//	}
//	if existErr != nil {
//		return nil, existErr
//	}
//	orgRes, err := gogsClient.CreateOrg(gogs.CreateOrgOption{
//		UserName: orgName,
//		FullName: orgName,
//	})
//	if err != nil {
//		return &VcsOrganization{
//			OrgName: orgRes.UserName,
//		}, nil
//	} else {
//		return nil, err
//	}
//}
//
//func (vcs *GogsVcsService) CreateRepositoryByOrg(orgName, repoName string) (*VcsRepository, error) {
//	gitUrl := vcs.config.GetString(GIT_URL)
//	gitToken := vcs.config.GetString(GIT_TOKEN)
//	gogsClient := gogs.NewClient(gitUrl, gitToken)
//	repoRes, err := gogsClient.CreateOrgRepo(orgName, gogs.CreateRepoOption{
//		Name: repoName,
//	})
//	if err != nil {
//		return &VcsRepository{
//				Name:        repoRes.Name,
//				FullName:    repoRes.FullName,
//				Description: repoRes.Description,
//				Size:        repoRes.Size,
//				CloneURL:    repoRes.CloneURL,
//				Updated:     repoRes.Updated,
//				Created:     repoRes.Created,
//			},
//			nil
//	}
//	return nil, err
//}
//
//func (vcs *GogsVcsService) CreateRepository(repoName string) (*VcsRepository, error) {
//	gitUrl := vcs.config.GetString(GIT_URL)
//	gitToken := vcs.config.GetString(GIT_TOKEN)
//	gogsClient := gogs.NewClient(gitUrl, gitToken)
//	repoRes, err := gogsClient.CreateRepo(gogs.CreateRepoOption{
//		Name: repoName,
//	})
//	if err != nil {
//		return &VcsRepository{
//				Name:        repoRes.Name,
//				FullName:    repoRes.FullName,
//				Description: repoRes.Description,
//				Size:        repoRes.Size,
//				CloneURL:    repoRes.CloneURL,
//				Updated:     repoRes.Updated,
//				Created:     repoRes.Created,
//			},
//			nil
//	}
//	return nil, err
//}
//
//func (vcs *GogsVcsService) GetGitBranches(gitAddr string, sourceType string, gitToken string) ([]string, error) {
//	client, err := NewScmProvider(sourceType, gitAddr, gitToken)
//	branchList := []*scm.Reference{}
//	listOptions := scm.ListOptions{
//		Page: 1,
//		Size: 100,
//	}
//	repoRes, err := getRepoNames(gitAddr)
//	if err != nil {
//		return nil, err
//	}
//
//	got, _, err := client.Git.ListBranches(context.Background(), fmt.Sprintf("%s/%s", repoRes.OrganizationName, repoRes.RepositoryName), listOptions)
//	if err != nil {
//		return nil, err
//	}
//	branchList = append(branchList, got...)
//
//	var branchesNameList []string
//	for _, branch := range branchList {
//		branchesNameList = append(branchesNameList, branch.Name)
//	}
//	return branchesNameList, nil
//}
//
//type GitRepoNames struct {
//	OrganizationName string
//	RepositoryName   string
//}
//
//func getRepoNames(gitAddr string) (*GitRepoNames, error) {
//	reg := regexp.MustCompile("^http.*/(\\w+)/([a-zA-Z-0-9]+).git")
//	groups := reg.FindStringSubmatch(gitAddr)
//
//	if len(groups) > 1 {
//		return &GitRepoNames{
//			OrganizationName: groups[1],
//			RepositoryName:   groups[2],
//		}, nil
//	} else {
//		return nil, errors.New("not found")
//	}
//}
