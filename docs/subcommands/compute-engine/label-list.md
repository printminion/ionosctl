---
description: List Labels from Resources
---

# LabelList

## Usage

```text
ionosctl label list [flags]
```

## Aliases

For `list` command:

```text
[l ls]
```

## Description

Use this command to list all Labels from all Resources under your account. If you want to list all Labels from a specific Resource, use `--resource-type` option together with the Resource Id: `--datacenter-id`, `--server-id`, `--volume-id`.

## Options

```text
  -u, --api-url string         Override default host url (default "https://api.ionos.com")
      --cols strings           Set of columns to be printed on output 
                               Available columns: [Key Value] (default [Key,Value])
  -c, --config string          Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
      --datacenter-id string   The unique Data Center Id
  -f, --force                  Force command to execute without user input
  -h, --help                   Print usage
      --ipblock-id string      The unique IpBlock Id
      --no-headers             When using text output, don't print headers
  -o, --output string          Desired output format [text|json] (default "text")
  -q, --quiet                  Quiet output
      --resource-type string   Type of the resource to list labels from
      --server-id string       The unique Server Id
      --snapshot-id string     The unique Snapshot Id
  -v, --verbose                Print step-by-step process when running command
      --volume-id string       The unique Volume Id
```

## Examples

```text
ionosctl label list

ionosctl label list --resource-type datacenter --datacenter-id DATACENTER_ID
```
