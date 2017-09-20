# Changelog

Items starting with `DEPRECATE` are important deprecation notices. For more
information on the list of deprecated flags and APIs please have a look at
https://docs.docker.com/engine/deprecated/ where target removal dates can also
be found.

## 17.09.0-ce (2017-09-dd)

### Builder

+ Add `--chown` flag to `ADD/COPY` commands in Dockerfile [moby/moby#34263](https://github.com/moby/moby/pull/34263)
* Fix cloning unneeded files while building from git repositories [moby/moby#33704](https://github.com/moby/moby/pull/33704)

### Client

* Allow extension fields in the v3.4 version of the compose format [docker/cli#452](https://github.com/docker/cli/pull/452)
* Make compose file allow to specify names for non-external volume [docker/cli#306](https://github.com/docker/cli/pull/306)
* Support `--compose-file -` as stdin [docker/cli#347](https://github.com/docker/cli/pull/347)
* Support `start_period` for healthcheck in Docker Compose [docker/cli#475](https://github.com/docker/cli/pull/475)
+ Add support for `stop-signal` in docker stack commands [docker/cli#388](https://github.com/docker/cli/pull/388)
+ Add support for update order in compose deployments [docker/cli#360](https://github.com/docker/cli/pull/360)
+ Add ulimits to unsupported compose fields [docker/cli#482](https://github.com/docker/cli/pull/482)
+ Add `--format` to `docker-search` [docker/cli#440](https://github.com/docker/cli/pull/440)
* Show images digests when `{{.Digest}}` is in format [docker/cli#439](https://github.com/docker/cli/pull/439)
* Print timestamp when `--human=true` [docker/cli#438](https://github.com/docker/cli/pull/438)


### Networking

+ Add the control plane MTU option in the daemon config [moby/moby#34103](https://github.com/moby/moby/pull/34103)
+ Add service virtual IP to sandbox's loopback address [docker/libnetwork#1877](https://github.com/docker/libnetwork/pull/1877)

### Runtime

* Graphdriver: promote overlay2 over aufs [moby/moby#34430](https://github.com/moby/moby/pull/34430)
* LCOW: Additional flags for VHD boot [moby/moby#34451](https://github.com/moby/moby/pull/34451)
* LCOW: Don't block export [moby/moby#34448](https://github.com/moby/moby/pull/34448)
* LCOW: Dynamic sandbox management [moby/moby#34170](https://github.com/moby/moby/pull/34170)
* LCOW: Force Hyper-V Isolation [moby/moby#34468](https://github.com/moby/moby/pull/34468)
* LCOW: Move toolsScratchPath to /tmp [moby/moby#34396](https://github.com/moby/moby/pull/34396)
* LCOW: Remove hard-coding [moby/moby#34398](https://github.com/moby/moby/pull/34398)
* LCOW: WORKDIR correct handling [moby/moby#34405](https://github.com/moby/moby/pull/34405)
* Windows: named pipe mounts [moby/moby#33852](https://github.com/moby/moby/pull/33852)

### Swarm Mode

* Include whether the managers in the swarm are autolocked as part of `docker info` [docker/cli#471](https://github.com/docker/cli/pull/471)
+ Add 'docker service rollback' subcommand [docker/cli#205](https://github.com/docker/cli/pull/205)

### Packaging

+ Build scripts for ppc64el on Ubuntu [docker/docker-ce-packaging#43](https://github.com/docker/docker-ce-packaging/pull/43)
