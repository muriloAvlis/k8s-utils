# K8s Installation Using the Kubeadm

To install a K8s cluster using Kubeadm, run the following commands:

## Requirement

- Ubuntu 22.04 or newer;
- [Docker/Containerd](/scripts/dockerInstall.sh);
- [kubeadm](/scripts/kubeadmInstall.sh);
- helm (for CNI installation)

## VMs Preparation

- Disable swap

```sh
sudo swapoff -a
sudo sed -i -e '/swap.img/ s/^#*/#/g' /etc/fstab
```

- Enable IPv4 packet forwarding and br_netfiler module

```sh
sudo modprobe br_netfilter
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF

# sysctl params required by setup, params persist across reboots
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.ipv4.ip_forward = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF

# Apply sysctl params without reboot
sudo sysctl --system
```

- Configure Containerd

```sh
## Enable CRI Plugins
sudo sed -i -e '/disabled_plugins/ s/^#*/#/g' /etc/containerd/config.toml

## Set Systemd Cgroup driver
sudo sed -i '14i version = 2' /etc/containerd/config.toml

cat <<EOF | sudo tee -a /etc/containerd/config.toml > /dev/null

[plugins."io.containerd.grpc.v1.cri"]
  sandbox_image = "registry.k8s.io/pause:3.10"
  [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
    runtime_type = "io.containerd.runc.v2"
    [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
      SystemdCgroup = true
EOF

sudo systemctl restart containerd.service
```

## Setup Cluster (Master)

- Create Cluster 

```sh
## Download k8s images
kubeadm config images pull

## Start cluster
sudo kubeadm init --pod-network-cidr 172.168.0.0/16
```

- Get Cluster Config

```sh
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

- Untaint Master Node

```sh
kubectl taint nodes --all node-role.kubernetes.io/control-plane-
kubectl taint nodes --all node-role.kubernetes.io/master-
```

### Install Container Network Interface (CNI)

- Calico

```sh
helm repo add projectcalico https://docs.tigera.io/calico/charts
kubectl create namespace tigera-operator
helm install calico projectcalico/tigera-operator --version v3.28.1 --namespace tigera-operator
```

- Flannel (TODO)

- Mutus (TODO)

### Join Others Nodes

- On master

```sh
kubeadm token create --print-join-command
```

- On worker

```sh
## Use script to install kubeadm, kubelet and kubectl
chmod +x kubeadmInstall
./kubeadmInstall.sh

## Run join command
kubeadm join <API_SERVER_IP>:<PORT> --token <TOKEN> --discovery-token-ca-cert-hash sha256:<HASH>
```