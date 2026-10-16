# resources-missing-plugin
A Kubectl plugin for finding objects that have no resources block defined, eg no memory / cpu requests and limits.

The plugin will return containers AND initContainers that are missing resource definitions.  It returns whether requests or limits are missing, so additional detective work will be needed in order to discover the exact context of the missing resources.

## Status
Still need more formal testing implemented, but some of the main functions are done and the functionality is working.

## Usage

After installation, you can use the plugin like so:

```bash
# All namespaces
kubectl resources-missing

# Filter by namespace
kubectl resources-missing -n kube-system

kubectl resources-missing --help
```

## Installation
```
# You can pull the repo down and install it using make build and make install commands like so:

make build      # -> bin/kubectl-resources_missing
make install    # copy into $(go env GOPATH)/bin
kubectl plugin list
kubectl resources-missing --help
```
kubectl discovers any `kubectl-resources_missing` executable on `PATH`; the `_`
becomes `-` in the subcommand name.
