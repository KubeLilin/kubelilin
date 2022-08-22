package tests

import (
	"errors"
	"fmt"
	gogs "github.com/gogs/go-gogs-client"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

var (
	gitClient  = gogs.NewClient("https://gogs/", "xxxxxxddf5f4db6b4f06c13cd6949")
	gitClient2 = gogs.NewClient("https://gogs/", "xxxxxxdb4ce13a3135700b58a5c8d7772")
)

func TestGogsCreateRepo(t *testing.T) {
	repo1, err := gitClient.GetRepo("sgr_platform", "testProject")
	_ = repo1

	repo, err := gitClient.CreateOrgRepo("sgr_platform", gogs.CreateRepoOption{
		Name:        "testProject",
		Description: "first project for git",
		Private:     false,
	})

	assert.Equal(t, repo.CloneURL != "", true)

	assert.NoError(t, err)
}

func TestGogsBranches(t *testing.T) {
	bs, err := gitClient.ListRepoBranches("sgr_platform", "testProject")

	assert.Equal(t, len(bs) >= 2, true)
	assert.NoError(t, err)
}

func TestCreateOrg(t *testing.T) {

	res, err := gitClient2.CreateOrg(gogs.CreateOrgOption{
		UserName: "sgr_666",
		FullName: "sgr_777",
	})
	fmt.Println(err)
	fmt.Println(res)
	assert.NoError(t, err)
}

func TestGetRegexpLibrary(t *testing.T) {
	git := "https://gihub.com/administration/nginx.git"

	names, _ := GetRepoNames(git)
	fmt.Println(names.OrganizationName, names.RepositoryName)

	bs, err := gitClient.ListRepoBranches(names.OrganizationName, names.RepositoryName)

	assert.Equal(t, len(bs) >= 2, true)
	assert.NoError(t, err)
}

type GitRepoNames struct {
	OrganizationName string
	RepositoryName   string
}

func GetRepoNames(gitAddr string) (*GitRepoNames, error) {
	reg := regexp.MustCompile("^http.*/(\\w+)/(\\w+).git")
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
