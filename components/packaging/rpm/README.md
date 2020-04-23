# Building your own Docker rpm package

`.rpm` packages can be built from this directory with the following syntax

```shell
make rpm
```

Artifacts will be located in `rpmbuild` under the following directory structure:
`rpmbuild/$distro-$distro_version/`

### Building from local source

Specify the location of the source repositories for the engine and cli when
building packages

* `ENGINE_DIR` -> Specifies the directory where the engine code is located, eg: `$GOPATH/src/github.com/docker/docker`
* `CLI_DIR` -> Specifies the directory where the cli code is located, eg: `$GOPATH/src/github.com/docker/cli`

```shell
make ENGINE_DIR=/path/to/engine CLI_DIR=/path/to/cli rpm
```

## Specifying a specific distro

```shell
make centos
```

## Specifying a specific distro version
```shell
make centos-8
```

## Building the for all distros

```shell
make rpm
```
