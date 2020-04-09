#!/bin/bash
docker build -t helloworld:1.0 -f docker/Dockerfile . && \
kubectl apply -f k8s/deployments/ && \
kubectl apply -f k8s/services/
