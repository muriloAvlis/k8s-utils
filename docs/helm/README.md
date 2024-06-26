# Helm

## Installation

```sh
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh
rm -rf get_helm.sh
```

## Enable completion

```sh
helm completion bash | sudo tee /etc/bash_completion.d/helm > /dev/null
exec bash
```