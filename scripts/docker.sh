# docker-compose 运行 sgr-api & mysql
docker-compose -f ./docker-compose.yaml up

docker-compose build --no-cache



## dev
docker build -t yoyofx/sgr-api:dev1.0 -f .\Dockerfile_Prod .




# 单独运行s gr-api
cd ../src

docker rmi yoyofx/sgr-api:v0.1.0

docker build -t yoyofx/sgr-api:v0.1.0 .

docker run --rm -p 8080:8080 yoyofx/sgr-api:v0.1.0



