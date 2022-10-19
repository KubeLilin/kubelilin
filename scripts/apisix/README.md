# 集群内部署APISIX
## 修改配置
在apisix.yaml中找到service配置成NodePort,端口只要与主机端口不冲突即可。

记录下文件中 admin_key的值，用于记录网关认证。
## 安装 
* kubectl apply -f ./etcd.yaml
* kubectl apply -f ./apisix.yaml
* kubectl apply -f ./apisix-dashboard.yaml

## Nginx Or Caddy 代理
使用反代服务将主机NodePort映射出来,这里反代IP即为网关的出口IP,可以绑定域名使用。

## 网关登记
PaaS数据库中,application_api_gateway表,新添加一条记录:
* access_token: {admin_key}
* admin_url:绑定的域名
* cluster_id:PaaS集群IP
* vip:内网IP
* cluster_ip:出口IP