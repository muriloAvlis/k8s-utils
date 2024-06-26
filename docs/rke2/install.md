# K8s Installation Using the RKE2

To install a K8s cluster using RKE2, run the following commands:

### Configuration

```bash
sudo mkdir -p /etc/rancher/rke2 && \
sudo cat <<EOF | sudo tee /etc/rancher/rke2/config.yaml
node-name: k8s-master
cni:
    - multus
    - calico
EOF
```

### Install

```bash
curl -sfL https://get.rke2.io | sudo sh -
```

> **NOTE**: Use INSTALL_RKE2_VERSION="vX.XX.XX+rke2r1" to set K8s version

### Start K8s cluster

```bash
sudo systemctl enable rke2-server.service 
sudo systemctl start rke2-server.service
```

### Get Cluster configuration

```bash
mkdir -p ~/.kube/
sudo cp /etc/rancher/rke2/rke2.yaml ~/.kube/config
```

### Get kubectl

```sh
sudo cp /var/lib/rancher/rke2/bin/kubectl /usr/local/bin/kubectl
kubectl completion bash | sudo tee /etc/bash_completion.d/kubectl > /dev/null
```

## Clean up

```bash
sudo sh /usr/local/bin/rke2-uninstall.sh
```