# Changelog

For official release notes for Docker Engine CE, visit the
[release notes page](https://docs.docker.com/engine/release-notes/).

## 20.10.0

### Deprecation / Removal

- Sterner warnings and deprecation notice for unauthenticated tcp access [moby/moby#41285](https://github.com/moby/moby/pull/41285)
- Deprecate KernelMemory (`docker run --kernel-memory`) [moby/moby#41254](https://github.com/moby/moby/pull/41254) [docker/cli#2652](https://github.com/docker/cli/pull/2652)
- Deprecate `aufs` storage driver [docker/cli#1484](https://github.com/docker/cli/pull/1484)
- Deprecate host-discovery and overlay networks with external k/v stores [moby/moby#40614](https://github.com/moby/moby/pull/40614) [moby/moby#40510](https://github.com/moby/moby/pull/40510)
- Deprecate Dockerfile legacy 'ENV name value' syntax, use `ENV name=value` instead [docker/cli#2743](https://github.com/docker/cli/pull/2743)
- Remove deprecated "filter" parameter for API v1.41 and up [moby/moby#40491](https://github.com/moby/moby/pull/40491)
- Disable distribution manifest v2 schema 1 on push [moby/moby#41295](https://github.com/moby/moby/pull/41295)
- Remove hack MalformedHostHeaderOverride breaking old docker clients (<= 1.12) in which case, set `DOCKER_API_VERSION` [moby/moby#39076](https://github.com/moby/moby/pull/39076)
- Remove "docker engine" subcommands [docker/cli#2207](https://github.com/docker/cli/pull/2207)
- Remove experimental "deploy" from "dab" files [docker/cli#2216](https://github.com/docker/cli/pull/2216)
- Remove deprecated `docker search --automated` and `--stars` flags [docker/cli#2338](https://github.com/docker/cli/pull/2338)
- No longer allow reserved namespaces in engine labels [docker/cli#2326](https://github.com/docker/cli/pull/2326)

### API

- Do not require "experimental" for metrics API [moby/moby#40427](https://github.com/moby/moby/pull/40427)
- `GET /events` now returns `prune` events after pruning resources have completed [moby/moby#41259](https://github.com/moby/moby/pull/41259)
  - Prune events are returned for `container`, `network`, `volume`, `image`, and `builder`, and have a `reclaimed` attribute, indicating the amount of space reclaimed (in bytes)
- Add `one-shot` stats option to not prime the stats [moby/moby#40478](https://github.com/moby/moby/pull/40478)
- Adding OS version info to the system info's API (`/info`) [moby/moby#38349](https://github.com/moby/moby/pull/38349)
- Add DefaultAddressPools to docker info [moby/moby#40714](https://github.com/moby/moby/pull/40714)
- Add API support for PidsLimit on services [moby/moby#39882](https://github.com/moby/moby/pull/39882)

### Builder

- buildkit: git: support for token authentication [moby/moby#41234](https://github.com/moby/moby/pull/41234) [docker/cli#2656](https://github.com/docker/cli/pull/2656) [moby/buildkit#1533](https://github.com/moby/buildkit/pull/1533)
- buildkit: secrets: allow providing secrets with env [moby/moby#41234](https://github.com/moby/moby/pull/41234) [docker/cli#2656](https://github.com/docker/cli/pull/2656) [moby/buildkit#1534](https://github.com/moby/buildkit/pull/1534)
  - Support `--secret id=foo,env=MY_ENV` as an alternative for storing a secret value to a file.
  - `--secret id=GIT_AUTH_TOKEN` will load env if it exists and the file does not.
- buildkit: Support for mirrors fallbacks, insecure TLS and custom TLS config [moby/moby#40814](https://github.com/moby/moby/pull/40814)
- buildkit: remotecache: Only visit each item once when walking results [moby/moby#41234](https://github.com/moby/moby/pull/41234) [moby/buildkit#1577](https://github.com/moby/buildkit/pull/1577)
  - Improves performance and CPU use on bigger graphs
- buildkit: Check remote when local image platform doesn't match [moby/moby#40629](https://github.com/moby/moby/pull/40629)
- buildkit: image export: Use correct media type when creating new layer blobs [moby/moby#41234](https://github.com/moby/moby/pull/41234) [moby/buildkit#1541](https://github.com/moby/buildkit/pull/1541)
- buildkit: progressui: fix logs time formatting [moby/moby#41234](https://github.com/moby/moby/pull/41234) [docker/cli#2656](https://github.com/docker/cli/pull/2656) [moby/buildkit#1549](https://github.com/moby/buildkit/pull/1549)
- buildkit: mitigate containerd issue on parallel push [moby/moby#41234](https://github.com/moby/moby/pull/41234) [moby/buildkit#1548](https://github.com/moby/buildkit/pull/1548)
- buildkit: inline cache: fix handling of duplicate blobs [moby/moby#41234](https://github.com/moby/moby/pull/41234) [moby/buildkit#1568](https://github.com/moby/buildkit/pull/1568)
  - Fixes https://github.com/moby/buildkit/issues/1388 cache-from working unreliably
  - Fixes https://github.com/moby/moby/issues/41219 Image built from cached layers is missing data
- Allow ssh:// for remote context URLs [moby/moby#40179](https://github.com/moby/moby/pull/40179)
- builder: remove legacy build's session handling (was experimental) [moby/moby#39983](https://github.com/moby/moby/pull/39983)

### Client

- Add swarm jobs support to CLI [docker/cli#2262](https://github.com/docker/cli/pull/2262)
- Add `-a/--all-tags` to docker push [docker/cli#2220](https://github.com/docker/cli/pull/2220)
- Add support for Kubernetes username/password auth [docker/cli#2308](https://github.com/docker/cli/pull/2308)
- Add `--pull=missing|always|never` to `run` and `create` commands [docker/cli#1498](https://github.com/docker/cli/pull/1498)
- Add `--env-file` flag to `docker exec` for parsing environment variables from a file [docker/cli#2602](https://github.com/docker/cli/pull/2602)
- Add shorthand `-n` for `--tail` option [docker/cli#2646](https://github.com/docker/cli/pull/2646)
- Add log-driver and options to service inspect "pretty" format [docker/cli#1950](https://github.com/docker/cli/pull/1950)
- docker run: specify cgroup namespace mode with `--cgroupns` [docker/cli#2024](https://github.com/docker/cli/pull/2024)
- `docker manifest rm` command to remove manifest list draft from local storage [docker/cli#2449](https://github.com/docker/cli/pull/2449)
- Add "context" to "docker version" and "docker info" [docker/cli#2500](https://github.com/docker/cli/pull/2500)
- Propagate platform flag to container create API [docker/cli#2551](https://github.com/docker/cli/pull/2551)
- The `docker ps --format` flag now has a `.State` placeholder to print the container's state without additional details about uptime and health check [docker/cli#2000](https://github.com/docker/cli/pull/2000)
- Add support for docker-compose schema v3.9 [docker/cli#2073](https://github.com/docker/cli/pull/2073)
- Add support for docker push `--quiet` [docker/cli#2197](https://github.com/docker/cli/pull/2197)
- Hide flags that are not supported by BuildKit, if BuildKit is enabled [docker/cli#2123](https://github.com/docker/cli/pull/2123)
- Update flag description for `docker rm -v` to clarify the option only removes anonymous (unnamed) volumes [docker/cli#2289](https://github.com/docker/cli/pull/2289)
- Improve tasks printing for docker services [docker/cli#2341](https://github.com/docker/cli/pull/2341)
- docker info: list CLI plugins alphabetically [docker/cli#2236](https://github.com/docker/cli/pull/2236)
- Fix order of processing of `--label-add/--label-rm`, `--container-label-add/--container-label-rm`, and `--env-add/--env-rm` flags on `docker service update` to allow replacing existing values [docker/cli#2668](https://github.com/docker/cli/pull/2668)
- Fix `docker rm --force` returning a non-zero exit code if one or more containers did not exist [docker/cli#2678](https://github.com/docker/cli/pull/2678)
- Improve memory stats display by using `total_inactive_file` instead of `cache` [docker/cli#2415](https://github.com/docker/cli/pull/2415)
- Mitigate against YAML files that has excessive aliasing [docker/cli#2117](https://github.com/docker/cli/pull/2117)
- Allow using advanced syntax when setting a config or secret with only the source field [docker/cli#2243](https://github.com/docker/cli/pull/2243)
- Fix reading config files containing `username` and `password` auth even if `auth` is empty [docker/cli#2122](https://github.com/docker/cli/pull/2122)
- docker cp: prevent NPE when failing to stat destination [docker/cli#2221](https://github.com/docker/cli/pull/2221)
- config: preserve ownership and permissions on configfile [docker/cli#2228](https://github.com/docker/cli/pull/2228)

### Logging

- Support reading `docker logs` with all logging drivers (best effort) [moby/moby#40543](https://github.com/moby/moby/pull/40543)
- Add `splunk-index-acknowledgment` log option to work with Splunk HECs with index acknowledgment enabled [moby/moby#39987](https://github.com/moby/moby/pull/39987)
- Add partial metadata to journald logs [moby/moby#41407](https://github.com/moby/moby/pull/41407)
- Reduce allocations for logfile reader [moby/moby#40796](https://github.com/moby/moby/pull/40796)
- Fluentd: add fluentd-async, fluentd-request-ack, and deprecate fluentd-async-connect [moby/moby#39086](https://github.com/moby/moby/pull/39086)

### Runtime

- Support cgroup2 [moby/moby#40174](https://github.com/moby/moby/pull/40174) [moby/moby#40657](https://github.com/moby/moby/pull/40657) [moby/moby#40662](https://github.com/moby/moby/pull/40662)
- cgroup2: use "systemd" cgroup driver by default when available [moby/moby#40846](https://github.com/moby/moby/pull/40846)
- new storage driver: fuse-overlayfs [moby/moby#40483](https://github.com/moby/moby/pull/40483)
- Update containerd binary to v1.4.0 [moby/moby#40982](https://github.com/moby/moby/pull/40982)
- `docker push` now defaults to `latest` tag instead of all tags [moby/moby#40302](https://github.com/moby/moby/pull/40302)
- Added ability to change the number of reconnect attempts during connection loss while pulling an image by adding max-download-attempts to the config file [moby/moby#39949](https://github.com/moby/moby/pull/39949)
- Add support for containerd v2 shim by using the now default `io.containerd.runc.v2` runtime [moby/moby#41182](https://github.com/moby/moby/pull/41182)
- cgroup v1: change the default runtime to io.containerd.runc.v2. Requires containerd v1.3.0 or later. v1.3.5 or later is recommended [moby/moby#41210](https://github.com/moby/moby/pull/41210)
- Start containers in their own cgroup namespaces [moby/moby#38377](https://github.com/moby/moby/pull/38377)
- Enable DNS Lookups for CIFS Volumes [moby/moby#39250](https://github.com/moby/moby/pull/39250)
- Use MemAvailable instead of MemFree to estimate actual available memory [moby/moby#39481](https://github.com/moby/moby/pull/39481)
- The `--device` flag in `docker run` will now be honored when the container is started in privileged mode [moby/moby#40291](https://github.com/moby/moby/pull/40291)
- Enforce reserved internal labels [moby/moby#40394](https://github.com/moby/moby/pull/40394)
- Raise minimum memory limit to 6M, to account for higher memory use by runtimes during container startup [moby/moby#41168](https://github.com/moby/moby/pull/41168)
- Add support for `CAP_PERFMON`, `CAP_BPF`, and `CAP_CHECKPOINT_RESTORE` on supported kernels [moby/moby#41460](https://github.com/moby/moby/pull/41460)
- vendor runc v1.0.0-rc92 [moby/moby#41344](https://github.com/moby/moby/pull/41344) [moby/moby#41317](https://github.com/moby/moby/pull/41317)
- info: add warnings about missing blkio cgroup support [moby/moby#41083](https://github.com/moby/moby/pull/41083)
- Accept platform spec on container create [moby/moby#40725](https://github.com/moby/moby/pull/40725)
- Fix handling of looking up user- and group-names with spaces [moby/moby#41377](https://github.com/moby/moby/pull/41377)

### Networking

- Support host.docker.internal in dockerd on Linux [moby/moby#40007](https://github.com/moby/moby/pull/40007)
- Include IPv6 address of linked containers in /etc/hosts [moby/moby#39837](https://github.com/moby/moby/pull/39837)
- Add alias for hostname if hostname != container name [moby/moby#39204](https://github.com/moby/moby/pull/39204)
- Better selection of DNS server (with systemd) [moby/moby#41022](https://github.com/moby/moby/pull/41022)
- Add docker interfaces to firewalld docker zone [moby/moby#41189](https://github.com/moby/moby/pull/41189) [moby/libnetwork#2548](https://github.com/moby/libnetwork/pull/2548)
  - Fixes DNS issue on CentOS8 [docker/for-linux#957](https://github.com/docker/for-linux/issues/957)
  - Fixes Port Forwarding on RHEL 8 with Firewalld running with FirewallBackend=nftables [moby/libnetwork#2496](https://github.com/moby/libnetwork/issues/2496)
- Fix an issue reporting 'failed to get network during CreateEndpoint' [moby/moby#41189](https://github.com/moby/moby/pull/41189) [moby/libnetwork#2554](https://github.com/moby/libnetwork/pull/2554)
- Log error instead of disabling IPv6 router advertisement failed [moby/moby#41189](https://github.com/moby/moby/pull/41189) [moby/libnetwork#2563](https://github.com/moby/libnetwork/pull/2563)
- No longer ignore `--default-address-pool` option in certain cases [moby/moby#40711](https://github.com/moby/moby/pull/40711)
- Produce an error with invalid address pool [moby/moby#40808](https://github.com/moby/moby/pull/40808) [moby/libnetwork#2538](https://github.com/moby/libnetwork/pull/2538)
- Fix `DOCKER-USER` chain not created when IPTableEnable=false [moby/moby#40808](https://github.com/moby/moby/pull/40808) [moby/libnetwork#2471](https://github.com/moby/libnetwork/pull/2471)
- Fix panic on startup in systemd environments [moby/moby#40808](https://github.com/moby/moby/pull/40808) [moby/libnetwork#2544](https://github.com/moby/libnetwork/pull/2544)
- Fix issue preventing containers to communicate over macvlan internal network [moby/moby#40596](https://github.com/moby/moby/pull/40596) [moby/libnetwork#2407](https://github.com/moby/libnetwork/pull/2407)
- Fix InhibitIPv4 nil panic [moby/moby#40596](https://github.com/moby/moby/pull/40596)
- Fix VFP leak in Windows overlay network deletion [moby/moby#40596](https://github.com/moby/moby/pull/40596) [moby/libnetwork#2524](https://github.com/moby/libnetwork/pull/2524)

### Packaging

- docker.service: Add multi-user.target to After= in unit file [moby/moby#41297](https://github.com/moby/moby/pull/41297)
- docker.service: Allow socket activation [moby/moby#37470](https://github.com/moby/moby/pull/37470)
- seccomp: Remove dependency in dockerd on libseccomp [moby/moby#41395](https://github.com/moby/moby/pull/41395)

### Rootless

- rootless: graduate from experimental [moby/moby#40759](https://github.com/moby/moby/pull/40759)
- Add dockerd-rootless-setuptool.sh [moby/moby#40950](https://github.com/moby/moby/pull/40950)
- Support `--exec-opt native.cgroupdriver=systemd` [moby/moby#40486](https://github.com/moby/moby/pull/40486)

### Security

- Fix CVE-2019-14271 loading of nsswitch based config inside chroot under Glibc [moby/moby#39612](https://github.com/moby/moby/pull/39612)
- seccomp: Whitelist `clock_adjtime`. `CAP_SYS_TIME` is still required for time adjustment [moby/moby#40929](https://github.com/moby/moby/pull/40929)
- seccomp: Add openat2 and faccessat2 to default seccomp profile [moby/moby#41353](https://github.com/moby/moby/pull/41353)
- seccomp: allow 'rseq' syscall in default seccomp profile [moby/moby#41158](https://github.com/moby/moby/pull/41158)
- seccomp: allow syscall membarrier [moby/moby#40731](https://github.com/moby/moby/pull/40731)
- seccomp: whitelist io-uring related system calls [moby/moby#39415](https://github.com/moby/moby/pull/39415)
- Add default sysctls to allow ping sockets and privileged ports with no capabilities [moby/moby#41030](https://github.com/moby/moby/pull/41030)
- Fix seccomp profile for clone syscall [moby/moby#39308](https://github.com/moby/moby/pull/39308)

### Swarm

- Add support for swarm jobs [moby/moby#40307](https://github.com/moby/moby/pull/40307)
- Add capabilities support to stack/service commands [docker/cli#2687](https://github.com/docker/cli/pull/2687) [docker/cli#2709](https://github.com/docker/cli/pull/2709) [moby/moby#39173](https://github.com/moby/moby/pull/39173) [moby/moby#41249](https://github.com/moby/moby/pull/41249)
- Add support for sending down service Running and Desired task counts [moby/moby#39231](https://github.com/moby/moby/pull/39231)
- service: support `--mount type=bind,bind-nonrecursive` [moby/moby#38788](https://github.com/moby/moby/pull/38788)
- Support ulimits on Swarm services. [moby/moby#41284](https://github.com/moby/moby/pull/41284) [docker/cli#2712](https://github.com/docker/cli/pull/2712)
- Fixed an issue where service logs could leak goroutines on the worker [moby/moby#40426](https://github.com/moby/moby/pull/40426)
