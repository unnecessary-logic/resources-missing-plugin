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

### Install by Release

You can download the latest release and architecture-specific binary from the GitHub releases page, then add the downloaded binary to your `PATH`.

```bash
# Example for macOSX / linux / WSL
export VERSION=0.1.0
export OS=$(uname | tr '[:upper:]' '[:lower:]')
export ARCH=$(uname -m | sed 's/x86_64/amd64/; s/aarch64/arm64/')

wget https://github.com/unnecessary-logic/resources-missing-plugin/releases/download/v${VERSION}/kubectl-resources_missing_${VERSION}_${OS}_${ARCH}.tar.gz

tar -xzf kubectl-resources_missing_${VERSION}_${OS}_${ARCH}.tar.gz kubectl-resources_missing

chmod +x kubectl-resources_missing

sudo mv kubectl-resources_missing /usr/local/bin/
```

### Install by Clone and Build
```
# You can pull the repo down and install it using make build and make install commands like so:

make build      # -> bin/kubectl-resources_missing
make install    # copy into $(go env GOPATH)/bin
kubectl plugin list
kubectl resources-missing --help
```
kubectl discovers any `kubectl-resources_missing` executable on `PATH`; the `_`
becomes `-` in the subcommand name.
