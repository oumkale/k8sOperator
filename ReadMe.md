# Podtatohead-operator for PodTatoHead installation on Kubernetes

## Steps to Install
 - Enter into root directory
### Install CRD
 - `kubectl apply -f config/crd`
### Install rbac
 - `kubectl apply -f config/rbac-manifest -n oumk` (`oumk` namespace hard coded in rbac manifests please update if you using own namespace)
 ### For Operator deployment
 - `kubectl apply -f config/operator/deployment.yaml -n oumk`
### PodTatoHead CR
 - `kubectl apply -f config/manifests/podtatohead.yaml -n oumk`

## To test locally
 - Run Operator locally :
   Run `make generate` and `make manifests` commands everytime when we change in `api/v1alpha1` in root directory
   Note: Whenever we change in `api/v1alpha1 (types)` then run `kubectl delete -f config/crd/bases` and `kubectl create -f config/crd/bases`.
 - Apply crd : `kubectl apply -f config/crd/bases`
 - Run Operator : `make run`
 - `kubectl apply -f config/manifests/podtatohead.yaml -n oumk`
