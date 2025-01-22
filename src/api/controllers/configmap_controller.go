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

// NewConfigmapController 配置中心构造函数
func NewConfigmapController(db *gorm.DB, configMapSupervisor *kubernetes.ConfigMapSupervisor) *ConfigmapController {
	return &ConfigmapController{db: db, configMapSupervisor: configMapSupervisor}
}

// GetList 获取当前集群下的k8s中已经配置的 config-map
func (controller ConfigmapController) GetList(req *requests.ConfigMapPageReq) mvc.ApiResult {
	configmaps, err := controller.configMapSupervisor.QueryList(req)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(configmaps)
}

// GetListByAppId 根据选中的应用，获取应用中关联的 CONFIG-MAP
func (controller ConfigmapController) GetListByAppId(req *requests.ConfigMapPageReq) mvc.ApiResult {
	configmaps, err := controller.configMapSupervisor.GetListByAppId(req.AppId)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(configmaps)
}

// PostApply 根据提交的配置信息，在当前集群中创建 config-map
func (controller ConfigmapController) PostApply(configmap *requests.ConfigMap) mvc.ApiResult {
	err := controller.configMapSupervisor.Apply(configmap)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(true)
}

// GetConfigMap 根据前端传入的参数获取 config-map 中的配置信息
func (controller ConfigmapController) GetConfigMap(req *requests.ConfigMapPageReq) mvc.ApiResult {
	configmap, err := controller.configMapSupervisor.GetConfigMap(req)
	if err != nil {
		return mvc.Success("")
	}
	return mvc.Success(configmap)
}

// DeleteConfigMap 删除选中的 config-map
func (controller ConfigmapController) DeleteConfigMap(configmap *requests.ConfigMap) mvc.ApiResult {
	err := controller.configMapSupervisor.Delete(configmap)
	if err != nil {
		return mvc.Fail(err.Error())
	}
	return mvc.Success(true)
}
