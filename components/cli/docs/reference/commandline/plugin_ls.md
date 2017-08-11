---
title: "plugin ls"
description: "The plugin ls command description and usage"
keywords: "plugin, list"
---

<!-- This file is maintained within the docker/cli Github
     repository at https://github.com/docker/cli/. Make all
     pull requests against that repo. If you see this file in
     another repository, consider it read-only there, as it will
     periodically be overwritten by the definitive file. Pull
     requests which include edits to this file in other repositories
     will be rejected.
-->

# plugin ls

```markdown
Usage:  docker plugin ls [OPTIONS]

List plugins

Aliases:
  ls, list

Options:
  -f, --filter filter   Provide filter values (e.g. 'enabled=true')
      --format string   Pretty-print plugins using a Go template
      --help            Print usage
      --no-trunc        Don't truncate output
  -q, --quiet           Only display plugin IDs
```

## Description

Lists all the plugins that are currently installed. You can install plugins
using the [`docker plugin install`](plugin_install.md) command.
You can also filter using the `-f` or `--filter` flag.
Refer to the [filtering](#filtering) section for more information about available filter options.

## Examples

```bash
$ docker plugin ls

ID                  NAME                             TAG                 DESCRIPTION                ENABLED
69553ca1d123        tiborvass/sample-volume-plugin   latest              A test plugin for Docker   true
```

### Filtering

The filtering flag (`-f` or `--filter`) format is of "key=value". If there is more
than one filter, then pass multiple flags (e.g., `--filter "foo=bar" --filter "bif=baz"`)

The currently supported filters are:

* enabled (boolean - true or false, 0 or 1)
* capability (string - currently `volumedriver`, `networkdriver`, `ipamdriver`, `logdriver`, `metricscollector`, or `authz`)

#### enabled

The `enabled` filter matches on plugins enabled or disabled.

#### capability

The `capability` filter matches on plugin capabilities. One plugin
might have multiple capabilities. Currently `volumedriver`, `networkdriver`,
`ipamdriver`, `logdriver`, `metricscollector`, and `authz` are supported capabilities.

```bash
$ docker plugin install --disable vieux/sshfs

Installed plugin vieux/sshfs

$ docker plugin ls --filter enabled=true

NAME                  TAG                 DESCRIPTION                ENABLED
```

### Formatting

The formatting options (`--format`) pretty-prints plugins output
using a Go template.

Valid placeholders for the Go template are listed below:

Placeholder    | Description
---------------|------------------------------------------------------------------------------------------
`.ID`              | Plugin ID
`.Name`            | Plugin name
`.Description`     | Plugin description
`.Enabled`         | Whether plugin is enabled or not
`.PluginReference` | The reference used to push/pull from a registry

When using the `--format` option, the `plugin ls` command will either
output the data exactly as the template declares or, when using the
`table` directive, includes column headers as well.

The following example uses a template without headers and outputs the
`ID` and `Name` entries separated by a colon for all plugins:

```bash
$ docker plugin ls --format "{{.ID}}: {{.Name}}"

4be01827a72e: vieux/sshfs:latest
```

## Related commands

* [plugin create](plugin_create.md)
* [plugin disable](plugin_disable.md)
* [plugin enable](plugin_enable.md)
* [plugin inspect](plugin_inspect.md)
* [plugin install](plugin_install.md)
* [plugin push](plugin_push.md)
* [plugin rm](plugin_rm.md)
* [plugin set](plugin_set.md)
* [plugin upgrade](plugin_upgrade.md)
