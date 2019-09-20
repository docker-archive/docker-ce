# Changelog

For official release notes for Docker Engine CE and Docker Engine EE, visit the
[release notes page](https://docs.docker.com/engine/release-notes/).

## 19.03.3 (2019-10-02)

### Builder

- Fix builder-next: resolve digest for third party registries. [docker/engine#339](https://github.com/docker/engine/pull/339)
- Fix builder-next: user namespace builds when daemon started with socket activation. [docker/engine#339](https://github.com/docker/engine/pull/373)
- Fix builder-next: session: release forwarded ssh socket connection per connection. [docker/engine#339](https://github.com/docker/engine/pull/373)
- Fix builder-next: llbsolver: error on multiple cache importers. [docker/engine#339](https://github.com/docker/engine/pull/373)

### Networking

- Fix various libnetwork issues for iptables, DNS queries, and more. [docker/engine#330](https://github.com/docker/engine/pull/330)

### Runtime

* Bump Golang to 1.12.9. [docker/engine#366](https://github.com/docker/engine/pull/366)
* Bump containerd to 1.2.9. [docker/engine#349](https://github.com/docker/engine/pull/349)
* Distribution: modify warning logic when pulling v2 schema1 manifests. [docker/engine#368](https://github.com/docker/engine/pull/368)
- Fix `POST /images/create` returning a 500 status code when providing an incorrect platform option. [docker/engine#365](https://github.com/docker/engine/pull/365)
- Fix `POST /build` returning a 500 status code when providing an incorrect platform option. [docker/engine#365](https://github.com/docker/engine/pull/365)
- Fix atomic: patch 64bit alignment on 32bit systems. [docker/engine#363](https://github.com/docker/engine/pull/363)
- Fix to ensure parent dir exists for mount cleanup file. [docker/engine#360](https://github.com/docker/engine/pull/360)
- Fix to return "invalid parameter" when linking to non-existing container. [docker/engine#352](https://github.com/docker/engine/pull/352)
- Fix overlay2 busy error on mount. [docker/engine#332](https://github.com/docker/engine/pull/332)
- Fix `docker rmi -f` hang. [docker/engine#335](https://github.com/docker/engine/pull/335)
- Fix handling of blocked I/O of exec'd processes. [docker/engine#296](https://github.com/docker/engine/pull/296)

### Swarm

- Fix for removing hard-coded IPAM config subnet value for ingress network. [docker/engine#369](https://github.com/docker/engine/pull/369)
- Fix update out of sequence and increase max recv gRPC message size for nodes and secrets. [docker/engine#352](https://github.com/docker/engine/pull/353)

## 19.03.2 (2019-08-29)

### Builder

- Fix "COPY --from" to non-existing directory on Windows. [moby/moby#39695](https://github.com/moby/moby/pull/39695)
- Fix builder-next: metadata commands not having created time in history. [moby/moby#39456](https://github.com/moby/moby/issues/39456)
- Fix builder-next: close progress on layer export error. [moby/moby#39782](https://github.com/moby/moby/pull/39782)
* Update buildkit to 588c73e1e4. [moby/moby#39781](https://github.com/moby/moby/pull/39781)

### Client

- Fix Windows absolute path detection on non-Windows. [docker/cli#1990](https://github.com/docker/cli/pull/1990)
- Fix to zsh completion script for `docker login --username`.
- Fix context: produce consistent output on `context create`. [docker/cli#1985](https://github.com/docker/cli/pull/1874)
- Fix support for HTTP proxy env variable. [docker/cli#2059](https://github.com/docker/cli/pull/2059)

### Logging

- Fix for reading journald logs. [moby/moby#37819](https://github.com/moby/moby/pull/37819) [moby/moby#38859](https://github.com/moby/moby/pull/38859)

### Networking

- Prevent panic on network attach to a container with disabled networking. [moby/moby#39589](https://github.com/moby/moby/pull/39589)

### Runtime

* Bump Golang to 1.12.8.
- Fix a potential engine panic when using XFS disk quota for containers. [moby/moby#39644](https://github.com/moby/moby/pull/39644)

### Swarm

- Fix an issue where nodes with lots of tasks could not be removed. [docker/swarmkit#2867](https://github.com/docker/swarmkit/pull/2867)

## 19.03.1 (2019-07-25)

### Runtime

- Fix [CVE-2019-14271](https://nvd.nist.gov/vuln/detail/CVE-2019-14271) loading of nsswitch based config inside chroot under Glibc.

## 19.03.0 (2019-07-22)

### Deprecation

* Deprecate image manifest v2 schema1 in favor of v2 schema2. Future version of Docker will remove support for v2 schema1 altogether. [moby/moby#39365](https://github.com/moby/moby/pull/39365)
* Remove v1.10 migrator. [moby/moby#38265](https://github.com/moby/moby/pull/38265)
* Skip deprecated storage-drivers in auto-selection. [moby/moby#38019](https://github.com/moby/moby/pull/38019)
* Deprecate `aufs` storage driver and add warning. [moby/moby#38090](https://github.com/moby/moby/pull/38090)

### Client

+ Add `--pids-limit` flag to `docker update`. [docker/cli#1765](https://github.com/docker/cli/pull/1765)
+ Add systctl support for services. [docker/cli#1754](https://github.com/docker/cli/pull/1754)
+ Add support for `template_driver` in composefiles. [docker/cli#1746](https://github.com/docker/cli/pull/1746)
+ Add --device support for Windows. [docker/cli#1606](https://github.com/docker/cli/pull/1606)
+ Data Path Port configuration support. [docker/cli#1509](https://github.com/docker/cli/pull/1509)
+ Fast context switch: commands. [docker/cli#1501](https://github.com/docker/cli/pull/1501)
+ Support --mount type=bind,bind-nonrecursive,... [docker/cli#1430](https://github.com/docker/cli/pull/1430)
+ Add maximum replicas per node. [docker/cli#1410](https://github.com/docker/cli/pull/1410) [docker/cli#1612](https://github.com/docker/cli/pull/1612)
+ Add option to pull images quietly. [docker/cli#882](https://github.com/docker/cli/pull/882)
+ Add a separate `--domainname` flag. [docker/cli#1130](https://github.com/docker/cli/pull/1130)
+ Add support for secret drivers in `docker stack deploy`. [docker/cli#1783](https://github.com/docker/cli/pull/1783)
+ Add ability to use swarm `Configs` as `CredentialSpecs` on services. [docker/cli#1781](https://github.com/docker/cli/pull/1781)
+ Add `--security-opt systempaths=unconfined` support. [docker/cli#1808](https://github.com/docker/cli/pull/1808)
+ Basic framework for writing and running CLI plugins. [docker/cli#1564](https://github.com/docker/cli/pull/1564) [docker/cli#1898](https://github.com/docker/cli/pull/1898)
+ Docker App v0.8.0. [docker/docker-ce-packaging#341](https://github.com/docker/docker-ce-packaging/pull/341)
+ Docker buildx. [docker/docker-ce-packaging#336](https://github.com/docker/docker-ce-packaging/pull/336)
* Bump google.golang.org/grpc to v1.20.1. [docker/cli#1884](https://github.com/docker/cli/pull/1884)
* Cli change to pass driver specific options to docker run. [docker/cli#1767](https://github.com/docker/cli/pull/1767)
* Bump Golang 1.12.5. [docker/cli#1875](https://github.com/docker/cli/pull/1875)
* The `docker system info` output now segregates information relevant to the client and daemon. [docker/cli#1638](https://github.com/docker/cli/pull/1638)
* (Experimental) When targetting Kubernetes, add support for `x-pull-secret: some-pull-secret` in compose-files service configs. [docker/cli#1617](https://github.com/docker/cli/pull/1617)
* (Experimental) When targetting Kubernetes, add support for `x-pull-policy: <Never|Always|IfNotPresent>` in compose-files service configs. [docker/cli#1617](https://github.com/docker/cli/pull/1617)
* cp, save, export: Prevent overwriting irregular files. [docker/cli#1515](https://github.com/docker/cli/pull/1515)
* Allow npipe volume type on stack file. [docker/cli#1195](https://github.com/docker/cli/pull/1195)
- Fix tty initial size error. [docker/cli#1529](https://github.com/docker/cli/pull/1529)
- Fix labels copying value from environment variables. [docker/cli#1671](https://github.com/docker/cli/pull/1671)

### API

+ Update API version to v1.40. [moby/moby#38089](https://github.com/moby/moby/pull/38089)
+ Add warnings to `/info` endpoint, and move detection to the daemon. [moby/moby#37502](https://github.com/moby/moby/pull/37502)
+ Add HEAD support for `/_ping` endpoint. [moby/moby#38570](https://github.com/moby/moby/pull/38570)
+ Add `Cache-Control` headers to disable caching `/_ping` endpoint. [moby/moby#38569](https://github.com/moby/moby/pull/38569)
+ Add containerd, runc, and docker-init versions to /version. [moby/moby#37974](https://github.com/moby/moby/pull/37974)
* Add undocumented `/grpc` endpoint and register BuildKit's controller. [moby/moby#38990](https://github.com/moby/moby/pull/38990)

### Builder

+ builder-next: allow setting buildkit outputs. [docker/cli#1766](https://github.com/docker/cli/pull/1766)
+ builder-next: look for a Dockerfile specific dockerignore file (eg. Dockerfile.dockerignore) for ignored paths. [docker/engine#215](https://github.com/docker/engine/pull/215)
+ builder-next: automatically detect if process execution is possible for x86, arm and arm64 binaries. [docker/engine#215](https://github.com/docker/engine/pull/215)
+ builder-next: added inline cache support `--cache-from`. [docker/engine#215](https://github.com/docker/engine/pull/215)
+ builder-next: allow outputs configuration. [moby/moby#38898](https://github.com/moby/moby/pull/38898)
* builder-next: update buildkit to 1f89ec1. [docker/engine#260](https://github.com/docker/engine/pull/260)
* builder-next: buildkit now also uses systemd's resolv.conf. [docker/engine#260](https://github.com/docker/engine/pull/260)
* builder-next: use Dockerfile frontend version `docker/dockerfile:1.1` by default. [docker/engine#215](https://github.com/docker/engine/pull/215)
* builder-next: no longer rely on an external image for COPY/ADD operations. [docker/engine#215](https://github.com/docker/engine/pull/215)
- Builder: fix `COPY --from` should preserve ownership. [moby/moby#38599](https://github.com/moby/moby/pull/38599)
- builder-next: fix gcr workaround token cache. [docker/engine#212](https://github.com/docker/engine/pull/212)
- builder-next: call stopprogress on download error. [docker/engine#215](https://github.com/docker/engine/pull/215)

### Experimental

+ Enable checkpoint/restore of containers with TTY. [moby/moby#38405](https://github.com/moby/moby/pull/38405)
+ LCOW: Add support for memory and CPU limits. [moby/moby#37296](https://github.com/moby/moby/pull/37296)
* Windows: Experimental: ContainerD runtime. [moby/moby#38541](https://github.com/moby/moby/pull/38541)
* Windows: Experimental: LCOW requires Windows RS5+. [moby/moby#39108](https://github.com/moby/moby/pull/39108)

### Security

* mount: add BindOptions.NonRecursive (API v1.40). [moby/moby#38003](https://github.com/moby/moby/pull/38003)
* seccomp: whitelist `io_pgetevents()`. [moby/moby#38895](https://github.com/moby/moby/pull/38895)
* seccomp: allow `ptrace(2)` for 4.8+ kernels. [moby/moby#38137](https://github.com/moby/moby/pull/38137)

### Runtime

+ Allow running dockerd as a non-root user (Rootless mode). [moby/moby#380050](https://github.com/moby/moby/pull/38050)
+ Rootless: optional support for `lxc-user-nic` SUID binary. [docker/engine#208](https://github.com/docker/engine/pull/208)
+ Add DeviceRequests to HostConfig to support NVIDIA GPUs. [moby/moby#38828](https://github.com/moby/moby/pull/38828)
+ Add `--device` support for Windows. [moby/moby#37638](https://github.com/moby/moby/pull/37638)
+ Add memory.kernelTCP support for linux. [moby/moby#37043](https://github.com/moby/moby/pull/37043)
* Making it possible to pass Windows credential specs directly to the engine. [moby/moby#38777](https://github.com/moby/moby/pull/38777)
* Add pids-limit support in docker update. [moby/moby#32519](https://github.com/moby/moby/pull/32519)
* Add support for exact list of capabilities. [moby/moby#38380](https://github.com/moby/moby/pull/38380)
* daemon: use 'private' ipc mode by default. [moby/moby#35621](https://github.com/moby/moby/pull/35621)
* daemon: switch to semaphore-gated WaitGroup for startup tasks. [moby/moby#38301](https://github.com/moby/moby/pull/38301)
* Use idtools.LookupGroup instead of parsing /etc/group file for docker.sock ownership to fix: api.go doesn't respect nsswitch.conf. [moby/moby#38126](https://github.com/moby/moby/pull/38126)
* cli: fix images filter when use multi reference filter. [moby/moby#38171](https://github.com/moby/moby/pull/38171)
* Bump Golang to 1.12.5. [docker/engine#209](https://github.com/docker/engine/pull/209)
* Bump containerd to 1.2.6. [moby/moby#39016](https://github.com/moby/moby/pull/39016)
* Bump runc to 1.0.0-rc8, opencontainers/selinux v1.2.2. [docker/engine#210](https://github.com/docker/engine/pull/210)
* Bump google.golang.org/grpc to v1.20.1. [docker/engine#215](https://github.com/docker/engine/pull/215)
* Performance optimizations in aufs and layer store for massively parallel container creation/removal. [moby/moby#39135](https://github.com/moby/moby/pull/39135) [moby/moby#39209](https://github.com/moby/moby/pull/39209)
* Pass root to chroot to for chroot Tar/Untar (CVE-2018-15664) [moby/moby#39292](https://github.com/moby/moby/pull/39292)
- Fix docker `--init` with `/dev` bind mount. [moby/moby#37665](https://github.com/moby/moby/pull/37665)
- Fix: fetch the right device number when greater than 255 and using `--device-read-bps` option. [moby/moby#39212](https://github.com/moby/moby/pull/39212)
- Fix: "Path does not exist" error when path definitely exists. [moby/moby#39251](https://github.com/moby/moby/pull/39251)
- Fix: [CVE-2018-15664](https://nvd.nist.gov/vuln/detail/CVE-2018-15664) symlink-exchange attack with directory traversal. [moby/moby#39357](https://github.com/moby/moby/pull/39357)
- Fix [CVE-2019-13509](https://nvd.nist.gov/vuln/detail/CVE-2019-13509) in DebugRequestMiddleware: unconditionally scrub data field.

### Networking

+ Move IPVLAN driver out of experimental. [moby/moby#38983](https://github.com/moby/moby/pull/38983) / [docker/libnetwork#2230](https://github.com/docker/libnetwork/pull/2230)
* Network: add support for 'dangling' filter. [moby/moby#31551](https://github.com/moby/moby/pull/31551)
* Windows: Forcing a nil IP specified in PortBindings to IPv4zero (0.0.0.0). [docker/libnetwork#2376](https://github.com/docker/libnetwork/pull/2376)
- Fix to make sure load balancer sandbox is deleted when a service is updated with `--network-rm`. [docker/engine#213](https://github.com/docker/engine/pull/213)

### Swarm

+ Add support for maximum replicas per node. [moby/moby#37940](https://github.com/moby/moby/pull/37940)
+ Add support for GMSA CredentialSpecs from Swarmkit configs. [moby/moby#38632](https://github.com/moby/moby/pull/38632)
+ Add support for sysctl options in services. [moby/moby#37701](https://github.com/moby/moby/pull/37701)
+ Add support for filtering on node labels. [moby/moby#37650](https://github.com/moby/moby/pull/37650)
+ Windows: Support named pipe mounts in docker service create + stack yml. [moby/moby#37400](https://github.com/moby/moby/pull/37400)
+ VXLAN UDP Port configuration support. [moby/moby#38102](https://github.com/moby/moby/pull/38102)
* Use Service Placement Constraints in Enforcer. [docker/swarmkit#2857](https://github.com/docker/swarmkit/pull/2857)
* Increase max recv gRPC message size for nodes and secrets. [docker/engine#256](https://github.com/docker/engine/pull/256)

### Logging

* Enable gcplogs driver on windows. [moby/moby#37717](https://github.com/moby/moby/pull/37717)
* Add zero padding for RFC5424 syslog format. [moby/moby#38335](https://github.com/moby/moby/pull/38335)
* Add IMAGE_NAME attribute to journald log events. [moby/moby#38032](https://github.com/moby/moby/pull/38032)
