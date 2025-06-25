# K8s Installation Using the Kubespray

To install a K8s cluster using RKE2, run the following commands:

## Requirements

- Ubuntu Hosts/VMs (memory >2GB)
- ansible-core (>2.14 and <2.17);
- Jinja (>2.11);
- python-netaddr;

## Scenario

- Manager VM -> ansible commands here;
- Master VM -> K8s master
- Worker-0 VM -> K8s Worker
- Worker-1 VM -> K8s Worker

## Commands on the Manager Host

### Install Prerequisites

```sh
sudo apt install python3-pip

pip install --user ansible==9.13.0 Jinja2 netaddr argcomplete --break-system-packages

## Add ansible bin to PATH
echo 'export PATH=$PATH:~/.local/bin' >> ~/.bashrc

## Enable ansible completion
activate-global-python-argcomplete --user
```

### Copy SSH Public Keys to K8s-nodes

```sh
## Generate key (if you don't have)
ssh-keygen -t ed25519 -C "kubespray-manager"

## Copy the public key
ssh-copy-id -i ~/.ssh/id_ed25519.pub user@master_ip
ssh-copy-id -i ~/.ssh/id_ed25519.pub user@worker_ip
```

### Clone Kubespray

```sh
git clone -b release-2.28 https://github.com/kubernetes-sigs/kubespray.git
```

### Configure Cluster

```sh
cd kubespray
cp -rfp inventory/sample inventory/mycluster

## Edit inventory.ini
nano inventory/mycluster/inventory.ini

nano inventory/mycluster/group_vars/all.yml # for every node, including etcd
nano inventory/mycluster/group_vars/k8s_cluster.yml # for every node in the cluster (not etcd when it's separate)
```

### Test K8s-nodes Connections

```sh
ansible -i inventory/mycluster/inventory.ini all -m ping
```

### Installing Cluster

```sh
ansible-playbook -i inventory/mycluster/inventory.ini cluster.yml -b -v --private-key ~/.ssh/id_ed25519
```

### Adding nodes

TODO

## Commands on the K8s Master

### Test Cluster API
```sh
sudo su
kubectl cluster-info
```