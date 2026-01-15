# SR-IOV Device Plugin Installation

## Requirements

- Kubernetes;
- Kubectl;
- SR-IOV CNI;

## Getting Started

### Install SR-IOV CNI

Run the following command to install the SR-IOV CNI:

```sh
kubectl apply -f yamls/sriov-cni-daemonset.yaml
```

### Install SR-IOV Device Plugin

First configure the SR-IOV Device Plugin by editing the [sriovdp-config.yaml](/yamls/sriovdp-config.yaml) file to match your environment and requirements. 

Then run the following command to install the SR-IOV Device Plugin:

```sh
kubectl apply -f yamls/sriovdp-daemonset.yaml
```