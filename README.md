## Constellation CLI tools

### Usage

```
Usage: cl-tools [options...] COMMAND [options...]
Constellation command line tools
Options:
  -h, --help            Display usage
  -v, --version         Display version
      --vv              Display version (extended)
Commands:
  check-alignment       Checks alignment on the cluster
  check-balance         Checks address balance on the cluster
    -a, --address       Wallet address
```

### Check cluster alignment

```bash
$ cl-tools check-alignment

Checking alignment on nodes: 20
Height: 244      Hashes: 1       Unique hashes: 1
Height: 246      Hashes: 1       Unique hashes: 1
Height: 248      Hashes: 1       Unique hashes: 1
Height: 250      Hashes: 1       Unique hashes: 1
Height: 252      Hashes: 1       Unique hashes: 1
Height: 254      Hashes: 1       Unique hashes: 1
Height: 256      Hashes: 3       Unique hashes: 1
Height: 258      Hashes: 7       Unique hashes: 1
Height: 260      Hashes: 20      Unique hashes: 1
Height: 262      Hashes: 20      Unique hashes: 1
Height: 264      Hashes: 20      Unique hashes: 1
Height: 266      Hashes: 20      Unique hashes: 1
...
Height: 410      Hashes: 20      Unique hashes: 1
Height: 412      Hashes: 20      Unique hashes: 1
Height: 414      Hashes: 19      Unique hashes: 1
Height: 416      Hashes: 17      Unique hashes: 1
Height: 418      Hashes: 13      Unique hashes: 1
Cluster is aligned!
```

### Check address balance

```bash
$ cl-tools check-balance -a DAG7GLsoEM7yoFsEqyUmiLXnE95haac4AKxcQpZh

Checking balance for address: DAG7GLsoEM7yoFsEqyUmiLXnE95haac4AKxcQpZh on nodes: 20
Same balance across nodes: 99900000000
```

