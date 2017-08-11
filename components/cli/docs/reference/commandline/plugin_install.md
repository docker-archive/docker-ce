---
title: "plugin install"
description: "the plugin install command description and usage"
keywords: "plugin, install"
---

<!-- This file is maintained within the docker/cli Github
     repository at https://github.com/docker/cli/. Make all
     pull requests against that repo. If you see this file in
     another repository, consider it read-only there, as it will
     periodically be overwritten by the definitive file. Pull
     requests which include edits to this file in other repositories
     will be rejected.
-->

# plugin install

```markdown
Usage:  docker plugin install [OPTIONS] PLUGIN [KEY=VALUE...]

Install a plugin

Options:
      --alias string            Local name for plugin
      --disable                 Do not enable the plugin on install
      --disable-content-trust   Skip image verification (default true)
      --grant-all-permissions   Grant all permissions necessary to run the plugin
      --help                    Print usage
```

## Description

Installs and enables a plugin. Docker looks first for the plugin on your Docker
host. If the plugin does not exist locally, then the plugin is pulled from
the registry. Note that the minimum required registry version to distribute
plugins is 2.3.0

## Examples

The following example installs `vieus/sshfs` plugin and [sets](plugin_set.md) its
`DEBUG` environment variable to `1`. To install, `pull` the plugin from Docker
Hub and prompt the user to accept the list of privileges that the plugin needs,
set the plugin's parameters and enable the plugin.

```bash
$ docker plugin install vieux/sshfs DEBUG=1

Plugin "vieux/sshfs" is requesting the following privileges:
 - network: [host]
 - device: [/dev/fuse]
 - capabilities: [CAP_SYS_ADMIN]
Do you grant the above permissions? [y/N] y
vieux/sshfs
```

After the plugin is installed, it appears in the list of plugins:

```bash
$ docker plugin ls

ID                  NAME                  TAG                 DESCRIPTION                ENABLED
69553ca1d123        vieux/sshfs           latest              sshFS plugin for Docker    true
```

## Related commands

* [plugin create](plugin_create.md)
* [plugin disable](plugin_disable.md)
* [plugin enable](plugin_enable.md)
* [plugin inspect](plugin_inspect.md)
* [plugin ls](plugin_ls.md)
* [plugin push](plugin_push.md)
* [plugin rm](plugin_rm.md)
* [plugin set](plugin_set.md)
* [plugin upgrade](plugin_upgrade.md)
