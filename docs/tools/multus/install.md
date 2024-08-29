# Multus Installation

## Requirements

- Git
- K8s Cluster
- kubectl

## Get Starting

### Install Multus from GitHub Repository

```sh
kubectl apply -f https://raw.githubusercontent.com/k8snetworkplumbingwg/multus-cni/v4.1.0/deployments/multus-daemonset-thick.yml
```