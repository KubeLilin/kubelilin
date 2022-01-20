package tests

import (
	"fmt"
	gogs "github.com/gogs/go-gogs-client"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	gitClient  = gogs.NewClient("https://gogs.xiaocui.site/", "b472ae0baaeb86d5dddf5f4db6b4f06c13cd6949")
	gitClient2 = gogs.NewClient("https://gogs.xiaocui.site/", "d2911632a4ac8db4ce13a3135700b58a5c8d7772")
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
