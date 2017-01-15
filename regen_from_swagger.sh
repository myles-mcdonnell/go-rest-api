#!/bin/bash

rm -rf swagger.json
json-refs resolve --filter relative ./swagger/root.yml >> swagger.json
rm -rf restapi/
rm -rf models/
swagger generate server -A go-rest-api -f swagger.json --model-package=models
cp swagger_post_autogen/configure_go_rest.go restapi/
cp swagger_post_autogen/initialiser.go restapi/

cd integration_tests
./regen_client.sh
cd ..