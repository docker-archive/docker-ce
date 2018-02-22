# Changelog 
Items starting with `DEPRECATE` are important deprecation notices. For more
information on the list of deprecated flags and APIs please have a look at
https://docs.docker.com/engine/deprecated/ where target removal dates can also
be found.

## 18.03.0-ce (2018-03-DD)


### Builder

* switch to -buildmode=pie [moby/moby#34369](https://github.com/moby/moby/pull/34369)
* Allow Dockerfile from outside build-context [docker/cli#886](https://github.com/docker/cli/pull/886)
* Builder: fix wrong cache hits building from tars [moby/moby#36329](https://github.com/moby/moby/pull/36329)
* Update containerd/continuity to fix ARM 32-bit builds [moby/moby#36339](https://github.com/moby/moby/pull/36339)
* Updated docker-on-docker build-notes. [moby/moby#35892](https://github.com/moby/moby/pull/35892)
- Fix string type for buildargs API definition [moby/moby#36246](https://github.com/moby/moby/pull/36246)

### Client

* Simplify the marshaling of compose types.Config [docker/cli#895](https://github.com/docker/cli/pull/895)
* Update reference docs for docker stack deploy [docker/cli#871](https://github.com/docker/cli/pull/871)
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
* Explain the columns shown in docker stats [docker/cli#860](https://github.com/docker/cli/pull/860)
* GetAll -> Get to retrieve credentials from credential helpers [docker/cli#840](https://github.com/docker/cli/pull/840)
+ Add Engine version to docker node ls [docker/cli#885](https://github.com/docker/cli/pull/885)
* Update event filter zsh completion with `disable`, `enable`, `install` and `remove` [docker/cli#372](https://github.com/docker/cli/pull/372)
* Produce errors when empty ids are passed into inspect calls. [moby/moby#36144](https://github.com/moby/moby/pull/36144)
* Marshall version [docker/cli#891](https://github.com/docker/cli/pull/891)
* Replace go-bindata with esc [docker/cli#874](https://github.com/docker/cli/pull/874)
* Set a non-zero timeout for HTTP client communication with plugin backend. [docker/cli#883](https://github.com/docker/cli/pull/883)
+ Add DOCKER_TLS environment variable for --tls option [docker/cli#863](https://github.com/docker/cli/pull/863)

### Logging

* Awslogs - don't add new lines to maximum sized events [moby/moby#36078](https://github.com/moby/moby/pull/36078)
* Move log validator logic after plugins are loaded [moby/moby#36306](https://github.com/moby/moby/pull/36306)
* Support a proxy in splunk log driver [moby/moby#36220](https://github.com/moby/moby/pull/36220)
- Fix log tail with empty logs [moby/moby#36305](https://github.com/moby/moby/pull/36305)

### Networking

* Bump libnetwork to 5ab4ab830062fe8a30a44b75b0bda6b1f4f166a4 [moby/moby#36099](https://github.com/moby/moby/pull/36099)
* Document long form of --network and --network-add [docker/cli#843](https://github.com/docker/cli/pull/843)
* Libnetwork revendoring [moby/moby#36137](https://github.com/moby/moby/pull/36137)
* Migrates TestContainersAPINetworkMountsNoChown to api tests [moby/moby#36198](https://github.com/moby/moby/pull/36198)
* Verify NetworkingConfig to make sure EndpointSettings is not nil [moby/moby#36077](https://github.com/moby/moby/pull/36077)
+ Add description to TestContainerNetworkMountsNoChown [moby/moby#36226](https://github.com/moby/moby/pull/36226)
- Fix `DockerNetworkInternalMode` issue [moby/moby#36298](https://github.com/moby/moby/pull/36298)
- Fix race in attachable network attachment [moby/moby#36191](https://github.com/moby/moby/pull/36191)
- Fix the network option table [docker/cli#848](https://github.com/docker/cli/pull/848)
- Fix timeout issue of `InspectNetwork` on AArch64 [moby/moby#36257](https://github.com/moby/moby/pull/36257)
* Verbose info is missing for partial overlay ID [moby/moby#35989](https://github.com/moby/moby/pull/35989)

### Runtime

* Enable HotAdd for Windows [moby/moby#35414](https://github.com/moby/moby/pull/35414)
* LCOW: Graphdriver fix deadlock in hotRemoveVHDs [moby/moby#36114](https://github.com/moby/moby/pull/36114)
* LCOW: Regular mount if only one layer [moby/moby#36052](https://github.com/moby/moby/pull/36052)
* Remove interim env var LCOW_API_PLATFORM_IF_OMITTED [moby/moby#36269](https://github.com/moby/moby/pull/36269)
* Revendor Microsoft/opengcs @ v0.3.6 [moby/moby#36108](https://github.com/moby/moby/pull/36108)
* Windows: Bump to final RS3 build number [moby/moby#36268](https://github.com/moby/moby/pull/36268)
- Fix issue of ExitCode and PID not show up in Task.Status.ContainerStatus [moby/moby#36150](https://github.com/moby/moby/pull/36150)
- Fix issue with plugin scanner going to deep [moby/moby#36119](https://github.com/moby/moby/pull/36119)
* Do not make graphdriver homes private mounts. [moby/moby#36047](https://github.com/moby/moby/pull/36047)
* Do not recursive unmount on cleanup of zfs/btrfs [moby/moby#36237](https://github.com/moby/moby/pull/36237)
* Don't restore image if layer does not exist [moby/moby#36304](https://github.com/moby/moby/pull/36304)
** Adjust minimum API version for templated configs/secrets [moby/moby#36366](https://github.com/moby/moby/pull/36366)
* Bump containerd to 1.0.2 (cfd04396dc68220d1cecbe686a6cc3aa5ce3667c) [moby/moby#36308](https://github.com/moby/moby/pull/36308)
* Bump Golang to 1.9.4 [moby/moby#36243](https://github.com/moby/moby/pull/36243)
* Ensure daemon root is unmounted on shutdown [moby/moby#36107](https://github.com/moby/moby/pull/36107)
- Fix import path [moby/moby#36322](https://github.com/moby/moby/pull/36322)
* Update runc to 6c55f98695e902427906eed2c799e566e3d3dfb5 [moby/moby#36222](https://github.com/moby/moby/pull/36222)
- Fix container cleanup on daemon restart [moby/moby#36249](https://github.com/moby/moby/pull/36249)
* Bump golang to 1.9.4 [docker/cli#868](https://github.com/docker/cli/pull/868)
* Support SCTP port mapping (bump up API to v1.37) [moby/moby#33922](https://github.com/moby/moby/pull/33922)
* Support SCTP port mapping [docker/cli#278](https://github.com/docker/cli/pull/278)
- Fix Volumes property definition in ContainerConfig [moby/moby#35946](https://github.com/moby/moby/pull/35946)
* Bump moby and dependencies [docker/cli#829](https://github.com/docker/cli/pull/829)
* Bump moby to 0ede01237c9ab871f1b8db0364427407f3e46541 [docker/cli#894](https://github.com/docker/cli/pull/894)
* Bump moby vendor and dependencies [docker/cli#892](https://github.com/docker/cli/pull/892)
* C.RWLayer: check for nil before use [moby/moby#36242](https://github.com/moby/moby/pull/36242)
+ Add `REMOVE` and `ORPHANED` to TaskState [moby/moby#36146](https://github.com/moby/moby/pull/36146)
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
* Templated secrets and configs [moby/moby#33702](https://github.com/moby/moby/pull/33702)
* Use rslave propagation for mounts from daemon root [moby/moby#36055](https://github.com/moby/moby/pull/36055)
+ Add /proc/keys to masked paths [moby/moby#36368](https://github.com/moby/moby/pull/36368)

### Swarm Mode

* Bump SwarmKit to f74983e7c015a38a81c8642803a78b8322cf7eac [moby/moby#36274](https://github.com/moby/moby/pull/36274)
* Clarify network plugins and swarm mode [docker/cli#869](https://github.com/docker/cli/pull/869)
* Migrates several swarm configs tests from integration-cli to api tests [moby/moby#36291](https://github.com/moby/moby/pull/36291)
* Migrates several swarm secrets from integration-cli to api tests [moby/moby#36283](https://github.com/moby/moby/pull/36283)
* Update swarmkit to 68a376dc30d8c4001767c39456b990dbd821371b [moby/moby#36131](https://github.com/moby/moby/pull/36131)
* [compose] Share the compose loading code between swarm and k8s stack deploy [docker/cli#845](https://github.com/docker/cli/pull/845)
+ Add swarm types to bash completion event type filter [docker/cli#888](https://github.com/docker/cli/pull/888)
- Fix issue where network inspect does not show Created time for networks in swarm scope [moby/moby#36095](https://github.com/moby/moby/pull/36095)
