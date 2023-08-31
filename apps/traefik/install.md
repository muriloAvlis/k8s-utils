# Traefik Installation on K8s Cluster

## Requiments

- Ready K8s Cluster
- Helm

## Install

```bash
helm repo add traefik https://traefik.github.io/charts
helm repo update
helm install -n traefik-v2 --create-namespace traefik traefik/traefik
```

## Export Traefik dashboard

```bash
kubectl apply -f ingressRoute.yaml
```