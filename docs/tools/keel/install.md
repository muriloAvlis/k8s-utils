# Keel installation using Helm

## Requirements

- K8s
- Helm


## Get Started

To install the Keel using helm chart, use the following command:

```bash
helm repo add keel https://charts.keel.sh 
helm repo update
helm upgrade --install keel --namespace=kube-system keel/keel
```
