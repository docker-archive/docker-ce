# Changelog

For more information on the list of deprecated flags and APIs, have a look at
https://docs.docker.com/engine/deprecated/ where you can find target removal dates.

## 18.02.0-ce (2018-02-07)

### Builder

- Gitutils: fix checking out submodules [moby/moby#35737](https://github.com/moby/moby/pull/35737)

### Client

* Attach: Ensure attach exit code matches container's [docker/cli#696](https://github.com/docker/cli/pull/696)
+ Added support for tmpfs-mode in compose file [docker/cli#808](https://github.com/docker/cli/pull/808)
+ Adds a new compose file version 3.6 [docker/cli#808](https://github.com/docker/cli/pull/808)
- Fix issue of filter in `docker ps` where `health=starting` returns nothing [moby/moby#35940](https://github.com/moby/moby/pull/35940)
+ Improve presentation of published port ranges [docker/cli#581](https://github.com/docker/cli/pull/581)
* Bump Go to 1.9.3 [docker/cli#827](https://github.com/docker/cli/pull/827)
- Fix broken Kubernetes stack flags [docker/cli#831](https://github.com/docker/cli/pull/831)
* Annotate "stack" commands to be "swarm" and "kubernetes" [docker/cli#804](https://github.com/docker/cli/pull/804)

### Experimental

+ Add manifest command [docker/cli#138](https://github.com/docker/cli/pull/138)
* LCOW remotefs - return error in Read() implementation [moby/moby#36051](https://github.com/moby/moby/pull/36051)
+ LCOW: Coalesce daemon stores, allow dual LCOW and WCOW mode [moby/moby#34859](https://github.com/moby/moby/pull/34859)
- LCOW: Fix OpenFile parameters [moby/moby#36043](https://github.com/moby/moby/pull/36043)
* LCOW: Raise minimum requirement to Windows RS3 RTM build (16299) [moby/moby#36065](https://github.com/moby/moby/pull/36065)

### Logging

* Improve daemon config reload; log active configuration [moby/moby#36019](https://github.com/moby/moby/pull/36019)
- Fixed error detection using IsErrNotFound and IsErrNotImplemented for the ContainerLogs method [moby/moby#36000](https://github.com/moby/moby/pull/36000)
+ Add journald tag as SYSLOG_IDENTIFIER [moby/moby#35570](https://github.com/moby/moby/pull/35570)
* Splunk: limit the reader size on error responses [moby/moby#35509](https://github.com/moby/moby/pull/35509)


### Networking

* Disable service on release network results in zero-downtime deployments with rolling upgrades [moby/moby#35960](https://github.com/moby/moby/pull/35960)
- Fix services failing to start if multiple networks with the same name exist in different spaces [moby/moby#30897](https://github.com/moby/moby/pull/30897)
- Fix duplicate networks being added with `docker service update --network-add` [docker/cli#780](https://github.com/docker/cli/pull/780)
- Fixing ingress network when upgrading from 17.09 to 17.12. [moby/moby#36003](https://github.com/moby/moby/pull/36003)
- Fix ndots configuration [docker/libnetwork#1995](https://github.com/docker/libnetwork/pull/1995)
- Fix IPV6 networking being deconfigured if live-restore is enabled [docker/libnetwork#2043](https://github.com/docker/libnetwork/pull/2043)
+ Add support for MX type DNS queries in the embedded DNS server [docker/libnetwork#2041](https://github.com/docker/libnetwork/pull/2041)

### Packaging

+ Added packaging for Fedora 26, Fedora 27, and Centos 7 on aarch64 [docker/docker-ce-packaging#71](https://github.com/docker/docker-ce-packaging/pull/71)
- Removed support for Ubuntu Zesty [docker/docker-ce-packaging#73](https://github.com/docker/docker-ce-packaging/pull/73)
- Removed support for Fedora 25 [docker/docker-ce-packaging#72](https://github.com/docker/docker-ce-packaging/pull/72)


### Runtime

- Fixes unexpected Docker Daemon shutdown based on pipe error [moby/moby#35968](https://github.com/moby/moby/pull/35968)
- Fix some occurrences of hcsshim::ImportLayer failed in Win32: The system cannot find the path specified [moby/moby#35924](https://github.com/moby/moby/pull/35924)
* Windows: increase the maximum layer size during build to 127GB [moby/moby#35925](https://github.com/moby/moby/pull/35925)
- Fix Devicemapper: Error running DeleteDevice dm_task_run failed [moby/moby#35919](https://github.com/moby/moby/pull/35919)
+ Introduce « exec_die » event [moby/moby#35744](https://github.com/moby/moby/pull/35744)
* Update API to version 1.36 [moby/moby#35744](https://github.com/moby/moby/pull/35744)
- Fix `docker update` not updating cpu quota, and cpu-period of a running container [moby/moby#36030](https://github.com/moby/moby/pull/36030)
* Make container shm parent unbindable [moby/moby#35830](https://github.com/moby/moby/pull/35830)
+ Make image (layer) downloads faster by using pigz [moby/moby#35697](https://github.com/moby/moby/pull/35697)
+ Protect the daemon from volume plugins that are slow or deadlocked [moby/moby#35441](https://github.com/moby/moby/pull/35441)
- Fix `DOCKER_RAMDISK` environment variable not being honoured [moby/moby#35957](https://github.com/moby/moby/pull/35957)
* Bump containerd to 1.0.1 (9b55aab90508bd389d7654c4baf173a981477d55) [moby/moby#35986](https://github.com/moby/moby/pull/35986)
* Update runc to fix hang during start and exec [moby/moby#36097](https://github.com/moby/moby/pull/36097) 
- Fix "--node-generic-resource" singular/plural [moby/moby#36125](https://github.com/moby/moby/pull/36125)
