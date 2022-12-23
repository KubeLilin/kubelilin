package controllers

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"gorm.io/gorm"
	"kubelilin/api/dto/requests"
	"kubelilin/domain/business/kubernetes"
)

type ConfigmapController struct {
	mvc.ApiController
	db                  *gorm.DB
	configMapSupervisor *kubernetes.ConfigMapSupervisor
}

func NewConfigmapController(db *gorm.DB, configMapSupervisor *kubernetes.ConfigMapSupervisor) *ConfigmapController {
	return &ConfigmapController{db: db, configMapSupervisor: configMapSupervisor}
}

func (controller ConfigmapController) GetList(req *requests.ConfigMapPageReq) mvc.ApiResult {
	configmaps, err := controller.configMapSupervisor.QueryList(req)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(configmaps)
}

func (controller ConfigmapController) GetListByAppId(req *requests.ConfigMapPageReq) mvc.ApiResult {
	configmaps, err := controller.configMapSupervisor.GetListByAppId(req.AppId)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(configmaps)
}

func (controller ConfigmapController) PostApply(configmap *requests.ConfigMap) mvc.ApiResult {
	err := controller.configMapSupervisor.Apply(configmap)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(true)
}

func (controller ConfigmapController) GetConfigMap(req *requests.ConfigMapPageReq) mvc.ApiResult {
	configmap, err := controller.configMapSupervisor.GetConfigMap(req)
	if err != nil {
		return mvc.Success("")
	}
	return mvc.Success(configmap)
}
