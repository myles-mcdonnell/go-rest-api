#!/bin/bash

rm -rf models/
rm -rf swagger.json
json-refs resolve --filter relative ../swagger/root.yml >> swagger.json
rm -rf client/
swagger generate client -f swagger.json -A Gorestapi