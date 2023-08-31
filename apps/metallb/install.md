# Metallb Installation on K8s Cluster

## Requiments

- Ready K8s Cluster
- Helm

## Install

```bash
helm repo add metallb https://metallb.github.io/metallb
helm repo update
helm install metallb -n metallb-system --create-namespace metallb/metallb
```

## Configuration

First is need define a IP address pool. A template that can be used is in [ipAddressPool](ipAddressPool.yaml).

```bash
kubectl apply -f ipAddressPool.yaml
```

Then advertise service IPs by layer 2 configuration (there are other methods). A template that can be used is in [l2Advertisement](l2Advertisement.yaml).

```bash
kubectl apply -f l2Advertisement.yaml
```