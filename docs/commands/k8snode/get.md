---
description: Get a Kubernetes Node
---

# Get

## Usage

```text
ionosctl k8s-node get [flags]
```

## Description

Use this command to retrieve details about a specific Kubernetes Node.

Required values to run command:

* K8s Cluster Id
* K8s NodePool Id
* K8s Node Id

## Options

```text
  -u, --api-url string       Override default API endpoint (default "https://api.ionos.com/cloudapi/v5")
      --cluster-id string    The unique K8s Cluster Id [Required flag]
      --cols strings         Columns to be printed in the standard output (default [NodeId,Name,K8sVersion,PublicIP,State])
  -c, --config string        Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -h, --help                 help for get
      --ignore-stdin         Force command to execute without user input
      --node-id string       The unique K8s Node Id [Required flag]
      --nodepool-id string   The unique K8s Node Pool Id [Required flag]
  -o, --output string        Desired output format [text|json] (default "text")
  -q, --quiet                Quiet output
```

## Examples

```text
ionosctl k8s-node get --cluster-id ba5e2960-4068-4aee-b972-092c254769a8 --nodepool-id 939811fe-cc13-41e2-8a49-87db58c7a812 --node-id a0e5d4c4-6b09-4965-8e98-59a749301d20 
NodeId                                 Name                   K8sVersion   PublicIP        State
a0e5d4c4-6b09-4965-8e98-59a749301d20   test12345-n3q55ggmap   1.19.8       x.x.x.x         UNKNOWN
```
