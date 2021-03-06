# 镜像
* Jenkins Server:  jenkins/jenkins:2.328
* Jenkins Slave :  jenkins/inbound-agent:4.10-3

# 1.依赖插件
* Kubernetes
* Pipeline
* Blue Ocean
* HTTP Request Plugin

#2. 开启匿名用户的可读权限
``为了保证 流水线的 日志详情可以正常打开
``
![img.png](images/img.png)

#3. 配置Kubernetes cloud
* 创建RABC账号ServiceAccount
* 注意默认授权的命名空间 **kube-lilin** .
## 获取ServiceAccount Auth
```bash
$ kubectl -n devops describe serviceaccount jenkins-admin
$ kubectl -n devops describe secret [jenkins-admin-token-name]
```
## 创建 Secret text 类型的Credentials
![img_1.png](images/img_1.png)

![img_2.png](images/img_2.png)

## Jenkins URL
配置Jenkins url为对应 {ServiceName}.{Namespace}:{TargetPort}

例如:  http://jenkins.kube-lilin:8080/

注意: 这里使用的Kubernetes Namespace 注意要和创建的 ServiceAccount的 Namespace 保持一致。


## Jenkins内部构建镜像
### 使用kaniko镜像构建Docker Image
```json
{
  Name:       "docker",
  Image:      "yoyofx/kaniko-executor:debug",
  WorkingDir: "/home/jenkins/agent",
  CommandArr: []string{"cat"},
}
```
#### Docker Hub Auth信息
需要将Auth信息打包到/kaniko/.docker/config.json 文件中。
```bash
## 环境变量,这里以 hub.docker.com 为例:
{Key: "SGR_REGISTRY_ADDR", Value: "https://index.docker.io/v1/"},
{Key: "SGR_REGISTRY_AUTH", Value: "eWxxxxxxOnpsMTI1MzMwMw=="},
{Key: "SGR_REGISTRY_CONFIG", Value: "/kaniko/.docker"},

echo '{"auths": {"'$SGR_REGISTRY_ADDR'": {"auth": "'$SGR_REGISTRY_AUTH'"}}}' > $SGR_REGISTRY_CONFIG/config.json

```
Auth信息生成($SGR_REGISTRY_AUTH)：
```bash
echo -n 用户名:密码 | base64
```

# 目前支持的阶段
Checkout (串行) -> Compile (并行) -> Build (并行) -> Deploy (并行)