package harbor

import (
	"kubelilin/domain/dto"
	"kubelilin/utils"
)

type HarborService struct {
}

func (svc *HarborService) CreateProject(projectName string) {
	reqData := dto.CreateHarborProjectDTO{
		ProjectName:  projectName,
		StorageLimit: -1,
		Metadata: dto.CreateHarborProjectMetadata{
			Public: "true",
		},
	}
	utils.PostHttpMessage()
}
