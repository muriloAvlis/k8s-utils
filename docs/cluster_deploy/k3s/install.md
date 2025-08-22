# K8s Installation Using the K3s

To install a K8s cluster using K3s, run the following commands:

### Server Configuration

```bash
sudo mkdir -p /etc/rancher/k3s && \
sudo cat <<EOF | sudo tee /etc/rancher/k3s/config.yaml
write-kubeconfig-mode: "0644"
node-label:
  - "testbed=openran-brasil"
  - "project=nori-blueprint-v2"
cluster-init: true
EOF
```

## Install Server Node

```bash
curl -sfL https://get.k3s.io | sh - 
```

> **NOTE**: Use INSTALL_K3S_VERSION="vX.XX.X+k3sX" to set K8s version

<!-- ### Get Cluster configuration

```bash
mkdir -p ~/.kube/
sudo cp /etc/rancher/k3s/k3s.yaml ~/.kube/config
sudo chown $USER:$USER ~/.kube/config
``` -->

### Get Cluster Configuration

```sh
mkdir -p ~/.kube/
sudo cp /etc/rancher/k3s/k3s.yaml ~/.kube/config
sudo chown $USER:$USER ~/.kube/config
```

### Get kubectl

```sh
kubectl completion bash | sudo tee /etc/bash_completion.d/kubectl > /dev/null
```

## Clean up

```bash
sudo sh /usr/local/bin/k3s-uninstall.sh
```
