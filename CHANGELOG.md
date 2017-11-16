# Changelog

Items starting with `DEPRECATE` are important deprecation notices. For more
information on the list of deprecated flags and APIs please have a look at
https://docs.docker.com/engine/deprecated/ where target removal dates can also
be found.

## 17.11.0-ce (2017-11-DD)

IMPORTANT: Docker CE 17.11 is the first Docker release based on
[containerd 1.0 beta](https://github.com/containerd/containerd/releases/tag/v1.0.0-beta.2).
Docker CE 17.11 and later won't recognize containers started with
previous Docker versions. If using
[Live Restore](https://docs.docker.com/engine/admin/live-restore/#enable-the-live-restore-option),
you must stop all containers before upgrading to Docker CE 17.11.
If you don't, any containers started by Docker versions that predate
17.11 won't be recognized by Docker after the upgrade and will keep
running, un-managed, on the system.

### Builder

* Test & Fix build with rm/force-rm matrix [moby/moby#35139](https://github.com/moby/moby/pull/35139)
- Fix build with `--stream` with a large context [moby/moby#35404](https://github.com/moby/moby/pull/35404)

### Client

* Hide help flag from help output [docker/cli#645](https://github.com/docker/cli/pull/645)
* Support parsing of named pipes for compose volumes [docker/cli#560](https://github.com/docker/cli/pull/560)
* [Compose] Cast values to expected type after interpolating values [docker/cli#601](https://github.com/docker/cli/pull/601)
+ Add output for "secrets" and "configs" on `docker stack deploy` [docker/cli#593](https://github.com/docker/cli/pull/593)
- Fix flag description for `--host-add` [docker/cli#648](https://github.com/docker/cli/pull/648)
* Do not truncate ID on docker service ps --quiet [docker/cli#579](https://github.com/docker/cli/pull/579)

### Deprecation

* Update bash completion and deprecation for synchronous service updates [docker/cli#610](https://github.com/docker/cli/pull/610)

### Logging

* copy to log driver's bufsize, fixes #34887 [moby/moby#34888](https://github.com/moby/moby/pull/34888)
+ Add TCP support for GELF log driver [moby/moby#34758](https://github.com/moby/moby/pull/34758)
+ Add credentials endpoint option for awslogs driver [moby/moby#35055](https://github.com/moby/moby/pull/35055)

### Networking

- Fix network name masking network ID on delete [moby/moby#34509](https://github.com/moby/moby/pull/34509)
- Fix returned error code for network creation from 500 to 409 [moby/moby#35030](https://github.com/moby/moby/pull/35030)
- Fix tasks fail with error "Unable to complete atomic operation, key modified" [docker/libnetwork#2004](https://github.com/docker/libnetwork/pull/2004)

### Runtime

* Switch to Containerd 1.0 client [moby/moby#34895](https://github.com/moby/moby/pull/34895)
* Increase container default shutdown timeout on Windows [moby/moby#35184](https://github.com/moby/moby/pull/35184)
* LCOW: API: Add `platform` to /images/create and /build [moby/moby#34642](https://github.com/moby/moby/pull/34642)
* Stop filtering Windows manifest lists by version [moby/moby#35117](https://github.com/moby/moby/pull/35117)
* Use windows console mode constants from Azure/go-ansiterm [moby/moby#35056](https://github.com/moby/moby/pull/35056)
* Windows Daemon should respect DOCKER_TMPDIR [moby/moby#35077](https://github.com/moby/moby/pull/35077)
* Windows: Fix startup logging [moby/moby#35253](https://github.com/moby/moby/pull/35253)
+ Add support for Windows version filtering on pull [moby/moby#35090](https://github.com/moby/moby/pull/35090)
- Fixes LCOW after containerd 1.0 introduced regressions [moby/moby#35320](https://github.com/moby/moby/pull/35320)
* ContainerWait on remove: don't stuck on rm fail [moby/moby#34999](https://github.com/moby/moby/pull/34999)
* oci: obey CL_UNPRIVILEGED for user namespaced daemon [moby/moby#35205](https://github.com/moby/moby/pull/35205)
* Don't abort when setting may_detach_mounts [moby/moby#35172](https://github.com/moby/moby/pull/35172)
- Fix panic on get container pid when live restore containers [moby/moby#35157](https://github.com/moby/moby/pull/35157)
- Mask `/proc/scsi` path for containers to prevent removal of devices (CVE-2017-16539) [moby/moby#35399](https://github.com/moby/moby/pull/35399)
* Update to github.com/vbatts/tar-split@v0.10.2 (CVE-2017-14992) [moby/moby#35424](https://github.com/moby/moby/pull/35424)
- Container: protect health monitor channel [moby/moby#35482](https://github.com/moby/moby/pull/35482)
- Libcontainerd: fix leaking container/exec state [moby/moby#35484](https://github.com/moby/moby/pull/35484)

### Swarm Mode

* Modifying integration test due to new ipam options in swarmkit [moby/moby#35103](https://github.com/moby/moby/pull/35103)
- Fix deadlock on getting swarm info [moby/moby#35388](https://github.com/moby/moby/pull/35388)
+ Expand the scope of the `Err` field in `TaskStatus` to also cover non-terminal errors that block the task from progressing [docker/swarmkit#2287](https://github.com/docker/swarmkit/pull/2287)

### Packaging

+ Build packages for Debian 10 (Buster) [docker/docker-ce-packaging#50](https://github.com/docker/docker-ce-packaging/pull/50)
+ Build packages for Ubuntu 17.10 (Artful) [docker/docker-ce-packaging#55](https://github.com/docker/docker-ce-packaging/pull/55)
