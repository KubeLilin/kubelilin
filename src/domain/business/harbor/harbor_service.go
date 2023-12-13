package harbor

import (
	"encoding/json"
	"fmt"
	"github.com/yoyofx/yoyogo/pkg/httpclient"
	"kubelilin/domain/database/models"
	"kubelilin/domain/dto"
)

func CreateProject(projectName string, harborDetail models.ServiceConnectionDetails) {
	reqData := dto.CreateHarborProjectDTO{
		ProjectName:  projectName,
		StorageLimit: -1,
		Metadata: dto.CreateHarborProjectMetadata{
			Public: "true",
		},
	}
	var scd dto.ServiceConnectionDetails
	json.Unmarshal([]byte(harborDetail.Detail), &scd)
	jsonStr, _ := json.Marshal(reqData)
	client := httpclient.NewClient()
	requestEntity := httpclient.Request{}

	requestEntity.POST(scd.Repo)

	requestEntity.WithBody(string(jsonStr))
	requestEntity.ContentType("application/json")
	resp, err := client.Post(&requestEntity)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}
