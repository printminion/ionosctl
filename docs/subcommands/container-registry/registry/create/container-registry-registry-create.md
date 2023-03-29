---
description: Create a registry
---

# ContainerRegistryRegistryCreate

## Usage

```text
ionosctl container-registry registry create [flags]
```

## Aliases

For `container-registry` command:

```text
[cr contreg cont-reg]
```

For `registry` command:

```text
[reg registries r]
```

For `create` command:

```text
[c]
```

## Description

Create a registry to hold container images or OCI compliant artifacts

## Options

```text
      --cols strings                               Set of columns to be printed on output 
                                                   Available columns: [RegistryId DisplayName Location Hostname GarbageCollectionDays GarbageCollectionTime]
      --garbage-collection-schedule-days strings   Specify the garbage collection schedule days
      --garbage-collection-schedule-time string    Specify the garbage collection schedule time of day using RFC3339 format
      --location string                            Specify the location of the registry (required)
  -n, --name string                                Specify the name of the registry (required)
```

## Examples

```text
ionosctl container-registry registry create
```
