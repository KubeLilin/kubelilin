# sudo npm install newman --global

# newman run postman_collection.json  -e postman_prod.json

newman run kubelilin.postman_collection.json -e postman_local.json
