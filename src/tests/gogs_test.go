package tests

import (
	gogs "github.com/gogs/go-gogs-client"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	gitClient = gogs.NewClient("https://gogs.xiaocui.site/", "b472ae0baaeb86d5dddf5f4db6b4f06c13cd6949")
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
