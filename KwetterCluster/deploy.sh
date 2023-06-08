#!/bin/bash

services=(
  "gateway"
  "account"
  "auth"
  "tweet"
  "trend"
)

for service in "${services[@]}"; do
  kubectl apply -f KwetterCluster/"$service"/deployment.yaml
  kubectl apply -f KwetterCluster/"$service"/service.yaml
done