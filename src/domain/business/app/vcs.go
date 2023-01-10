package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/gitee"
	"github.com/drone/go-scm/scm/driver/github"
	"github.com/drone/go-scm/scm/driver/gitlab"
	"github.com/drone/go-scm/scm/driver/gogs"
	"github.com/drone/go-scm/scm/transport"
	"net/http"
	"regexp"
	"strings"
	"time"
)

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

func GetGitBranches(gitAddr string, sourceType string, gitToken string) ([]string, error) {
	client, err := NewScmProvider(sourceType, gitAddr, gitToken)
	branchList := []*scm.Reference{}
	listOptions := scm.ListOptions{
		Page: 1,
		Size: 100,
	}
	repoRes, err := getRepoNames(gitAddr)
	if err != nil {
		return nil, err
	}

	got, _, err := client.Git.ListBranches(context.Background(), fmt.Sprintf("%s/%s", repoRes.OrganizationName, repoRes.RepositoryName), listOptions)
	if err != nil {
		return nil, err
	}
	branchList = append(branchList, got...)

	var branchesNameList []string
	for _, branch := range branchList {
		branchesNameList = append(branchesNameList, branch.Name)
	}
	return branchesNameList, nil
}

// NewScmProvider ..
func NewScmProvider(vcsType, vcsPath, token string) (*scm.Client, error) {
	var err error
	var client *scm.Client
	switch strings.ToLower(vcsType) {
	case "gogs", "gitlab":
		if strings.HasSuffix(vcsPath, ".git") {
			vcsPath = strings.TrimSuffix(vcsPath, ".git")
		}

		vcsPathSplit := strings.Split(vcsPath, "://")
		// TODO: verify vcsPath, only support http, do not support git@gitlab.com:/dddd.git
		projectPathSplit := strings.Split(vcsPathSplit[1], "/")
		//projectName := strings.Join(projectPathSplit[1:], "/")
		schema := vcsPathSplit[0]
		gitRepo := strings.ToLower(vcsType)
		if "gogs" == gitRepo {
			client, err = gogs.New(schema + "://" + projectPathSplit[0])
			client.Client = &http.Client{}
			if token != "" {
				client.Client.Transport = &transport.BearerToken{Token: token}
			}
		} else {
			client, err = gitlab.New(schema + "://" + projectPathSplit[0])
			client.Client = &http.Client{}
			if token != "" {
				client.Client.Transport = &transport.PrivateToken{Token: token}
			}
		}
	case "github":
		client = github.NewDefault()
		client.Client = &http.Client{}
		if token != "" {
			client.Client.Transport = &transport.BearerToken{Token: token}
		}

	case "gitee":
		client = gitee.NewDefault()
		client.Client = &http.Client{}
		if token != "" {
			client.Client.Transport = &transport.BearerToken{Token: token}
		}

	default:
		err = fmt.Errorf("source code management system not configured")
	}
	return client, err
}

func getRepoNames(gitAddr string) (*GitRepoNames, error) {
	reg := regexp.MustCompile("^http.*/(\\w+)/([a-zA-Z-0-9.]+).git")
	groups := reg.FindStringSubmatch(gitAddr)

	if len(groups) > 1 {
		return &GitRepoNames{
			OrganizationName: groups[1],
			RepositoryName:   groups[2],
		}, nil
	} else {
		return nil, errors.New("not found")
	}
}

type GitRepoNames struct {
	OrganizationName string
	RepositoryName   string
}
