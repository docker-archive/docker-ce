# Changelog

For more information on the list of deprecated flags and APIs, have a look at
https://docs.docker.com/engine/deprecated/ where you can find the target removal dates

## 18.09.0-ce (2018-MM-DD)

### Builder

* Build --secret with buildkit. [docker/cli#1288](https://github.com/docker/cli/pull/1288)
* Build: change --console=[auto,false,true] to --progress=[auto,plain,tty]. [docker/cli#1276](https://github.com/docker/cli/pull/1276)
* Builder: Implement `builder prune` to prune build cache. [docker/cli#1295](https://github.com/docker/cli/pull/1295)
* Builder: do not cancel buildkit status request. [moby/moby#37597](https://github.com/moby/moby/pull/37597)
* Buildkit: implement PullParent option with buildkit. [moby/moby#37613](https://github.com/moby/moby/pull/37613)
* Buildkit: enable net modes and bridge. [moby/moby#37620](https://github.com/moby/moby/pull/37620)
* Remove "experimental" annotations for buildkit. [docker/cli#1303](https://github.com/docker/cli/pull/1303)
* Remove experimental guard for buildkit. [moby/moby#37686](https://github.com/moby/moby/pull/37686)
* Buildkit: Set BuildKit's ExportedProduct variable to show useful errors in the future. [moby/moby#37439](https://github.com/moby/moby/pull/37439)
* [enhancement] add optional fields in daemon.json to enable buildkit. [moby/moby#37593](https://github.com/moby/moby/pull/37593)
* [enhancement] enable buildkit from daemon side. [docker/cli#1275](https://github.com/docker/cli/pull/1275)

### Client

+ Add missing fields in compose/types. [docker/cli#1235](https://github.com/docker/cli/pull/1235)

### Logging

* Pass endpoint to the CloudWatch Logs logging driver. [moby/moby#37374](https://github.com/moby/moby/pull/37374)
* Pass log-level to containerd. [moby/moby#37419](https://github.com/moby/moby/pull/37419)
+ Add "local" log driver. [moby/moby#37092](https://github.com/moby/moby/pull/37092)

### Networking

* Builder: temporarily disable bridge networking when using buildkit. [moby/moby#37691](https://github.com/moby/moby/pull/37691)
* Use sortorder library for sorting network list output. [docker/cli#1266](https://github.com/docker/cli/pull/1266)
- Fix faulty error type checking in removeNetwork(). [moby/moby#37409](https://github.com/moby/moby/pull/37409)

### Runtime

* LCOW: Ensure platform is populated on COPY/ADD. [moby/moby#37563](https://github.com/moby/moby/pull/37563)
* LCOW: Mount to short container paths to avoid command-line length limit. [moby/moby#37659](https://github.com/moby/moby/pull/37659)
* LCOW: lazycontext: Use correct lstat, fix archive check. [moby/moby#37356](https://github.com/moby/moby/pull/37356)
* Select polling based watcher for docker log file watcher on Windows. [moby/moby#37412](https://github.com/moby/moby/pull/37412)
* Updated the go-winio library to release 0.4.8 that has the fix for Windows Container. [docker/cli#1157](https://github.com/docker/cli/pull/1157)
+ Add --chown flag support for ADD/COPY commands for Windows. [moby/moby#35521](https://github.com/moby/moby/pull/35521)
+ Adds LinuxMetadata support by default on Windows. [moby/moby#37514](https://github.com/moby/moby/pull/37514)

### Swarm Mode

* Propagate the provided external CA certificate to the external CA object in swarm. [docker/cli#1178](https://github.com/docker/cli/pull/1178)

