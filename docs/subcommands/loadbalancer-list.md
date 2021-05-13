---
description: List Load Balancers
---

# LoadbalancerList

## Usage

```text
ionosctl loadbalancer list [flags]
```

## Description

Use this command to retrieve a list of Load Balancers within a Virtual Data Center on your account.

Required values to run command:

* Data Center Id

## Options

```text
  -u, --api-url string         Override default API endpoint (default "https://api.ionos.com/cloudapi/v5")
      --cols strings           Columns to be printed in the standard output (default [LoadBalancerId,Name,Dhcp])
  -c, --config string          Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
      --datacenter-id string   The unique Data Center Id (required)
      --force                  Force command to execute without user input
  -h, --help                   help for list
  -o, --output string          Desired output format [text|json] (default "text")
  -q, --quiet                  Quiet output
```

## Examples

```text
ionosctl loadbalancer list --datacenter-id f28c0edd-d5ef-48f2-b8a3-aa8f6b55da3d 
LoadbalancerId                         Name               Dhcp
f16dfcc1-9181-400b-a08d-7fe15ca0e9af   demoLoadbalancer   true
3f9f14a9-5fa8-4786-ba86-a91f9daded2c   demoLoadBalancer   false
```
