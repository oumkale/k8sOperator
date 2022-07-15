# Podtatohead-operator for PodTatoHead installation on Kubernetes


## To test locally
 - Run demo application in diffrent terminal `/demo_application`
   Command to run : `go run main.go`
 - Run Operator locally :
   Run `make generate` and `make manifests` commands everytime when we change in `api/v1alpha1` in root directory
   Note: Whenever we change in `api/v1alpha1 (types)` then run `kubectl delete -f config/crd/bases` and `kubectl create -f config/crd/bases`.
 - Apply crd : `kubectl apply -f config/crd/bases`
 - Run Operator : `make run`
   Test things are working fine

## [Not Required]
## Steps to Install
 - Enter into root directory
### Install CRD
 - `kubectl apply -f config/crd`
### Install rbac
 - `kubectl apply -f config/rbac-manifest -n oumk` (`oumk` namespace hard coded in rbac manifests please update if you using own namespace)
 ### For Operator deployment
 - `kubectl apply -f config/operator/deployment.yaml -n oumk`
### For Base CR and PodTatoHead CR
 - `kubectl apply -f config/manifests/base.yaml -n oumk` : Make sure, correct encoded format of systemYaml and keys updated to base.yaml
    - To encode `echo "<systemyaml file data>" | base64`
 - `kubectl apply -f config/manifests/podtatohead.yaml -n oumk`
