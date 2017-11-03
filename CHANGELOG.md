# Changelog

Items starting with `DEPRECATE` are important deprecation notices. For more
information on the list of deprecated flags and APIs please have a look at
https://docs.docker.com/engine/deprecated/ where target removal dates can also
be found.

## 17.11.0-ce (2017-11-DD)

IMPORTANT: Docker CE 17.11 is the first Docker release based on
[containerd 1.0 beta](https://github.com/containerd/containerd/releases/tag/v1.0.0-beta.2).
Docker CE 17.11 and later won't recognize containers started with
previous Docker versions. If using
[Live Restore](https://docs.docker.com/engine/admin/live-restore/#enable-the-live-restore-option),
you must stop all containers before upgrading to Docker CE 17.11.
If you don't, any containers started by Docker versions that predate
17.11 won't be recognized by Docker after the upgrade and will keep
running, un-managed, on the system.

NOTE: This is the first release with official support for Ubuntu 17.10 (Artful) and Debian Buster
