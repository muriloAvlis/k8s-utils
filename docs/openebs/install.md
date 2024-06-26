# OpenEBS install using Helm

## Requirements

- Helm


## Get Started

To install the OpenEBS using helm chart, use the following command:

```bash
helm repo add openebs https://openebs.github.io/charts
helm repo update
helm install openebs --namespace openebs openebs/openebs --create-namespace
```