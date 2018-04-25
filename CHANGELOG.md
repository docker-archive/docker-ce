# Changelog

For more information on the list of deprecated flags and APIs, have a look at
https://docs.docker.com/engine/deprecated/ where you can find the target removal dates

## 18.05.0-ce (2018-05-DD)

### Builder

*  Adding `netbsd` compatibility to the package `pkg/term`. [moby/moby#36887](https://github.com/moby/moby/pull/36887)
*  Standardizes output path for artifacts of intermediate builds to `/build/`. [moby/moby#36858](https://github.com/moby/moby/pull/36858)

### Client

- Fix `docker stack deploy` reference flag. [docker/cli#981](https://github.com/docker/cli/pull/981)
- Fix docker stack deploy re-deploying services after the service was updated with `--force`. [docker/cli#963](https://github.com/docker/cli/pull/963)
+ Add bash completion for `secret|config create --template-driver`. [docker/cli#1004](https://github.com/docker/cli/pull/1004)
+ Add fish completions for docker trust subcommand. [docker/cli#984](https://github.com/docker/cli/pull/984)
- Fix --format example for docker history. [docker/cli#980](https://github.com/docker/cli/pull/980)
- Fix error with merge composefile with networks. [docker/cli#983](https://github.com/docker/cli/pull/983)

### Logging
* Standardized the properties of storage-driver log messages. [moby/moby#36492](https://github.com/moby/moby/pull/36492)
* Improve partial message support in logger. [moby/moby#35831](https://github.com/moby/moby/pull/35831)

### Networking

- Allow for larger preset property values, do not override. [docker/libnetwork#2124](https://github.com/docker/libnetwork/pull/2124)
- networkdb: User write lock in handleNodeEvent.  [docker/libnetwork#2136](https://github.com/docker/libnetwork/pull/2136)
* Import libnetwork fix for rolling updates. [moby/moby#36638](https://github.com/moby/moby/pull/36638)
* Update libnetwork to improve scalabiltiy of bridge network isolation rules. [moby/moby#36774](https://github.com/moby/moby/pull/36774)
- Fix a misused network object name. [moby/moby#36745](https://github.com/moby/moby/pull/36745)

### Runtime

* LCOW: Implement `docker save`. [moby/moby#36599](https://github.com/moby/moby/pull/36599)
* Pkg: devmapper: dynamically load dm_task_deferred_remove. [moby/moby#35518](https://github.com/moby/moby/pull/35518)
* Windows: Add GetLayerPath implementation in graphdriver. [moby/moby#36738](https://github.com/moby/moby/pull/36738)
- Fix Windows layer leak when write fails. [moby/moby#36728](https://github.com/moby/moby/pull/36728)
- Fix FIFO, sockets and device files when run in user NS. [moby/moby#36756](https://github.com/moby/moby/pull/36756)
- Fix docker version output alignment. [docker/cli#965](https://github.com/docker/cli/pull/965)
* Always make sysfs read-write with privileged. [moby/moby#36808](https://github.com/moby/moby/pull/36808)
* Bump Golang to 1.10.1. [moby/moby#35739](https://github.com/moby/moby/pull/35739)
* Bump containerd client. [moby/moby#36684](https://github.com/moby/moby/pull/36684)
* Bump golang.org/x/net to go1.10 release commit. [moby/moby#36894](https://github.com/moby/moby/pull/36894)
* Context.WithTimeout: do call the cancel func. [moby/moby#36920](https://github.com/moby/moby/pull/36920)
* Copy: avoid using all system memory with authz plugins. [moby/moby#36595](https://github.com/moby/moby/pull/36595)
* Daemon/cluster: handle partial attachment entries during configure. [moby/moby#36769](https://github.com/moby/moby/pull/36769)
* Don't make container mount unbindable. [moby/moby#36768](https://github.com/moby/moby/pull/36768)
* Extra check before unmounting on shutdown. [moby/moby#36879](https://github.com/moby/moby/pull/36879)
* Move mount parsing to separate package. [moby/moby#36896](https://github.com/moby/moby/pull/36896)
* No global volume driver store. [moby/moby#36637](https://github.com/moby/moby/pull/36637)
* Pkg/mount improvements. [moby/moby#36091](https://github.com/moby/moby/pull/36091)
* Relax some libcontainerd client locking. [moby/moby#36848](https://github.com/moby/moby/pull/36848)
* Remove daemon dependency on api packages. [moby/moby#36912](https://github.com/moby/moby/pull/36912)
* Remove the retries for service update. [moby/moby#36827](https://github.com/moby/moby/pull/36827)
* Revert unencryted storage warning prompt. [docker/cli#1008](https://github.com/docker/cli/pull/1008)
* Support cancellation in `directory.Size()`. [moby/moby#36734](https://github.com/moby/moby/pull/36734)
* Switch from x/net/context -> context. [moby/moby#36904](https://github.com/moby/moby/pull/36904)
* Fixed a function to check Content-type is `application/json` or not. [moby/moby#36778](https://github.com/moby/moby/pull/36778)
+ Add default pollSettings config functions. [moby/moby#36706](https://github.com/moby/moby/pull/36706)
+ Add if judgment before receiving operations on daemonWaitCh. [moby/moby#36651](https://github.com/moby/moby/pull/36651)
- Fix issues with running volume tests as non-root.. [moby/moby#36935](https://github.com/moby/moby/pull/36935)

### Swarm Mode

* RoleManager will remove detected nodes from the cluster membership [docker/swarmkit#2548](https://github.com/docker/swarmkit/pull/2548)
* Scheduler/TaskReaper: handle unassigned tasks marked for shutdown [docker/swarmkit#2574](https://github.com/docker/swarmkit/pull/2574)
* Avoid predefined error log. [docker/swarmkit#2561](https://github.com/docker/swarmkit/pull/2561)
* Task reaper should delete tasks with removed slots that were not yet assigned. [docker/swarmkit#2557](https://github.com/docker/swarmkit/pull/2557)
* Agent reports FIPS status. [docker/swarmkit#2587](https://github.com/docker/swarmkit/pull/2587)
- Fix: timeMutex critical operation outside of critical section. [docker/swarmkit#2603](https://github.com/docker/swarmkit/pull/2603)
* Expose swarmkit's Raft tuning parameters in engine config. [moby/moby#36726](https://github.com/moby/moby/pull/36726)
* Make internal/test/daemon.Daemon swarm aware. [moby/moby#36826](https://github.com/moby/moby/pull/36826)
