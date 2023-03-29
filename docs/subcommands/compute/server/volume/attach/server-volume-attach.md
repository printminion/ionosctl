---
description: Attach a Volume to a Server
---

# ServerVolumeAttach

## Usage

```text
ionosctl server volume attach [flags]
```

## Aliases

For `server` command:

```text
[s svr]
```

For `volume` command:

```text
[v vol]
```

For `attach` command:

```text
[a]
```

## Description

Use this command to attach a pre-existing Volume to a Server.

You can wait for the Request to be executed using `--wait-for-request` option.

Required values to run command:

* Data Center Id
* Server Id
* Volume Id

## Options

```text
      --cols strings           Set of columns to be printed on output 
                               Available columns: [VolumeId Name Size Type LicenceType State Image Bus AvailabilityZone BackupunitId DeviceNumber UserData BootServerId DatacenterId] (default [VolumeId,Name,Size,Type,LicenceType,State,Image])
      --datacenter-id string   The unique Data Center Id (required)
      --server-id string       The unique Server Id (required)
  -t, --timeout int            Timeout option for Request for Volume attachment [seconds] (default 60)
  -i, --volume-id string       The unique Volume Id (required)
  -w, --wait-for-request       Wait for the Request for Volume attachment to be executed
```

## Examples

```text
ionosctl server volume attach --datacenter-id DATACENTER_ID --server-id SERVER_ID --volume-id VOLUME_ID
```
