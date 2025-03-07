# K8s Pod Deployer
Interact with a K8s API to control the lifecycle of a pod

## Kind
```bash
kind create cluster
kubectl cluster-info \
--context kind-kind
```

## Methods
### POST
```bash
curl -X POST \
http://localhost:8080/pods \
-H "Content-Type: application/json" \
-d '{
  "apiVersion": "v1",
  "kind": "Pod",
  "metadata": {
    "name": "example-pod",
    "namespace": "default"
  },
  "spec": {
    "containers": [{
      "name": "example-container",
      "image": "nginx",
      "ports": [{
        "containerPort": 80
      }]
    }]
  }
}'
```
### GET
```bash
curl -X GET \
http://localhost:8080/pods/example-pod
```
### DELETE
```bash
curl -X DELETE \
http://localhost:8080/pods/example-pod 
```