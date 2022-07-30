# Podtatohead-operator for PodTatoHead installation on Kubernetes

## Steps to Install
 - Enter into root directory
### Install CRD
 - `kubectl apply -f https://raw.githubusercontent.com/oumkale/k8sOperator/main/config/crd/bases/crd.demo.com_podtatoheads.yaml`
### Install rbac
 - `kubectl apply -f https://raw.githubusercontent.com/oumkale/k8sOperator/main/config/operator/deployment.yaml`
### PodTatoHead CR
 - `kubectl apply -f https://raw.githubusercontent.com/oumkale/k8sOperator/main/config/CRmanifests/podtatohead.yaml`
### Check Resources in `operator` namespace
 - `kubectl get po -n operator`

## To test locally
 - Run Operator locally :
   Run `make generate` and `make manifests` commands everytime when we change in `api/v1alpha1` in root directory
   Note: Whenever we change in `api/v1alpha1 (types)` then run `kubectl delete -f config/crd/bases` and `kubectl create -f config/crd/bases`.
 - Apply crd : `kubectl apply -f config/crd/bases`
 - Run Operator : `make run`
 - `kubectl apply -f config/manifests/podtatohead.yaml -n oumk`
