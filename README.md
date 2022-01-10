# sgr-platform-api
SGR cloud native backend .

## PaaS 平台运行
### Backend & MySQL
```bash
git clone https://github.com/yoyofxteam/sgr-platform-api.git
cd sgr-platform-api/scripts
docker-compose -f ./docker-compose.yaml up -d
```
### FrontEnd
```bash
git clone https://github.com/yoyofxteam/sgr-platform-ui.git
cd sgr-platform-ui/src
npm install
npm run start
```

# Todo:
## fix bugs
-[X] 为用户分配角色
-[ ] 导入集群时，创建集群对应的namespace （一个租户 对应一个集群一个namespace 1:N ）
-[ ] namespace分配配额 （cpu、内存、磁盘、pod）等

## 流水线(jenkins)
流水线（应用级） 一个流水线可以有多个应用
1. 编译过程BUILD：
-  git pull (应用中保存的git)
-  编译环境 (java 、go 、 nodejs) + 命令行
-  docker build xxxxxx .
-  docker push  (  "hub.yoyogo.run/" + "{应用名+部署环境名+}" + ":build.number"  )

2. 部署DEPLOY
-  部署 k8s Deployment (应用+部署环境 确定)
-  生成 k8s service (应用+部署环境 确定)
-  网关绑定 apixsix admin api ,添加默认路由并绑定 k8s service name

--------
    服务治理：（不在本期范围）
    配置管理：（不在本期范围）
    
    运维中心：（不在本期范围）
    链路分析
    服务监控
    日志服务
    注册中心： 导入 nacos 信息
    
    第一期就是管理中心和集群管理，其它不做

## 资源大盘
没想好 显示点啥


# 已完成功能：
## 管理中心(SAAS功能)  （已完成）
* 角色管理
* 用户管理
* 权限管理： (用户 角色 菜单)
## 资源中心：   （已完成）
### 集群管理：  （已完成）
 功能：管理k8s集群，导入k8sconfig配置，添加新配置到集群列表。  
### 集群列表    （已完成）
集群信息 id ,状态，版本号，运行时，描述，网络，配置 ， api server 地址 k8s config 等。
#### 下钻菜单  （已完成）
1. 节点管理列表: 显示节点ID，名称，IP，状态，已分配/总资源 等，(CPU，内存使用量)。
### 命名空间：管理对应K8s的命名空间，附加一些描述信息 （已完成）
### 部署环境管理： （已完成）
部署环境是一种对应k8s部署的描述信息 ，一个部署对应一个集群的一个命名空间( k8s:  id & namespace ) ，和一组容器 工作负载的配置信息 ( pod 的 yaml )
#### 部署环境列表  （已完成）
部署名称，所在集群，集群下的命名空间，运行状态，容器镜像，备注，运行中/预期实例数， 负载均衡IP/服务IP, 更新时间, 操作()
* 新建部署        （已完成）
* 操作子菜单      （已完成）
   ##### 部署应用  （已完成）
   ##### 删除部署   （已完成）
   ##### 部署页面   （已完成）
     用于实际部署到可运行环境中配置，根据部署中的配置信息生成 K8s的部署资源并apply. (信息包括：部署名，镜像/版本，启动参数，环境变量(多), 更新方式，健康检查，资源配置cpu,内存限制，实例数等)
   ##### 扩缩容     （已完成）
     控制实例数量
   ##### 清空实例    （已完成）
     实例清0
## 应用中心：  （已完成）
* 应用管理：   应用显示出多个部署的运行状态信息  （已完成）
1. 新建应用： 名称 ， git地址， 所在部门 ，应用 只是一个概念，它只记录，一些基础信息，比如， 应用的级别 ， 来源(git地址) ，项目描述，开发语言等信息
2. 应用列表： (列表) id，应用名称, 部署环境数量， 运行实例数/总实例数  , 操作(删除)
   * 应用名称下钻：
   1. 部署环境列表：
        ID/部署环境名，集群，命名空间，运行状态，镜像/版本 ，运行实例数/预期实例数， 负载均衡IP/服务IP, 更新时间,
        应用对应的部署环境列表(同部署环境列表)
   2. 容器列表, POD -> Container