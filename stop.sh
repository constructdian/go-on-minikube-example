#!/bin/bash
kubectl delete -f k8s/deployments/ && \
kubectl delete -f k8s/services/