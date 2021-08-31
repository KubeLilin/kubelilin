# Postman自动化测试

## 安装
```bash
 sudo npm install newman --global
```

## 自动化
```bash
# newman run sgr.postman_collection.json  -e postman_prod.json
newman run sgr.postman_collection.json  -e postman_local.json
```