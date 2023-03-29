---
description: Get Certificate by ID
---

# CertificateManagerGet

## Usage

```text
ionosctl certificate-manager get [flags]
```

## Aliases

For `get` command:

```text
[g]
```

## Description

Use this command to retrieve a Certificate by ID.

## Options

```text
      --certificate             Print the certificate
      --certificate-chain       Print the certificate chain
  -i, --certificate-id string   Response get a single certificate (required)
      --cols strings            Set of columns to be printed on output 
                                Available columns: [CertId DisplayName]
```

## Examples

```text
ionosctl certificate-manager get --certificate-id 47c5d9cc-b613-4b76-b0cc-dc531787a422
```
