---
description: List IpConsumers
---

# IpconsumerList

## Usage

```text
ionosctl ipconsumer list [flags]
```

## Aliases

For `ipconsumer` command:
```text
[ipc]
```

For `list` command:
```text
[l ls]
```

## Description

Use this command to get a list of Resources where each IP address from an IpBlock is being used.

Required values to run command:

* IpBlock Id

## Options

```text
  -u, --api-url string      Override default API endpoint (default "https://api.ionos.com/cloudapi/v6")
      --cols strings        Set of columns to be printed on output 
                            Available columns: [Ip Mac NicId ServerId ServerName DatacenterId DatacenterName K8sNodePoolId K8sClusterId] (default [Ip,NicId,ServerId,DatacenterId,K8sNodePoolId,K8sClusterId])
  -c, --config string       Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -f, --force               Force command to execute without user input
  -h, --help                help for list
      --ipblock-id string   The unique IpBlock Id (required)
  -o, --output string       Desired output format [text|json] (default "text")
  -q, --quiet               Quiet output
```

## Examples

```text
ionosctl ipconsumer list --ipblock-id 564f4984-8349-40c1-bcd8-ba177ebf2fb6
```
