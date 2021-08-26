cd ../src

docker rmi sgr-api

docker build -t sgr-api .

docker run --rm -p 8080:8080 sgr-api