# Building your own Docker deb package

`.deb` packages can be built from this directory with the following syntax

```shell
make deb
```

Artifacts will be located in `debbuild` under the following directory structure:
`debbuild/$distro-$distro_version/`

### Building from local source

Specify the location of the source repositories for the engine and cli when
building packages

* `ENGINE_DIR` -> Specifies the directory where the engine code is located, eg: `$GOPATH/src/github.com/docker/docker`
* `CLI_DIR` -> Specifies the directory where the cli code is located, eg: `$GOPATH/src/github.com/docker/cli`

```shell
make ENGINE_DIR=/path/to/engine CLI_DIR=/path/to/cli deb
```

## Specifying a specific distro

```shell
make ubuntu
```

## Specifying a specific distro version
```shell
make ubuntu-focal
```

## Building the for all distros

```shell
make deb
```
