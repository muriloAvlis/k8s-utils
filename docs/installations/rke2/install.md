# K8s Installation Using the RKE2

To install a K8s cluster using RKE2, run the following commands:

## Configuration

```bash
mkdir -p /etc/rancher/rke2 && \
cat <<EOF | sudo tee /etc/rancher/rke2/config.yaml
node-name: oranlocal-master
disable: rke2-ingress-nginx
cni:
    - multus
    - calico
EOF
```

## Install

```bash
curl -sfL https://get.rke2.io | INSTALL_RKE2_VERSION="v1.24.16+rke2r1" sh -
```

## Start K8s cluster

```bash
systemctl enable rke2-server.service 
systemctl start rke2-server.service
journalctl -u rke2-server -f
```

## Get Cluster configuration

```bash
mkdir -p ~/.kube/
cp /etc/rancher/rke2/rke2.yaml ~/.kube/config
```