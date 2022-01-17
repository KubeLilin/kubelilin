package tests

import (
	gogs "github.com/gogs/go-gogs-client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGogsCreateRepo(t *testing.T) {
	// token: b472ae0baaeb86d5dddf5f4db6b4f06c13cd6949
	gitClient := gogs.NewClient("https://gogs.xiaocui.site/", "b472ae0baaeb86d5dddf5f4db6b4f06c13cd6949")
	repo, err := gitClient.CreateOrgRepo("sgr_platform", gogs.CreateRepoOption{
		Name:        "testProject",
		Description: "first project for git",
		Private:     false,
	})

	assert.Equal(t, repo.CloneURL != "", true)

	assert.NoError(t, err)
}
