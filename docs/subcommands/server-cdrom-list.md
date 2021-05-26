---
description: List attached CD-ROMs from a Server
---

# ServerCdromList

## Usage

```text
ionosctl server cdrom list [flags]
```

## Aliases

For `server` command:
```text
[s svr]
```

For `cdrom` command:
```text
[cd]
```

For `list` command:
```text
[l ls]
```

## Description

Use this command to retrieve a list of CD-ROMs attached to the Server.

Required values to run command:

* Data Center Id
* Server Id

## Options

```text
  -u, --api-url string         Override default API endpoint (default "https://api.ionos.com/cloudapi/v6")
      --cols strings           Set of columns to be printed on output 
                               Available columns: [ImageId Name ImageAliases Location Size LicenceType ImageType Description Public CloudInit] (default [ImageId,Name,ImageAliases,Location,LicenceType,ImageType,CloudInit])
  -c, --config string          Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
      --datacenter-id string   The unique Data Center Id (required)
  -f, --force                  Force command to execute without user input
  -h, --help                   help for list
  -o, --output string          Desired output format [text|json] (default "text")
  -q, --quiet                  Quiet output
      --server-id string       The unique Server Id (required)
```

## Examples

```text
ionosctl server cdrom list --datacenter-id 4fd7996d-2b08-4c04-9c47-d9d884ee179a --server-id f7438b0c-2f36-4bec-892f-af027930b81e 
ImageId                                Name                              ImageAliases                       Location   LicenceType   ImageType   CloudInit
80c63662-49a0-11ea-94e0-525400f64d8d   CentOS-8.1.1911-x86_64-boot.iso   [centos:8_iso centos:latest_iso]   us/las     LINUX         CDROM       NONE
```
