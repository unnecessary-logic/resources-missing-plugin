# resources-missing-plugin
A Kubectl plugin for finding objects that have no resources block defined, eg no memory / cpu requests and limits.

## Status
Skeleton only. `main.go` sets up the cobra command and the standard kubectl
config flags; the actual scan is a `TODO`.

## Develop
```
make build      # -> bin/kubectl-resources_missing
make install    # copy into $(go env GOPATH)/bin
kubectl plugin list
kubectl resources-missing --help
```
kubectl discovers any `kubectl-resources_missing` executable on `PATH`; the `_`
becomes `-` in the subcommand name.
