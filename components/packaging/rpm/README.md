# Building your own Docker rpm package

`.rpm` packages can be built from this directory with the following syntax

```shell
make ENGINE_DIR=/path/to/engine CLI_DIR=/path/to/cli rpm
```

If you want to specify a specific distro:

```shell
make ENGINE_DIR=/path/to/engine CLI_DIR=/path/to/cli fedora
```

If you want to specify a specific distro version:

```shell
make ENGINE_DIR=/path/to/engine CLI_DIR=/path/to/cli fedora-25
```

