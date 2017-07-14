# Building your own Docker rpm package

`.rpm` packages can be built from this directory with the following syntax

```shell
make ENGINE_DIR=/path/to/engine CLI_DIR=/path/to/cli rpm
```

Artifacts will be located in `rpmbuild` under the following directory structure:
`rpmbuild/$distro-$distro_version/`

### NOTES:
* `ENGINE_DIR` -> Specifies the directory where the engine code is located, eg: `$GOPATH/src/github.com/docker/docker`
* `CLI_DIR` -> Specifies the directory where the cli code is located, eg: `$GOPATH/src/github.com/docker/cli`

## Specifying a specific distro

```shell
make ENGINE_DIR=/path/to/engine CLI_DIR=/path/to/cli fedora
```

## Specifying a specific distro version
```shell
make ENGINE_DIR=/path/to/engine CLI_DIR=/path/to/cli fedora-25
```

## Building the latest docker-ce

```shell
git clone https://github.com/docker/docker-ce.git
make ENGINE_DIR=docker-ce/components/engine CLI_DIR=docker-ce/components/cli rpm
```
