---
description: List Kubernetes NodePool LANs
---

# K8sNodepoolLanList

## Usage

```text
ionosctl k8s nodepool lan list [flags]
```

## Aliases

For `nodepool` command:
```text
[np]
```

For `list` command:
```text
[l ls]
```

## Description

Use this command to get a list of all contained NodePool LANs in a selected Kubernetes Cluster.

Required values to run command:

* K8s Cluster Id
* K8s NodePool Id

## Options

```text
  -u, --api-url string       Override default API endpoint (default "https://api.ionos.com/cloudapi/v6")
      --cluster-id string    The unique K8s Cluster Id (required)
      --cols strings         Set of columns to be printed on output 
                             Available columns: [LanId Dhcp RoutesNetwork RoutesGatewayIp] (default [LanId,Dhcp,RoutesNetwork,RoutesGatewayIp])
  -c, --config string        Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -f, --force                Force command to execute without user input
  -h, --help                 help for list
      --nodepool-id string   The unique K8s Node Pool Id (required)
  -o, --output string        Desired output format [text|json] (default "text")
  -q, --quiet                Quiet output
```

## Examples

```text
ionosctl k8s nodepool lan list --cluster-id CLUSTER_ID --nodepool-id NODEPOOL_ID
```
