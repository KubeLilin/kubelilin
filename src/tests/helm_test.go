package tests

import (
	"fmt"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
	"log"
	"testing"
)

func TestHelm(t *testing.T) {
	repoUrl := "https://charts.bitnami.com/bitnami"

	// 创建一个 Chart 存储库配置
	chartRepo := repo.Entry{
		URL:  repoUrl,
		Name: "stable",
	}

	// 创建一个 Chart 存储库客户端
	providers := getter.All(&cli.EnvSettings{})
	chartRepoClient, err := repo.NewChartRepository(&chartRepo, providers)
	if err != nil {
		log.Fatal(err)
	}

	// 更新 Chart 存储库索引
	path, err := chartRepoClient.DownloadIndexFile()
	if err != nil {
		log.Fatal(err)
	}
	// 加载 Chart 存储库索引
	indexFile, err := repo.LoadIndexFile(path)
	if err != nil {
		log.Fatal(err)
	}
	// 搜索特定的 Chart
	//searchTerm := "nginx"
	for chartName, chartVersions := range indexFile.Entries {
		//if strings.Contains(chartName, searchTerm) {
		fmt.Println("Chart Name:", chartName)
		for _, chartVersion := range chartVersions {
			fmt.Println("  Version:", chartVersion.Version)
		}
		//}
	}
}
