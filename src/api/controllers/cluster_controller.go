package controllers

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"gorm.io/gorm"
	"sgr/domain/business/kubernetes"
)

type ClusterController struct {
	mvc.ApiController

	db *gorm.DB
}

func NewClusterController(database *gorm.DB) *ClusterController {
	return &ClusterController{db: database}
}

func (controller ClusterController) GetPods() mvc.ApiResult {
	client, err := kubernetes.NewClientSet("")
	if err != nil {
		panic(err)
	}

	podList := kubernetes.GetPodList(client, "yoyogo")

	return controller.OK(podList)
}
