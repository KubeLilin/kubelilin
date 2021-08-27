cd ../src

docker rmi yoyofx/sgr-api:v0.1.0

docker build -t yoyofx/sgr-api:v0.1.0 .

docker run --rm -p 8080:8080 yoyofx/sgr-api:v0.1.0
