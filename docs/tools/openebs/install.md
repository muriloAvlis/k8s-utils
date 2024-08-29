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

To enable Jiva or cStor support, run the following commands:

- To Jiva support

```sh
helm upgrade --install openebs openebs/openebs --namespace openebs --set jiva.enabled=true
```

- To cStor support

```sh
helm install openebs --namespace openebs openebs/openebs --create-namespace --set cstor.enabled=true
```