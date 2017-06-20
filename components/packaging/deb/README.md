# Building your own Docker deb package

`.deb` packages can be built from this directory with the following syntax

```shell
make ENGINE_DIR=/path/to/engine CLI_DIR=/path/to/cli deb
```

If you want to specify a specific distro:

```shell
make ENGINE_DIR=/path/to/engine CLI_DIR=/path/to/cli ubuntu
```

If you want to specify a specific distro version:

```shell
make ENGINE_DIR=/path/to/engine CLI_DIR=/path/to/cli ubuntu-xenial
```

