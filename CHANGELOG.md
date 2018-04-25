# Changelog

For more information on the list of deprecated flags and APIs, have a look at
https://docs.docker.com/engine/deprecated/ where you can find the target removal dates

## 18.05.0-ce (2018-05-DD)

### ---Delete---

* Bump Golang to 1.9.5. [docker/cli#986](https://github.com/docker/cli/pull/986)
* Update Golang to 1.9.5. [moby/moby#36779](https://github.com/moby/moby/pull/36779)
* Cleaning some manifest documentation typos. [docker/cli#1022](https://github.com/docker/cli/pull/1022)
* Ci: quote bash variable. [moby/moby#36727](https://github.com/moby/moby/pull/36727)
* Bump version to 18.05.0-dev. [docker/cli#992](https://github.com/docker/cli/pull/992)
* E2e integration cli run. [moby/moby#36631](https://github.com/moby/moby/pull/36631)
* Migrate image tag tests from integration-cli to api tests. [moby/moby#36841](https://github.com/moby/moby/pull/36841)
* Migrate test-integration-cli experimental ipvlan test to integration. [moby/moby#36722](https://github.com/moby/moby/pull/36722)
* Migrate test-integration-cli experimental macvlan test to integration. [moby/moby#36697](https://github.com/moby/moby/pull/36697)
* Migrate test-integration-cli experimental plugin tests to integration. [moby/moby#36886](https://github.com/moby/moby/pull/36886)
* More integration-cli/integration refactoring + request package. [moby/moby#36832](https://github.com/moby/moby/pull/36832)
- Fix tests for pkg/archive. [moby/moby#36770](https://github.com/moby/moby/pull/36770)
- Fix version mismatch in API the documentation. [moby/moby#36927](https://github.com/moby/moby/pull/36927)
- Fixed gometalinter errors on test files. [docker/cli#994](https://github.com/docker/cli/pull/994)
* Copy: remove kernel version test. [moby/moby#36736](https://github.com/moby/moby/pull/36736)
* Refactor code in cmd/dockerd/daemon.go. [moby/moby#36845](https://github.com/moby/moby/pull/36845)
* Some improvement in restart_test.go. [moby/moby#36922](https://github.com/moby/moby/pull/36922)
* [docs] Fix typo in manifest command docs: updated `MANFEST` to `MANIFEST`.. [docker/cli#978](https://github.com/docker/cli/pull/978)
* [test/integration-cli] small cleanups of FIXME(s). [moby/moby#36875](https://github.com/moby/moby/pull/36875)
* Skip some tests requires root uid when run as user…. [moby/moby#36913](https://github.com/moby/moby/pull/36913)
* Some enhancement in integration tests. [moby/moby#36862](https://github.com/moby/moby/pull/36862)
* TestDaemonNoSpaceLeftOnDeviceError: simplify. [moby/moby#36744](https://github.com/moby/moby/pull/36744)
+ Add Silvin as maintainer. [docker/cli#1012](https://github.com/docker/cli/pull/1012)
* Make testing helpers as such…. [docker/cli#1011](https://github.com/docker/cli/pull/1011)
* Make testing helpers as such…. [moby/moby#36888](https://github.com/moby/moby/pull/36888)
* Move and refactor integration-cli/registry to internal/test. [moby/moby#36839](https://github.com/moby/moby/pull/36839)
* Build.md: Document --build-arg without value. [docker/cli#999](https://github.com/docker/cli/pull/999)
* Clarify --build-arg documentation. [docker/cli#970](https://github.com/docker/cli/pull/970)
* Migrate test-integration-cli experimental build tests to integration. [moby/moby#36746](https://github.com/moby/moby/pull/36746)
* Migrate TestAPISwarmServicesPlugin to integration. [moby/moby#36865](https://github.com/moby/moby/pull/36865)
* [test/integration] Small daemon refactoring and add swarm init/join helpers. [moby/moby#36854](https://github.com/moby/moby/pull/36854)
* Update Notary vendor to 0.6.1. [docker/cli#997](https://github.com/docker/cli/pull/997)
* Clean some integration-cli/fixtures package/files. [moby/moby#36838](https://github.com/moby/moby/pull/36838)
* Config integration tests use unique resource names. [moby/moby#36806](https://github.com/moby/moby/pull/36806)
* Docs: Typofix in example "docker image ls". [docker/cli#1007](https://github.com/docker/cli/pull/1007)
* Handle some TODOs in tests. [docker/cli#1020](https://github.com/docker/cli/pull/1020)
* Move integration-cli daemon package to internal/test…. [moby/moby#36824](https://github.com/moby/moby/pull/36824)
+ Add myself to poule to get 2w/old PRs. [moby/moby#36866](https://github.com/moby/moby/pull/36866)
+ Add quotation marks for $CURDIR. [moby/moby#36762](https://github.com/moby/moby/pull/36762)
* Update examples to reflect docker-runc's runtime root for plugins.. [docker/cli#988](https://github.com/docker/cli/pull/988)
* Dockerfile: restore yamllint. [moby/moby#36741](https://github.com/moby/moby/pull/36741)
* Use printf, not echo when creating secrets. [docker/cli#979](https://github.com/docker/cli/pull/979)
+ Add target field to build API docs. [moby/moby#36724](https://github.com/moby/moby/pull/36724)
* Move fakecontext, fakegit and fakestorage to internal/test. [moby/moby#36868](https://github.com/moby/moby/pull/36868)

### ---Triage---


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
