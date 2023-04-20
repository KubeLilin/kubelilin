# Kubelilin
An Cloud-Native application platform for Kubernetes.
![](https://mnur-prod-public.oss-cn-beijing.aliyuncs.com/0/tech/physical_architecture.png)
![](https://mnur-prod-public.oss-cn-beijing.aliyuncs.com/0/tech/functional_architecture.png)
## Dapr运行时支持
![](https://mnur-prod-public.oss-cn-beijing.aliyuncs.com/0/tech/20230411143436.png)# Kubelilin ApiServer
Kubelilin cloud native backend .

## PaaS 平台运行
### Docker-Compose FrontEnd & Backend & MySQL
```bash
git clone https://github.com/KubeLilin/kubelilin.git
cd kubelilin/scripts
docker-compose  up -d
```
#### 登录
* 用户名: admin
* 密  码: 1234abcd

### FrontEnd
```bash
git clone https://github.com/KubeLilin/dashboard.git
cd dashboard/src
npm install --force
npm run start
```

# Todo:
## v0.2.0
1. [] 平台文档化
2. [] 服务治理 (注册中心，配置管理)
3. [] 应用&部署删除的审计记录
4. [] 流水线Action抽象扩展
5. [] 流水线镜像&包安全扫描Action
6. [] 流水线上传文件&制品库

## v0.1.8
1. [] helm商店&安装
2. [] 服务监控&图表展示

## v0.1.7
1. [] 从集群管理所有类型的工作负载 
2. [] pod监控 (promethus)

## v0.1.6
1. [x] 从git导入应用,并生成部署以及流水线
2. [x] 应用配置：部署绑定数据卷(卷、configmap,环境变量 )
3. [x] Dapr运行时集成

## v0.1.5
1. [x] 部署编辑改版使用tab展示配置
2. [x] 部署环境的健康检查功能
3. [x] 项目中创建及修改应用

