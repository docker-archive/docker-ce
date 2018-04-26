# Changelog 
For more information on the list of deprecated flags and APIs please have a look at
https://docs.docker.com/engine/deprecated/ where you can find the target removal dates 

## 18.03.1-ce (2018-04-26)

### Client

- Fix error with merge compose file with networks [docker/cli#983](https://github.com/docker/cli/pull/983)
* Fix docker stack deploy re-deploying services after the service was updated with `--force` [docker/cli#963](https://github.com/docker/cli/pull/963)
* Fix docker version output alignment [docker/cli#965](https://github.com/docker/cli/pull/965)

### Runtime

- Fix AppArmor profiles not being applied to `docker exec` processes [moby/moby#36466](https://github.com/moby/moby/pull/36466)
- Don't sort plugin mount slice [moby/moby#36711](https://github.com/moby/moby/pull/36711)
- Daemon/cluster: handle partial attachment entries during configure [moby/moby#36769](https://github.com/moby/moby/pull/36769)
* Bump Golang to 1.9.5 [moby/moby#36779](https://github.com/moby/moby/pull/36779) [docker/cli#986](https://github.com/docker/cli/pull/986)
- Daemon/stats: more resilient cpu sampling [moby/moby#36519](https://github.com/moby/moby/pull/36519)
* Containerd: update to 1.0.3 release [moby/moby#36749](https://github.com/moby/moby/pull/36749)
- Fix Windows layer leak when write fails [moby/moby#36728](https://github.com/moby/moby/pull/36728)
* Don't make container mount unbindable [moby/moby#36768](https://github.com/moby/moby/pull/36768)
- Fix Daemon panics on container export after a daemon restart [moby/moby/36586](https://github.com/moby/moby/pull/36586)
- Fix digest cache being removed on autherrors [moby/moby#36509](https://github.com/moby/moby/pull/36509)
- Make sure plugin container is removed on failure [moby/moby#36715](https://github.com/moby/moby/pull/36715)
- Copy: avoid using all system memory with authz plugins [moby/moby#36595](https://github.com/moby/moby/pull/36595)
- Relax some libcontainerd client locking [moby/moby#36848](https://github.com/moby/moby/pull/36848)

### Swarm Mode

* Increase raft Election tick to 10 times Heartbeat tick [moby/moby#36672](https://github.com/moby/moby/pull/36672)

### Networking

* Gracefully remove LB endpoints from services [docker/libnetwork#2112](https://github.com/docker/libnetwork/pull/2112)
* Retry other external DNS servers on ServFail [docker/libnetwork#2121](https://github.com/docker/libnetwork/pull/2121)
* Improve scalabiltiy of bridge network isolation rules [docker/libnetwork#2117](https://github.com/docker/libnetwork/pull/2117)
* Allow for larger preset property values, do not override [docker/libnetwork#2124](https://github.com/docker/libnetwork/pull/2124)
* Prevent panics on concurrent reads/writes when calling `changeNodeState` [docker/libnetwork#2136](https://github.com/docker/libnetwork/pull/2136)

## 18.03.0-ce (2018-03-21)

### Builder

* Switch to -buildmode=pie [moby/moby#34369](https://github.com/moby/moby/pull/34369)
* Allow Dockerfile to be outside of build-context [docker/cli#886](https://github.com/docker/cli/pull/886) 
* Builder: fix wrong cache hits building from tars [moby/moby#36329](https://github.com/moby/moby/pull/36329)
- Fixes files leaking to other images in a multi-stage build [moby/moby#36338](https://github.com/moby/moby/pull/36338)

### Client

* Simplify the marshaling of compose types.Config [docker/cli#895](https://github.com/docker/cli/pull/895)
+ Add support for multiple composefile when deploying [docker/cli#569](https://github.com/docker/cli/pull/569)
- Fix broken Kubernetes stack flags [docker/cli#831](https://github.com/docker/cli/pull/831)
- Fix stack marshaling for Kubernetes [docker/cli#890](https://github.com/docker/cli/pull/890)
- Fix and simplify bash completion for service env, mounts and labels [docker/cli#682](https://github.com/docker/cli/pull/682)
- Fix `before` and `since` filter for `docker ps` [moby/moby#35938](https://github.com/moby/moby/pull/35938)
- Fix `--label-file` weird behavior [docker/cli#838](https://github.com/docker/cli/pull/838)
- Fix compilation of defaultCredentialStore() on unsupported platforms [docker/cli#872](https://github.com/docker/cli/pull/872)
* Improve and fix bash completion for images [docker/cli#717](https://github.com/docker/cli/pull/717)
+ Added check for empty source in bind mount [docker/cli#824](https://github.com/docker/cli/pull/824)
- Fix TLS from environment variables in client [moby/moby#36270](https://github.com/moby/moby/pull/36270)
* docker build now runs faster when registry-specific credential helper(s) are configured [docker/cli#840](https://github.com/docker/cli/pull/840)
* Update event filter zsh completion with `disable`, `enable`, `install` and `remove` [docker/cli#372](https://github.com/docker/cli/pull/372)
* Produce errors when empty ids are passed into inspect calls [moby/moby#36144](https://github.com/moby/moby/pull/36144)
* Marshall version for the k8s controller [docker/cli#891](https://github.com/docker/cli/pull/891)
* Set a non-zero timeout for HTTP client communication with plugin backend [docker/cli#883](https://github.com/docker/cli/pull/883)
+ Add DOCKER_TLS environment variable for --tls option [docker/cli#863](https://github.com/docker/cli/pull/863)
+ Add --template-driver option for secrets/configs [docker/cli#896](https://github.com/docker/cli/pull/896)
+ Move `docker trust` commands out of experimental [docker/cli#934](https://github.com/docker/cli/pull/934) [docker/cli#935](https://github.com/docker/cli/pull/935) [docker/cli#944](https://github.com/docker/cli/pull/944)

### Logging

* AWS logs - don't add new lines to maximum sized events [moby/moby#36078](https://github.com/moby/moby/pull/36078)
* Move log validator logic after plugins are loaded [moby/moby#36306](https://github.com/moby/moby/pull/36306)
* Support a proxy in Splunk log driver [moby/moby#36220](https://github.com/moby/moby/pull/36220)
- Fix log tail with empty logs [moby/moby#36305](https://github.com/moby/moby/pull/36305)

### Networking

* Libnetwork revendoring [moby/moby#36137](https://github.com/moby/moby/pull/36137)
- Fix for deadlock on exit with Memberlist revendor [docker/libnetwork#2040](https://github.com/docker/libnetwork/pull/2040)
* Fix user specified ndots option [docker/libnetwork#2065](https://github.com/docker/libnetwork/pull/2065) 
- Fix to use ContainerID for Windows instead of SandboxID [docker/libnetwork#2010](https://github.com/docker/libnetwork/pull/2010)
* Verify NetworkingConfig to make sure EndpointSettings is not nil [moby/moby#36077](https://github.com/moby/moby/pull/36077)
- Fix `DockerNetworkInternalMode` issue [moby/moby#36298](https://github.com/moby/moby/pull/36298)
- Fix race in attachable network attachment [moby/moby#36191](https://github.com/moby/moby/pull/36191)
- Fix timeout issue of `InspectNetwork` on AArch64 [moby/moby#36257](https://github.com/moby/moby/pull/36257)
* Verbose info is missing for partial overlay ID [moby/moby#35989](https://github.com/moby/moby/pull/35989)
* Update `FindNetwork` to address network name duplications [moby/moby#30897](https://github.com/moby/moby/pull/30897)
* Disallow attaching ingress network [docker/swarmkit#2523](https://github.com/docker/swarmkit/pull/2523)
- Prevent implicit removal of the ingress network [moby/moby#36538](https://github.com/moby/moby/pull/36538)
- Fix stale HNS endpoints on Windows [moby/moby#36603](https://github.com/moby/moby/pull/36603)
- IPAM fixes for duplicate IP addresses [docker/libnetwork#2104](https://github.com/docker/libnetwork/pull/2104) [docker/libnetwork#2105](https://github.com/docker/libnetwork/pull/2105)

### Runtime

* Enable HotAdd for Windows [moby/moby#35414](https://github.com/moby/moby/pull/35414)
* LCOW: Graphdriver fix deadlock in hotRemoveVHDs [moby/moby#36114](https://github.com/moby/moby/pull/36114)
* LCOW: Regular mount if only one layer [moby/moby#36052](https://github.com/moby/moby/pull/36052)
* Remove interim env var LCOW_API_PLATFORM_IF_OMITTED [moby/moby#36269](https://github.com/moby/moby/pull/36269)
* Revendor Microsoft/opengcs @ v0.3.6 [moby/moby#36108](https://github.com/moby/moby/pull/36108)
- Fix issue of ExitCode and PID not show up in Task.Status.ContainerStatus [moby/moby#36150](https://github.com/moby/moby/pull/36150)
- Fix issue with plugin scanner going too deep [moby/moby#36119](https://github.com/moby/moby/pull/36119)
* Do not make graphdriver homes private mounts [moby/moby#36047](https://github.com/moby/moby/pull/36047)
* Do not recursive unmount on cleanup of zfs/btrfs [moby/moby#36237](https://github.com/moby/moby/pull/36237)
* Don't restore image if layer does not exist [moby/moby#36304](https://github.com/moby/moby/pull/36304)
* Adjust minimum API version for templated configs/secrets [moby/moby#36366](https://github.com/moby/moby/pull/36366)
* Bump containerd to 1.0.2 (cfd04396dc68220d1cecbe686a6cc3aa5ce3667c) [moby/moby#36308](https://github.com/moby/moby/pull/36308)
* Bump Golang to 1.9.4 [moby/moby#36243](https://github.com/moby/moby/pull/36243)
* Ensure daemon root is unmounted on shutdown [moby/moby#36107](https://github.com/moby/moby/pull/36107)
- Fix container cleanup on daemon restart [moby/moby#36249](https://github.com/moby/moby/pull/36249)
* Support SCTP port mapping (bump up API to v1.37) [moby/moby#33922](https://github.com/moby/moby/pull/33922)
* Support SCTP port mapping [docker/cli#278](https://github.com/docker/cli/pull/278)
- Fix Volumes property definition in ContainerConfig [moby/moby#35946](https://github.com/moby/moby/pull/35946)
* Bump moby and dependencies [docker/cli#829](https://github.com/docker/cli/pull/829)
* C.RWLayer: check for nil before use [moby/moby#36242](https://github.com/moby/moby/pull/36242)
+ Add `REMOVE` and `ORPHANED` to TaskState [moby/moby#36146](https://github.com/moby/moby/pull/36146)
- Fixed error detection using `IsErrNotFound` and `IsErrNotImplemented` for `ContainerStatPath`, `CopyFromContainer`, and `CopyToContainer` methods [moby/moby#35979](https://github.com/moby/moby/pull/35979)
+ Add an integration/internal/container helper package [moby/moby#36266](https://github.com/moby/moby/pull/36266)
+ Add canonical import path [moby/moby#36194](https://github.com/moby/moby/pull/36194)
+ Add/use container.Exec() to integration [moby/moby#36326](https://github.com/moby/moby/pull/36326)
- Fix "--node-generic-resource" singular/plural [moby/moby#36125](https://github.com/moby/moby/pull/36125)
* Daemon.cleanupContainer: nullify container RWLayer upon release [moby/moby#36160](https://github.com/moby/moby/pull/36160)
* Daemon: passdown the `--oom-kill-disable` option to containerd [moby/moby#36201](https://github.com/moby/moby/pull/36201)
* Display a warn message when there is binding ports and net mode is host [moby/moby#35510](https://github.com/moby/moby/pull/35510)
* Refresh containerd remotes on containerd restarted [moby/moby#36173](https://github.com/moby/moby/pull/36173)
* Set daemon root to use shared propagation [moby/moby#36096](https://github.com/moby/moby/pull/36096)
* Optimizations for recursive unmount [moby/moby#34379](https://github.com/moby/moby/pull/34379)
* Perform plugin mounts in the runtime [moby/moby#35829](https://github.com/moby/moby/pull/35829)
* Graphdriver: Fix RefCounter memory leak [moby/moby#36256](https://github.com/moby/moby/pull/36256)
* Use continuity fs package for volume copy [moby/moby#36290](https://github.com/moby/moby/pull/36290)
* Use proc/exe for reexec [moby/moby#36124](https://github.com/moby/moby/pull/36124)
+ Add API support for templated secrets and configs [moby/moby#33702](https://github.com/moby/moby/pull/33702) and [moby/moby#36366](https://github.com/moby/moby/pull/36366)
* Use rslave propagation for mounts from daemon root [moby/moby#36055](https://github.com/moby/moby/pull/36055)
+ Add /proc/keys to masked paths [moby/moby#36368](https://github.com/moby/moby/pull/36368)
* Bump Runc to 1.0.0-rc5 [moby/moby#36449](https://github.com/moby/moby/pull/36449)
- Fixes `runc exec` on big-endian architectures [moby/moby#36449](https://github.com/moby/moby/pull/36449)
* Use chroot when mount namespaces aren't provided [moby/moby#36449](https://github.com/moby/moby/pull/36449)
- Fix systemd slice expansion so that it could be consumed by cAdvisor [moby/moby#36449](https://github.com/moby/moby/pull/36449)
- Fix devices mounted with wrong uid/gid [moby/moby#36449](https://github.com/moby/moby/pull/36449)
- Fix read-only containers with IPC private mounts `/dev/shm` read-only [moby/moby#36526](https://github.com/moby/moby/pull/36526)

### Swarm Mode

* Replace EC Private Key with PKCS#8 PEMs [docker/swarmkit#2246](https://github.com/docker/swarmkit/pull/2246)
* Fix IP overlap with empty EndpointSpec [docker/swarmkit #2505](https://github.com/docker/swarmkit/pull/2505)
* Add support for Support SCTP port mapping [docker/swarmkit#2298](https://github.com/docker/swarmkit/pull/2298)
* Do not reschedule tasks if only placement constraints change and are satisfied by the assigned node [docker/swarmkit#2496](https://github.com/docker/swarmkit/pull/2496)
* Ensure task reaper stopChan is closed no more than once [docker/swarmkit #2491](https://github.com/docker/swarmkit/pull/2491)
* Synchronization fixes [docker/swarmkit#2495](https://github.com/docker/swarmkit/pull/2495)
* Add log message to indicate message send retry if streaming unimplemented [docker/swarmkit#2483](https://github.com/docker/swarmkit/pull/2483)
* Debug logs for session, node events on dispatcher, heartbeats [docker/swarmkit#2486](https://github.com/docker/swarmkit/pull/2486)
+ Add swarm types to bash completion event type filter [docker/cli#888](https://github.com/docker/cli/pull/888)
- Fix issue where network inspect does not show Created time for networks in swarm scope [moby/moby#36095](https://github.com/moby/moby/pull/36095)
