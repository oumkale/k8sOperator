# Podtatohead-operator for PodTatoHead installation on Kubernetes

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
   Run `make generate` and `make manifests` commands to run on every change in `api/v1alpha1` in root directory
 - Apply crd : `kubectl apply -f config/crd/bases`
 - Run Operator : `make run`
 - `kubectl apply -f config/manifests/podtatohead.yaml -n oumk`
