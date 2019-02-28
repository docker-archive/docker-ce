# Changelog

For official release notes for Docker Engine CE and Docker Engine EE, visit the
[https://docs.docker.com/engine/release-notes/](release notes page).

## 18.09.3 (2019-02-28)

### Networking

- Windows: avoid regeneration of network ids to prevent broken references to networks. [docker/engine#149](https://github.com/docker/engine/pull/149)

### Runtime

* Update to Go 1.10.8.
* Modify some of the names in the container name generator. [docker/engine#159](https://github.com/docker/engine/pull/159)
- When copying existing folder, ignore xattr set errors when the target filesystem doesn't support xattr. [docker/engine#135](https://github.com/docker/engine/pull/135)
- Graphdriver: fix "device" mode not being detected if "character-device" bit is set. [docker/engine#160](https://github.com/docker/engine/pull/160)
- Fix nil pointer derefence on failure to connect to containerd. [docker/engine#162](https://github.com/docker/engine/pull/162)
- Delete stale containerd object on start failure. [docker/engine#154](https://github.com/docker/engine/pull/154)

## 18.09.2 (2019-02-11)

### Security

- Update `runc` to address a critical vulnerability that allows specially-crafted containers to gain administrative privileges on the host. ([CVE-2019-5736](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2019-5736))

## 18.09.1 (2019-01-09)

### Builder

- Fix inefficient networking config. [docker/engine#123](https://github.com/docker/engine/pull/123)
- Fix docker system prune doesn't accept until filter. [docker/engine#122](https://github.com/docker/engine/pull/122)
- Avoid unset credentials in containerd. [docker/engine#122](https://github.com/docker/engine/pull/122)
* Update to BuildKit 0.3.3. [docker/engine#122](https://github.com/docker/engine/pull/122)
+ Additional warnings for use of deprecated legacy overlay and devicemapper storage dirvers. [docker/engine#85](https://github.com/docker/engine/pull/85)

### Client

+ Add bash completion for experimental CLI commands (manifest). [docker/cli#1542](https://github.com/docker/cli/pull/1542)
- Fix yamldocs outputing `[flags]` in usage output. [docker/cli#1540](https://github.com/docker/cli/pull/1540)
- Fix setting default schema to tcp for docker host. [docker/cli#1454](https://github.com/docker/cli/pull/1454)
- prune: perform image pruning before build cache pruning. [docker/cli#1532](https://github.com/docker/cli/pull/1532)
- Fix bash completion for `service update --force`. [docker/cli#1526](https://github.com/docker/cli/pull/1526)

### Networking

- Fix iptables compatibility on debian. [docker/engine#107](https://github.com/docker/engine/pull/107)

### Packaging

+ Add docker.socket requirement for docker.service. [docker/docker-ce-packaging#276](https://github.com/docker/docker-ce-packaging/pull/276)
+ Add socket activation for RHEL-based distributions. [docker/docker-ce-packaging#274](https://github.com/docker/docker-ce-packaging/pull/274)
- Add libseccomp requirement for RPM packages. [docker/docker-ce-packaging#266](https://github.com/docker/docker-ce-packaging/pull/266)

### Runtime

* Add `/proc/asound` to masked paths. [docker/engine#126](https://github.com/docker/engine/pull/126)
* Update to containerd 1.2.1-rc.0. [docker/engine#121](https://github.com/docker/engine/pull/121)
+ Windows: allow process isolation. [docker/engine#81](https://github.com/docker/engine/pull/81)
- Windows: DetachVhd attempt in cleanup [docker/engine#113](https://github.com/docker/engine/pull/113)
- API: properly handle invalid JSON to return a 400 status. [docker/engine#110](https://github.com/docker/engine/pull/110)
- API: ignore default address-pools on API < 1.39. [docker/engine#118](https://github.com/docker/engine/pull/118)
- API: add missing default address pool fields to swagger. [docker/engine#119](https://github.com/docker/engine/pull/119)
- awslogs: account for UTF-8 normalization in limits. [docker/engine#112](https://github.com/docker/engine/pull/112)
- Prohibit reading more than 1MB in HTTP error responses. [docker/engine#114](https://github.com/docker/engine/pull/114)
- apparmor: allow receiving of signals from `docker kill`. [docker/engine#116](https://github.com/docker/engine/pull/116)
- overlay2: use index=off if possible (fix EBUSY on mount). [docker/engine#84](https://github.com/docker/engine/pull/84)

## 18.09.0 (2018-11-08)

### Deprecation

For more information on the list of deprecated flags and APIs, have a look at
https://docs.docker.com/engine/deprecated/ where you can find the target removal dates

* Deprecate devicemapper storage driver [docker/cli#1455](https://github.com/docker/cli/pull/1455) / [docker/cli#1424](https://github.com/docker/cli/pull/1424)
* Deprecate legacy overlay storage driver [docker/cli#1455](https://github.com/docker/cli/pull/1455) / [docker/cli#1425](https://github.com/docker/cli/pull/1425) 
* Remove support for TLS < 1.2 [moby/moby#37660](https://github.com/moby/moby/pull/37660)
* Remove Ubuntu 14.04 "Trusty Tahr" as a supported platform [docker-ce-packaging#255](https://github.com/docker/docker-ce-packaging/pull/255) / [docker-ce-packaging#254](https://github.com/docker/docker-ce-packaging/pull/254)
* Remove Debian 8 "Jessie" as a supported platform [docker-ce-packaging#255](https://github.com/docker/docker-ce-packaging/pull/255) / [docker-ce-packaging#254](https://github.com/docker/docker-ce-packaging/pull/254)


### API

+ Update API version to 1.39 [moby/moby#37640](https://github.com/moby/moby/pull/37640)
+ Add support for remote connections using SSH [docker/cli#1014](https://github.com/docker/cli/pull/1014)
+ Builder: add prune options to the API [moby/moby#37651](https://github.com/moby/moby/pull/37651)
+ Add "Warnings" to `/info` endpoint, and move detection to the daemon [moby/moby#37502](https://github.com/moby/moby/pull/37502)
* Do not return "`<unknown>`" in /info response [moby/moby#37472](https://github.com/moby/moby/pull/37472)


### Builder

+ Allow BuildKit builds to run without experimental mode enabled. Buildkit can now be configured with an option in daemon.json [moby/moby#37593](https://github.com/moby/moby/pull/37593) [moby/moby#37686](https://github.com/moby/moby/pull/37686) [moby/moby#37692](https://github.com/moby/moby/pull/37692) [docker/cli#1303](https://github.com/docker/cli/pull/1303)  [docker/cli#1275](https://github.com/docker/cli/pull/1275)
+ Add support for build-time secrets using a `--secret` flag when using BuildKit [docker/cli#1288](https://github.com/docker/cli/pull/1288)
+ Add SSH agent socket forwarder (`docker build --ssh $SSHMOUNTID=$SSH_AUTH_SOCK`) when using BuildKit [docker/cli#1438](https://github.com/docker/cli/pull/1438) / [docker/cli#1419](https://github.com/docker/cli/pull/1419)
+ Add `--chown` flag support for `ADD` and `COPY` commands on Windows [moby/moby#35521](https://github.com/moby/moby/pull/35521)
+ Add `builder prune` subcommand to prune BuildKit build cache [docker/cli#1295](https://github.com/docker/cli/pull/1295) [docker/cli#1334](https://github.com/docker/cli/pull/1334)
+ BuildKit: Add configurable garbage collection policy for the BuildKit build cache [docker/engine#59](https://github.com/docker/engine/pull/59) / [moby/moby#37846](https://github.com/moby/moby/pull/37846)
+ BuildKit: Add support for `docker build --pull ...` when using BuildKit [moby/moby#37613](https://github.com/moby/moby/pull/37613)
+ BuildKit: Add support or "registry-mirrors" and "insecure-registries" when using BuildKit docker/engine#59](https://github.com/docker/engine/pull/59) / [moby/moby#37852](https://github.com/moby/moby/pull/37852)
+ BuildKit: Enable net modes and bridge. [moby/moby#37620](https://github.com/moby/moby/pull/37620)
* BuildKit: Change `--console=[auto,false,true]` to `--progress=[auto,plain,tty]` [docker/cli#1276](https://github.com/docker/cli/pull/1276)
* BuildKit: Set BuildKit's ExportedProduct variable to show useful errors in the future. [moby/moby#37439](https://github.com/moby/moby/pull/37439)
- BuildKit: Do not cancel buildkit status request. [moby/moby#37597](https://github.com/moby/moby/pull/37597)
- Fix no error is shown if build args are missing during docker build [moby/moby#37396](https://github.com/moby/moby/pull/37396)
- Fix error "unexpected EOF" when adding an 8GB file [moby/moby#37771](https://github.com/moby/moby/pull/37771)
- LCOW: Ensure platform is populated on `COPY`/`ADD`. [moby/moby#37563](https://github.com/moby/moby/pull/37563)


### Client

+ Add `docker engine` subcommand to manage the lifecycle of a Docker Engine running as a privileged container on top of containerd, and to allow upgrades to Docker Engine Enterprise [docker/cli#1260](https://github.com/docker/cli/pull/1260)
+ Expose product license in `docker info` output [docker/cli#1313](https://github.com/docker/cli/pull/1313)
+ Show warnings produced by daemon in `docker info` output [docker/cli#1225](https://github.com/docker/cli/pull/1225)
* Hide `--data-path-addr` flags when connected to a daemon that doesn't support this option [docker/docker/cli#1240](https://github.com/docker/cli/pull/1240)
* Only show buildkit-specific flags if BuildKit is enabled [docker/cli#1438](https://github.com/docker/cli/pull/1438) / [docker/cli#1427](https://github.com/docker/cli/pull/1427)
* Improve version output alignment [docker/cli#1204](https://github.com/docker/cli/pull/1204)
* Sort plugin names and networks in a natural order [docker/cli#1166](https://github.com/docker/cli/pull/1166), [docker/cli#1266](https://github.com/docker/cli/pull/1266)
* Updated bash and zsh [completion scripts](https://github.com/docker/cli/issues?q=label%3Aarea%2Fcompletion+milestone%3A18.09.0+is%3Aclosed)
- Fix mapping a range of host ports to a single container port [docker/cli#1102](https://github.com/docker/cli/pull/1102)
- Fix `trust inspect` typo: "`AdminstrativeKeys`" [docker/cli#1300](https://github.com/docker/cli/pull/1300)
- Fix environment file parsing for imports of absent variables and those with no name. [docker/cli#1019](https://github.com/docker/cli/pull/1019)
- Fix a potential "out of memory exception" when running `docker image prune` with a large list of dangling images [docker/cli#1432](https://github.com/docker/cli/pull/1432) / [docker/cli#1423](https://github.com/docker/cli/pull/1423)
- Fix pipe handling in ConEmu and ConsoleZ on Windows [moby/moby#37600](https://github.com/moby/moby/pull/37600)
- Fix long startup on windows, with non-hns governed Hyper-V networks [docker/engine#67](https://github.com/docker/engine/pull/67) / [moby/moby#37774](https://github.com/moby/moby/pull/37774)


### Daemon

- Fix daemon won't start when "runtimes" option is defined both in config file and cli [docker/engine#57](https://github.com/docker/engine/pull/57) / [moby/moby#37871](https://github.com/moby/moby/pull/37871)
- Loosen permissions on `/etc/docker` directory to prevent "permission denied" errors when using `docker manifest inspect` [docker/engine#56](https://github.com/docker/engine/pull/56) / [moby/moby#37847](https://github.com/moby/moby/pull/37847)
- Fix denial of service with large numbers in `cpuset-cpus` and `cpuset-mems` [docker/engine#70](https://github.com/docker/engine/pull/70) / [moby/moby#37967](https://github.com/moby/moby/pull/37967)


### Experimental

- LCOW: Add `--platform` to `docker import` [docker/cli#1375](https://github.com/docker/cli/pull/1375) / [docker/cli#1371](https://github.com/docker/cli/pull/1371)
- LCOW: Add LinuxMetadata support by default on Windows [moby/moby#37514](https://github.com/moby/moby/pull/37514)
- LCOW: Mount to short container paths to avoid command-line length limit [moby/moby#37659](https://github.com/moby/moby/pull/37659)
- LCOW: Fix builder using wrong cache layer [moby/moby#37356](https://github.com/moby/moby/pull/37356)


### Logging

+ Add "local" log driver [moby/moby#37092](https://github.com/moby/moby/pull/37092)
+ Amazon CloudWatch: add `awslogs-endpoint` logging option [moby/moby#37374](https://github.com/moby/moby/pull/37374)
* Pass log-level to containerd. [moby/moby#37419](https://github.com/moby/moby/pull/37419)
- Fix json-log file descriptors leaking when using `--follow` [docker/engine#48](https://github.com/docker/engine/pull/48) [moby/moby#37576](https://github.com/moby/moby/pull/37576) [moby/moby#37734](https://github.com/moby/moby/pull/37734)
- Fix a possible deadlock on closing the watcher on kqueue [moby/moby#37392](https://github.com/moby/moby/pull/37392)
- Use poller based watcher to work around the file caching issue in Windows [moby/moby#37412](https://github.com/moby/moby/pull/37412)


### Networking

+ Add support for global default address pools [moby/moby#37558](https://github.com/moby/moby/pull/37558) [docker/cli#1233](https://github.com/docker/cli/pull/1233)
* Use direct server return (DSR) in east-west overlay load balancing [docker/engine#93](https://github.com/docker/engine/pull/93) / [docker/libnetwork#2270](https://github.com/docker/libnetwork/pull/2270)
* Builder: temporarily disable bridge networking when using buildkit. [moby/moby#37691](https://github.com/moby/moby/pull/37691)
- Handle systemd-resolved case by providing appropriate resolv.conf to networking layer [moby/moby#37485](https://github.com/moby/moby/pull/37485)


### Runtime

+ Configure containerd log-level to be the same as dockerd [moby/moby#37419](https://github.com/moby/moby/pull/37419)
+ Add configuration option for cri-containerd [moby/moby#37519](https://github.com/moby/moby/pull/37519)
+ Update containerd client to v1.2.0-rc.1 [moby/moby#37664](https://github.com/moby/moby/pull/37664), [docker/engine#75](https://github.com/docker/engine/pull/75) / [moby/moby#37710](https://github.com/moby/moby/pull/37710)


### Security

- Remove support for TLS < 1.2 [moby/moby#37660](https://github.com/moby/moby/pull/37660)
- Seccomp: Whitelist syscalls linked to `CAP_SYS_NICE` in default seccomp profile [moby/moby#37242](https://github.com/moby/moby/pull/37242)
- Seccomp: move the syslog syscall to be gated by `CAP_SYS_ADMIN` or `CAP_SYSLOG` [docker/engine#64](https://github.com/docker/engine/pull/64) / [moby/moby#37929](https://github.com/moby/moby/pull/37929)
- SELinux: Fix relabeling of local volumes specified via Mounts API on selinux-enabled systems [moby/moby#37739](https://github.com/moby/moby/pull/37739)
- Add warning if REST API is accessible through an insecure connection [moby/moby#37684](https://github.com/moby/moby/pull/37684)
- Mask proxy credentials from URL when displayed in system info [docker/engine#72](https://github.com/docker/engine/pull/72) / [moby/moby#37934](https://github.com/moby/moby/pull/37934)


### Storage drivers

- Fix mount propagation for btrfs [docker/engine#86](https://github.com/docker/engine/pull/86) / [moby/moby#38026](https://github.com/moby/moby/pull/38026)


### Swarm Mode

+ Add support for global default address pools [moby/moby#37558](https://github.com/moby/moby/pull/37558) [docker/cli#1233](https://github.com/docker/cli/pull/1233)
* Block task starting until node attachments are ready [moby/moby#37604](https://github.com/moby/moby/pull/37604)
* Propagate the provided external CA certificate to the external CA object in swarm. [docker/cli#1178](https://github.com/docker/cli/pull/1178)
- Fix nil pointer dereference in node allocation [docker/engine#94](https://github.com/docker/engine/pull/94) / [docker/swarmkit#2764](https://github.com/docker/swarmkit/pull/2764)


## Packaging

* Remove Ubuntu 14.04 "Trusty Tahr" as a supported platform [docker-ce-packaging#255](https://github.com/docker/docker-ce-packaging/pull/255) / [docker-ce-packaging#254](https://github.com/docker/docker-ce-packaging/pull/254)
* Remove Debian 8 "Jessie" as a supported platform [docker-ce-packaging#255](https://github.com/docker/docker-ce-packaging/pull/255) / [docker-ce-packaging#254](https://github.com/docker/docker-ce-packaging/pull/254)
* Remove 'docker-' prefix for containerd and runc binaries [docker/engine#61](https://github.com/docker/engine/pull/61) / [moby/moby#37907](https://github.com/moby/moby/pull/37907), [docker-ce-packaging#241](https://github.com/docker/docker-ce-packaging/pull/241)
* Split "engine", "cli", and "containerd" to separate packages, and run containerd as a separate systemd service [docker-ce-packaging#131](https://github.com/docker/docker-ce-packaging/pull/131), [docker-ce-packaging#158](https://github.com/docker/docker-ce-packaging/pull/158)
* Build binaries with Go 1.10.4 [docker-ce-packaging#181](https://github.com/docker/docker-ce-packaging/pull/181)
* Remove `-ce` / `-ee` suffix from version string [docker-ce-packaging#206](https://github.com/docker/docker-ce-packaging/pull/206)
