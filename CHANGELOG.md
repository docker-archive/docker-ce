# Changelog

Items starting with `DEPRECATE` are important deprecation notices. For more
information on the list of deprecated flags and APIs please have a look at
https://docs.docker.com/engine/deprecated/ where target removal dates can also
be found.

## 17.10.0-ce (2017-10-DD)

IMPORTANT: Starting with this release, `docker service create`, `docker service update`,
`docker service scale` and `docker service rollback` use non-detached mode as default,
use `--detach` to keep the old behaviour.

### Builder

* Reset uid/gid to 0 in uploaded build context to share build cache with other clients [docker/cli#513](https://github.com/docker/cli/pull/513)
+ Add support for `ADD` urls without any sub path [moby/moby#34217](https://github.com/moby/moby/pull/34217)

### Client

* Move output of `docker stack rm` to stdout [docker/cli#491](https://github.com/docker/cli/pull/491)
* Use natural sort secrets and configs in cli [docker/cli#307](https://github.com/docker/cli/pull/307)
* Use non-detached mode as default for `docker service` commands [docker/cli#525](https://github.com/docker/cli/pull/525)
* Set APIVersion on the client, even when Ping fails [docker/cli#546](https://github.com/docker/cli/pull/546)
- Fix loader error with different build syntax in `docker stack deploy` [docker/cli#544](https://github.com/docker/cli/pull/544)
* Change the default output format for `docker container stats` to show `CONTAINER ID` and `NAME` [docker/cli#565](https://github.com/docker/cli/pull/565)
+ Add `--no-trunc` flag to `docker container stats` [docker/cli#565](https://github.com/docker/cli/pull/565)
+ Add experimental `docker trust`: `view`, `revoke`, `sign` subcommands [docker/cli#472](https://github.com/docker/cli/pull/472)

### Networking

* Enabling ILB/ELB on windows using per-node, per-network LB endpoint [moby/moby#34674](https://github.com/moby/moby/pull/34674)

### Runtime

* LCOW: Add UVM debugability by grabbing logs before tear-down [moby/moby#34846](https://github.com/moby/moby/pull/34846)
* LCOW: Prepare work for bind mounts [moby/moby#34258](https://github.com/moby/moby/pull/34258)
* LCOW: Support for docker cp, ADD/COPY on build [moby/moby#34252](https://github.com/moby/moby/pull/34252)
* LCOW: VHDX boot to readonly [moby/moby#34754](https://github.com/moby/moby/pull/34754)
* Volume: evaluate symlinks before relabeling mount source [moby/moby#34792](https://github.com/moby/moby/pull/34792)
- Fixing ‘docker cp’ to allow new target file name in a host symlinked directory [moby/moby#31993](https://github.com/moby/moby/pull/31993)

### Swarm Mode

* Produce an error if `docker swarm init --force-new-cluster` is executed on worker nodes [moby/moby#34881](https://github.com/moby/moby/pull/34881)
+ Add support for `.Node.Hostname` templating in swarm services [moby/moby#34686](https://github.com/moby/moby/pull/34686)
