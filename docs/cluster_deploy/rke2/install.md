# K8s Installation Using the RKE2

To install a K8s cluster using RKE2, run the following commands:

## Install Server Node

```bash
curl -sfL https://get.rke2.io | sudo sh -
```

> **_NOTE_** Use INSTALL_RKE2_VERSION="vX.XX.XX+rke2r1" to set K8s version

### Server Configuration

```bash
sudo mkdir -p /etc/rancher/rke2 && \
sudo cat <<EOF | sudo tee /etc/rancher/rke2/config.yaml
cni:
    - multus
    - calico
EOF
```

### Start Server Node

```bash
sudo systemctl enable rke2-server.service 
sudo systemctl start rke2-server.service
```

### Get Cluster configuration

```bash
mkdir -p ~/.kube/
sudo cp /etc/rancher/rke2/rke2.yaml ~/.kube/config
sudo chown $USER:$USER ~/.kube/config
```

### Get kubectl

```sh
sudo cp /var/lib/rancher/rke2/bin/kubectl /usr/local/bin/kubectl
kubectl completion bash | sudo tee /etc/bash_completion.d/kubectl > /dev/null
```

## Install Worker Node (optional)

On another machine, run the following commands:

```sh
curl -sfL https://get.rke2.io | sudo INSTALL_RKE2_TYPE="agent" sh -
```


### Worker Configuration

```sh
sudo mkdir -p /etc/rancher/rke2 && \
sudo cat <<EOF | sudo tee /etc/rancher/rke2/config.yaml
server: https://<server-addr>:9345
token: <token from server node>
EOF
```

> **_NOTE_** The token is available on /var/lib/rancher/rke2/server/node-token

### Start Worker Node

```sh
sudo systemctl enable rke2-agent.service
sudo systemctl start rke2-agent.service
```

## Clean up

On each machine run the following command:

```bash
sudo sh /usr/local/bin/rke2-uninstall.sh
```