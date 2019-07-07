# Local Cluster Creation

My local cloud environment with vagrant boxes. Including small utilities for storing
and retrieving cluster tokens to and from vault storage.

## Usage

By default, `run-vms.sh` will up the vagrant boxes by storing tokens in local file system, if you want to store them in `vault`, run the script as follows:

```
>  VAULT_ENDPOINT = your_vault_endpoint
>  VAULT_PREFEX = prefix_where_to_store_token
>  bash run-vms.sh vault
```
