# Changelog

For more information on the list of deprecated flags and APIs, have a look at
https://docs.docker.com/engine/deprecated/ where you can find the target removal dates

## 18.04.0-ce (2018-04-DD)

### Builder

- Fix typos in builder and client. [moby/moby#36424](https://github.com/moby/moby/pull/36424)

### Client

* Print Stack API and Kubernetes versions in version command. [docker/cli#898](https://github.com/docker/cli/pull/898)
- Fix Kubernetes duplication in version command. [docker/cli#953](https://github.com/docker/cli/pull/953)
* Use HasAvailableFlags instead of HasFlags for Options in help. [docker/cli#959](https://github.com/docker/cli/pull/959)
+ Add support for mandatory variables to stack deploy. [docker/cli#893](https://github.com/docker/cli/pull/893)
- Fix docker stack services command Port output. [docker/cli#943](https://github.com/docker/cli/pull/943)
* Deprecate unencrypted storage. [docker/cli#561](https://github.com/docker/cli/pull/561)
* Don't set a default filename for ConfigFile. [docker/cli#917](https://github.com/docker/cli/pull/917)
- Fix compose network name. [docker/cli#941](https://github.com/docker/cli/pull/941)

### Logging

* Make LogFile perms configurable. [moby/moby#36523](https://github.com/moby/moby/pull/36523)
* Silent login: use credentials from cred store to login. [docker/cli#139](https://github.com/docker/cli/pull/139)
+ Add support for compressibility of log file. [moby/moby#29932](https://github.com/moby/moby/pull/29932)
- Fix empty LogPath with non-blocking logging mode. [moby/moby#36272](https://github.com/moby/moby/pull/36272)

### Networking

- Prevent explicit removal of ingress network. [moby/moby#36538](https://github.com/moby/moby/pull/36538)

### Runtime

* Devmapper cleanup improvements. [moby/moby#36307](https://github.com/moby/moby/pull/36307)
* Devmapper.Mounted: remove. [moby/moby#36437](https://github.com/moby/moby/pull/36437)
* Devmapper/Remove(): use Rmdir, ignore errors. [moby/moby#36438](https://github.com/moby/moby/pull/36438)
* LCOW - Change platform parser directive to FROM statement flag. [moby/moby#35089](https://github.com/moby/moby/pull/35089)
* Split daemon service code to windows file. [moby/moby#36653](https://github.com/moby/moby/pull/36653)
* Windows: Block pulling uplevel images. [moby/moby#36327](https://github.com/moby/moby/pull/36327)
* Windows: Hyper-V containers are broken after 36586 was merged. [moby/moby#36610](https://github.com/moby/moby/pull/36610)
* Windows: Move kernel_windows to use golang registry functions. [moby/moby#36617](https://github.com/moby/moby/pull/36617)
* Windows: Pass back system errors on container exit. [moby/moby#35967](https://github.com/moby/moby/pull/35967)
* Windows: Remove servicing mode. [moby/moby#36267](https://github.com/moby/moby/pull/36267)
* Windows: Report Version and UBR. [moby/moby#36451](https://github.com/moby/moby/pull/36451)
* Bump Runc to 1.0.0-rc5. [moby/moby#36449](https://github.com/moby/moby/pull/36449)
* Mount failure indicates the path that failed. [moby/moby#36407](https://github.com/moby/moby/pull/36407)
* Change return for errdefs.getImplementer(). [moby/moby#36489](https://github.com/moby/moby/pull/36489)
* Client: fix hijackedconn reading from buffer. [moby/moby#36663](https://github.com/moby/moby/pull/36663)
* Content encoding negotiation added to archive request. [moby/moby#36164](https://github.com/moby/moby/pull/36164)
* Daemon/stats: more resilient cpu sampling. [moby/moby#36519](https://github.com/moby/moby/pull/36519)
* Daemon/stats: remove obnoxious types file. [moby/moby#36494](https://github.com/moby/moby/pull/36494)
* Daemon: use context error rather than inventing new one. [moby/moby#36670](https://github.com/moby/moby/pull/36670)
* Enable CRIU on non-amd64 architectures (v2). [moby/moby#36676](https://github.com/moby/moby/pull/36676)
- Fixes intermittent client hang after closing stdin to attached container [moby/moby#36517](https://github.com/moby/moby/pull/36517)
- Fix daemon panic on container export after restart [moby/moby#36586](https://github.com/moby/moby/pull/36586)
- Follow-up fixes on multi-stage moby's Dockerfile. [moby/moby#36425](https://github.com/moby/moby/pull/36425)
* Freeze busybox and latest glibc in Docker image. [moby/moby#36375](https://github.com/moby/moby/pull/36375)
* If container will run as non root user, drop permitted, effective caps early. [moby/moby#36587](https://github.com/moby/moby/pull/36587)
* Layer: remove metadata store interface. [moby/moby#36504](https://github.com/moby/moby/pull/36504)
* Minor optimizations to dockerd. [moby/moby#36577](https://github.com/moby/moby/pull/36577)
* Whitelist statx syscall. [moby/moby#36417](https://github.com/moby/moby/pull/36417)
+ Add missing error return for plugin creation. [moby/moby#36646](https://github.com/moby/moby/pull/36646)
- Fix AppArmor not being applied to Exec processes. [moby/moby#36466](https://github.com/moby/moby/pull/36466)
* Daemon/logger/ring.go: log error not instance. [moby/moby#36475](https://github.com/moby/moby/pull/36475)
- Fix stats collector spinning CPU if no stats are collected. [moby/moby#36609](https://github.com/moby/moby/pull/36609)
- Fix(distribution): digest cache should not be moved if it was an auth. [moby/moby#36509](https://github.com/moby/moby/pull/36509)

### Swarm Mode

* Fixes for synchronizing the dispatcher shutdown with in-progress rpcs. [moby/moby#36371](https://github.com/moby/moby/pull/36371)
* Increase raft ElectionTick to 10xHeartbeatTick. [moby/moby#36672](https://github.com/moby/moby/pull/36672)
