# Changelog

For more information on the list of deprecated flags and APIs, have a look at
https://docs.docker.com/engine/deprecated/ where you can find the target removal dates

## 18.09.0-ce (2018-MM-DD)

### Builder

* -buildmode=pie is not supported on Linux on MIPS either. [moby/moby#37489](https://github.com/moby/moby/pull/37489)
* Build --secret with buildkit. [docker/cli#1288](https://github.com/docker/cli/pull/1288)
* Build: Remove API requirement for --progress as it is CLI only. [docker/cli#1296](https://github.com/docker/cli/pull/1296)
* Build: add -buildmode=pie. [docker/cli#1242](https://github.com/docker/cli/pull/1242)
* Build: change --console=[auto,false,true] to --progress=[auto,plain,tty]. [docker/cli#1276](https://github.com/docker/cli/pull/1276)
* Builder: Implement `builder prune` to prune build cache. [docker/cli#1295](https://github.com/docker/cli/pull/1295)
* Builder: do not cancel buildkit status request. [moby/moby#37597](https://github.com/moby/moby/pull/37597)
* Builder: implement PullParent option with buildkit. [moby/moby#37613](https://github.com/moby/moby/pull/37613)
* Buildkit: enable net modes and bridge. [moby/moby#37620](https://github.com/moby/moby/pull/37620)
* Import TestBuildIidFileSquash from moby to cli. [docker/cli#990](https://github.com/docker/cli/pull/990)
* Re-add support for a custom .bashrc file in build env. [moby/moby#37370](https://github.com/moby/moby/pull/37370)
* Remove "experimental" annotations for buildkit. [docker/cli#1303](https://github.com/docker/cli/pull/1303)
* Remove experimental guard for buildkit. [moby/moby#37686](https://github.com/moby/moby/pull/37686)
* Set BuildKit's ExportedProduct variable to show useful errors in the future. [moby/moby#37439](https://github.com/moby/moby/pull/37439)
* Vendor: update buildkit to 9acf51e491. [moby/moby#37385](https://github.com/moby/moby/pull/37385)
* Vndr buildkit, containerd, console. [docker/cli#1302](https://github.com/docker/cli/pull/1302)
* Vndr containerd to a88b631, buildkit to e57eed4, and fsutil to b19464c. [moby/moby#37582](https://github.com/moby/moby/pull/37582)
* [enhancement] add optional fields in daemon.json to enable buildkit. [moby/moby#37593](https://github.com/moby/moby/pull/37593)
* [enhancement] enable buildkit from daemon side. [docker/cli#1275](https://github.com/docker/cli/pull/1275)
+ Add osusergo build tar for static binaries. [moby/moby#37500](https://github.com/moby/moby/pull/37500)

### Client

* Exposes compose `loader.Transform` function…. [docker/cli#1244](https://github.com/docker/cli/pull/1244)
* Extract StackConverter from the StackClient. [docker/cli#1152](https://github.com/docker/cli/pull/1152)
* Remove composefiles length check on k8s RunDeploy. [docker/cli#1172](https://github.com/docker/cli/pull/1172)
+ Add a doc.go file so the compose/schema/data directory can be vendore…. [docker/cli#1169](https://github.com/docker/cli/pull/1169)
+ Add a new `ExtractVariables` function to `compose/template` package. [docker/cli#1249](https://github.com/docker/cli/pull/1249)
+ Add missing fields in compose/types. [docker/cli#1235](https://github.com/docker/cli/pull/1235)
+ Add omitempty on compose config top-level types. [docker/cli#1170](https://github.com/docker/cli/pull/1170)

### Logging

* Integration-cli: error logging improvements. [moby/moby#37635](https://github.com/moby/moby/pull/37635)
* Integration: fix log message. [moby/moby#37542](https://github.com/moby/moby/pull/37542)
* Loggerutils: fix a typo. [moby/moby#37570](https://github.com/moby/moby/pull/37570)
* Pass endpoint to the CloudWatch Logs logging driver. [moby/moby#37374](https://github.com/moby/moby/pull/37374)
* Pass log-level to containerd. [moby/moby#37419](https://github.com/moby/moby/pull/37419)
+ Add "local" log driver. [moby/moby#37092](https://github.com/moby/moby/pull/37092)
- Fix logic when enabling buildkit. [moby/moby#37688](https://github.com/moby/moby/pull/37688)

### Networking

* Builder: temporarily disable bridge networking when using buildkit. [moby/moby#37691](https://github.com/moby/moby/pull/37691)
* Move network conversions out of API router. [moby/moby#37156](https://github.com/moby/moby/pull/37156)
* Use sortorder library for sorting network list output. [docker/cli#1266](https://github.com/docker/cli/pull/1266)
- Fix faulty error type checking in removeNetwork(). [moby/moby#37409](https://github.com/moby/moby/pull/37409)

### Runtime

* Disable TestExecWindowsOpenHandles on RS5 temporarily. [moby/moby#37679](https://github.com/moby/moby/pull/37679)
* LCOW: Ensure platform is populated on COPY/ADD. [moby/moby#37563](https://github.com/moby/moby/pull/37563)
* LCOW: Mount to short container paths to avoid command-line length limit. [moby/moby#37659](https://github.com/moby/moby/pull/37659)
* LCOW: lazycontext: Use correct lstat, fix archive check. [moby/moby#37356](https://github.com/moby/moby/pull/37356)
* Lcow: fix debug in startServiceVMIfNotRunning(). [moby/moby#37446](https://github.com/moby/moby/pull/37446)
* Revendor Microsoft/opengcs @ v0.3.8. [moby/moby#37515](https://github.com/moby/moby/pull/37515)
* Select polling based watcher for docker log file watcher on Windows. [moby/moby#37412](https://github.com/moby/moby/pull/37412)
* Updated the go-winio library to release 0.4.8 that has the fix for Windows Container. [docker/cli#1157](https://github.com/docker/cli/pull/1157)
+ Add --chown flag support for ADD/COPY commands for Windows. [moby/moby#35521](https://github.com/moby/moby/pull/35521)
+ Adds LinuxMetadata support by default on Windows. [moby/moby#37514](https://github.com/moby/moby/pull/37514)

### Swarm Mode

* Propagate the provided external CA certificate to the external CA object in swarm. [docker/cli#1178](https://github.com/docker/cli/pull/1178)
