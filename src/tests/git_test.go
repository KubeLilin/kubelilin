package tests

import (
	"context"
	"fmt"
	"github.com/drone/go-scm/scm"
	"kubelilin/domain/business/app"
	"testing"
)

func TestGitList(t *testing.T) {
	gitAddr := "http://gitlab.yoyogo.run/yoyofx/dapr-demo.git"
	client, _ := app.NewScmProvider("gitlab", gitAddr, "")
	repoRes, _ := app.GetRepoNames(gitAddr)

	//branchList := []*scm.Reference{}
	//listOptions := scm.ListOptions{
	//	Page: 1,
	//	Size: 100,
	//}
	//got, _, _ := client.Git.ListBranches(context.Background(), fmt.Sprintf("%s/%s", repoRes.OrganizationName, repoRes.RepositoryName), listOptions)

	commits, _, _ := client.Git.ListCommits(context.Background(), fmt.Sprintf("%s/%s", repoRes.OrganizationName, repoRes.RepositoryName), scm.CommitListOptions{Size: 1})
	latestCommit := commits[0]
	fmt.Printf("Latest commit Message %s SHA: %s\n", latestCommit.Message, latestCommit.Sha)

	files, err := app.FindFiles(gitAddr, "gitlab", "", "main", "Dockerfile")
	if err != nil {
		return
	}
	fmt.Println(files)
}
