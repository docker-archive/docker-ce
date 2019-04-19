# Changelog

For official release notes for Docker Engine CE and Docker Engine EE, visit the
[release notes page](https://docs.docker.com/engine/release-notes/).

## 19.03.0 (2019-05-21)

### Client

* Update buildkit to 62e55427. [docker/cli#1800](https://github.com/docker/cli/pull/1800)
* Cli change to pass driver specific options to docker run. [docker/cli#1767](https://github.com/docker/cli/pull/1767)
* build: allow setting buildkit outputs. [docker/cli#1766](https://github.com/docker/cli/pull/1766)
* Add `--pids-limit` flag to `docker update`. [docker/cli#1765](https://github.com/docker/cli/pull/1765)
* Add systctl support for services. [docker/cli#1754](https://github.com/docker/cli/pull/1754)
* Add support for `template_driver` in composefiles. [docker/cli#1746](https://github.com/docker/cli/pull/1746)
* Bump Golang 1.12.4. [docker/cli#1832](https://github.com/docker/cli/pull/1832)
* Fix labels copying value from environment variables. [docker/cli#1671](https://github.com/docker/cli/pull/1671)
* The `docker system info` output now segregates information relevant to the client and daemon. [docker/cli#1638](https://github.com/docker/cli/pull/1638)
* (Experimental) When targetting Kubernetes, add support for `x-pull-secret: some-pull-secret` in compose-files service configs. [docker/cli#1617](https://github.com/docker/cli/pull/1617)
* (Experimental) When targetting Kubernetes, add support for `x-pull-policy: <Never|Always|IfNotPresent>` in compose-files service configs. [docker/cli#1617](https://github.com/docker/cli/pull/1617)
* Add support for maximum replicas per node without stack. [docker/cli#1612](https://github.com/docker/cli/pull/1612)
* Add --device support for Windows. [docker/cli#1606](https://github.com/docker/cli/pull/1606)
* Basic framework for writing and running CLI plugins. [docker/cli#1564](https://github.com/docker/cli/pull/1564)
* Fix tty initial size error. [docker/cli#1529](https://github.com/docker/cli/pull/1529)
* cp, save, export: Prevent overwriting irregular files. [docker/cli#1515](https://github.com/docker/cli/pull/1515)
* Data Path Port configuration support. [docker/cli#1509](https://github.com/docker/cli/pull/1509)
* Fast context switch: commands. [docker/cli#1501](https://github.com/docker/cli/pull/1501)
* Support --mount type=bind,bind-nonrecursive,... [docker/cli#1430](https://github.com/docker/cli/pull/1430)
* Deprecate legacy overlay storage driver. [docker/cli#1425](https://github.com/docker/cli/pull/1425)
* Deprecate "devicemapper" storage driver. [docker/cli#1424](https://github.com/docker/cli/pull/1424)
* build: add SSH agent socket forwarder (`docker build --ssh $SSHMOUNTID=$SSH_AUTH_SOCK`) [docker/cli#1419](https://github.com/docker/cli/pull/1419)
* Add maximum replicas per node support to stack version 3.8. [docker/cli#1410](https://github.com/docker/cli/pull/1410)
* Allow npipe volume type on stack file. [docker/cli#1195](https://github.com/docker/cli/pull/1195)
* Add option to pull images quietly. [docker/cli#882](https://github.com/docker/cli/pull/882)
* Add a separate `--domainname` flag. [docker/cli#1130](https://github.com/docker/cli/pull/1130)
* Add `--from` flag to `context create`. [docker/cli#1773](https://github.com/docker/cli/pull/1773)
* Add support for secret drivers in `docker stack deploy`. [docker/cli#1783](https://github.com/docker/cli/pull/1783)
* Add ability to use swarm `Configs` as `CredentialSpecs` on services. [docker/cli#1781](https://github.com/docker/cli/pull/1781)
* Add `--security-opt systempaths=unconfined` support. [docker/cli#1808](https://github.com/docker/cli/pull/1808)
* Bump Docker App to v0.8.0-beta1. [docker/docker-ce-packaging#324](https://github.com/docker/docker-ce-packaging/pull/324)

### API

* Update API version to v1.40. [moby/moby#38089](https://github.com/moby/moby/pull/38089)
* Add warnings to `/info` endpoint, and move detection to the daemon. [moby/moby#37502](https://github.com/moby/moby/pull/37502)
* Add HEAD support for `/_ping` endpoint. [moby/moby#38570](https://github.com/moby/moby/pull/38570)
* Add `Cache-Control` headers to disable caching `/_ping` endpoint. [moby/moby#38569](https://github.com/moby/moby/pull/38569)
* Add containerd, runc, and docker-init versions to /version. [moby/moby#37974](https://github.com/moby/moby/pull/37974)
* Add undocumented `/grpc` endpoint and register BuildKit's controller. [moby/moby#38990](https://github.com/moby/moby/pull/38990)

### Builder

* Builder: fix `COPY --from` should preserve ownership. [moby/moby#38599](https://github.com/moby/moby/pull/38599)
* builder-next: update buildkit to c3541087 (v0.4.0). [moby/moby#38882](https://github.com/moby/moby/pull/38882)
  * This brings in inline cache support. --cache-from can now point to an existing image
  if it was built with `--build-arg BUILDKIT_INLINE_CACHE=true` and contains the cache metadata in the image config.
* builder-next: allow outputs configuration. [moby/moby#38898](https://github.com/moby/moby/pull/38898)
* TODO changes from BuildKit

### Experimental

* Enable checkpoint/restore of containers with TTY. [moby/moby#38405](https://github.com/moby/moby/pull/38405)
* LCOW: Add support for memory and CPU limits. [moby/moby#37296](https://github.com/moby/moby/pull/37296)
* Windows: Experimental: ContainerD runtime. [moby/moby#38541](https://github.com/moby/moby/pull/38541)

### Security

* mount: add BindOptions.NonRecursive (API v1.40). [moby/moby#38003](https://github.com/moby/moby/pull/38003)
* seccomp: whitelist `io_pgetevents()`. [moby/moby#38895](https://github.com/moby/moby/pull/38895)
* seccomp: allow `ptrace(2)` for 4.8+ kernels. [moby/moby#38137](https://github.com/moby/moby/pull/38137)

### Runtime

* Allow running dockerd as a non-root user (Rootless mode). [moby/moby#380050](https://github.com/moby/moby/pull/38050)
* Add DeviceRequests to HostConfig to support NVIDIA GPUs. [moby/moby#38828](https://github.com/moby/moby/pull/38828)
* Making it possible to pass Windows credential specs directly to the engine. [moby/moby#38777](https://github.com/moby/moby/pull/38777)
* Add pids-limit support in docker update. [moby/moby#32519](https://github.com/moby/moby/pull/32519)
* Add support for exact list of capabilities. [moby/moby#38380](https://github.com/moby/moby/pull/38380)
* daemon: use 'private' ipc mode by default. [moby/moby#35621](https://github.com/moby/moby/pull/35621)
* daemon: switch to semaphore-gated WaitGroup for startup tasks. [moby/moby#38301](https://github.com/moby/moby/pull/38301)
* Add --device support for Windows. [moby/moby#37638](https://github.com/moby/moby/pull/37638)
* Add memory.kernelTCP support for linux. [moby/moby#37043](https://github.com/moby/moby/pull/37043)
* Use idtools.LookupGroup instead of parsing /etc/group file for docker.sock ownership to fix: api.go doesn't respect nsswitch.conf. [moby/moby#38126](https://github.com/moby/moby/pull/38126)
* Fix docker --init with /dev bind mount. [moby/moby#37665](https://github.com/moby/moby/pull/37665)
* cli: fix images filter when use multi reference filter. [moby/moby#38171](https://github.com/moby/moby/pull/38171)
* Bump Golang to 1.12.4. [moby/moby#39063](https://github.com/moby/moby/pull/39063)
* Bump containerd to 1.2.6 and runc to 029124d. [moby/moby#39016](https://github.com/moby/moby/pull/39016)

### Networking

* Network: add support for 'dangling' filter. [moby/moby#31551](https://github.com/moby/moby/pull/31551)
* Move IPVLAN driver out of experimental. [moby/moby#38983](https://github.com/moby/moby/pull/38983) / [docker/libnetwork#2230](https://github.com/docker/libnetwork/pull/2230)

### Swarm

* Added support for maximum replicas per node. [moby/moby#37940](https://github.com/moby/moby/pull/37940)
* Add support for GMSA CredentialSpecs from Swarmkit configs. [moby/moby#38632](https://github.com/moby/moby/pull/38632)
* Add support for sysctl options in services. [moby/moby#37701](https://github.com/moby/moby/pull/37701)
* Add support for filtering on node labels. [moby/moby#37650](https://github.com/moby/moby/pull/37650)
* Windows: Support named pipe mounts in docker service create + stack yml. [moby/moby#37400](https://github.com/moby/moby/pull/37400)
* VXLAN UDP Port configuration support. [moby/moby#38102](https://github.com/moby/moby/pull/38102)

### Logging

* Enable gcplogs driver on windows. [moby/moby#37717](https://github.com/moby/moby/pull/37717)
* Add zero padding for RFC5424 syslog format. [moby/moby#38335](https://github.com/moby/moby/pull/38335)
* Add IMAGE_NAME attribute to journald log events. [moby/moby#38032](https://github.com/moby/moby/pull/38032)

### Deprecation

* Remove v1 manifest support, remove `--disable-legacy-registry`. [moby/moby#37874](https://github.com/moby/moby/pull/37874)
* Remove v1.10 migrator. [moby/moby#38265](https://github.com/moby/moby/pull/38265)
* Skip deprecated storage-drivers in auto-selection. [moby/moby#38019](https://github.com/moby/moby/pull/38019)
* Deprecate AuFS storage driver, and add warning. [moby/moby#38090](https://github.com/moby/moby/pull/38090)
