---
description: List Servers
---

# ServerList

## Usage

```text
ionosctl server list [flags]
```

## Description

Use this command to list Servers from a specified Virtual Data Center.

Required values to run command:

* Data Center Id

## Options

```text
  -u, --api-url string         Override default API endpoint (default "https://api.ionos.com/cloudapi/v5")
      --cols strings           Columns to be printed in the standard output (default [ServerId,Name,AvailabilityZone,State,Cores,Ram,CpuFamily])
  -c, --config string          Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
      --datacenter-id string   The unique Data Center Id (required)
      --force                  Force command to execute without user input
  -h, --help                   help for list
  -o, --output string          Desired output format [text|json] (default "text")
  -q, --quiet                  Quiet output
```

## Examples

```text
ionosctl server list --datacenter-id f28c0edd-d5ef-48f2-b8a3-aa8f6b55da3d 
ServerId                               Name         AvailabilityZone   State       Cores   Ram     CpuFamily
f45f435e-8d6c-4170-ab90-858b59dab9ff   demoServer   AUTO               AVAILABLE   4       256MB   AMD_OPTERON
```
