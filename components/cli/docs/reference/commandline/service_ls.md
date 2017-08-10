---
title: "service ls"
description: "The service ls command description and usage"
keywords: "service, ls"
---

<!-- This file is maintained within the docker/cli Github
     repository at https://github.com/docker/cli/. Make all
     pull requests against that repo. If you see this file in
     another repository, consider it read-only there, as it will
     periodically be overwritten by the definitive file. Pull
     requests which include edits to this file in other repositories
     will be rejected.
-->

# service ls

```Markdown
Usage:	docker service ls [OPTIONS]

List services

Aliases:
  ls, list

Options:
  -f, --filter filter   Filter output based on conditions provided
      --format string   Pretty-print services using a Go template
      --help            Print usage
  -q, --quiet           Only display IDs
```

## Description

This command when run targeting a manager, lists services are running in the
swarm.

## Examples

On a manager node:

```bash
$ docker service ls

ID            NAME      MODE        REPLICAS    IMAGE
c8wgl7q4ndfd  frontend  replicated  5/5         nginx:alpine
dmu1ept4cxcf  redis     replicated  3/3         redis:3.0.6
iwe3278osahj  mongo     global      7/7         mongo:3.3
```

The `REPLICAS` column shows both the *actual* and *desired* number of tasks for
the service.

### Filtering

The filtering flag (`-f` or `--filter`) format is of "key=value". If there is more
than one filter, then pass multiple flags (e.g., `--filter "foo=bar" --filter "bif=baz"`)

The currently supported filters are:

* [id](service_ls.md#id)
* [label](service_ls.md#label)
* [mode](service_ls.md#mode)
* [name](service_ls.md#name)

#### id

The `id` filter matches all or part of a service's id.

```bash
$ docker service ls -f "id=0bcjw"
ID            NAME   MODE        REPLICAS  IMAGE
0bcjwfh8ychr  redis  replicated  1/1       redis:3.0.6
```

#### label

The `label` filter matches services based on the presence of a `label` alone or
a `label` and a value.

The following filter matches all services with a `project` label regardless of
its value:

```bash
$ docker service ls --filter label=project
ID            NAME       MODE        REPLICAS  IMAGE
01sl1rp6nj5u  frontend2  replicated  1/1       nginx:alpine
36xvvwwauej0  frontend   replicated  5/5       nginx:alpine
74nzcxxjv6fq  backend    replicated  3/3       redis:3.0.6
```

The following filter matches only services with the `project` label with the
`project-a` value.

```bash
$ docker service ls --filter label=project=project-a
ID            NAME      MODE        REPLICAS  IMAGE
36xvvwwauej0  frontend  replicated  5/5       nginx:alpine
74nzcxxjv6fq  backend   replicated  3/3       redis:3.0.6
```

#### mode

The `mode` filter matches on the mode (either `replicated` or `global`) of a service.

The following filter matches only `global` services.

```bash
$ docker service ls --filter mode=global
ID                  NAME                MODE                REPLICAS            IMAGE
w7y0v2yrn620        top                 global              1/1                 busybox
```

#### name

The `name` filter matches on all or part of a service's name.

The following filter matches services with a name containing `redis`.

```bash
$ docker service ls --filter name=redis
ID            NAME   MODE        REPLICAS  IMAGE
0bcjwfh8ychr  redis  replicated  1/1       redis:3.0.6
```

### Formatting

The formatting options (`--format`) pretty-prints services output
using a Go template.

Valid placeholders for the Go template are listed below:

Placeholder | Description
------------|------------------------------------------------------------------------------------------
`.ID`       | Service ID
`.Name`     | Service name
`.Mode`     | Service mode (replicated, global)
`.Replicas` | Service replicas
`.Image`    | Service image
`.Ports`    | Service ports published in ingress mode

When using the `--format` option, the `service ls` command will either
output the data exactly as the template declares or, when using the
`table` directive, includes column headers as well.

The following example uses a template without headers and outputs the
`ID`, `Mode`, and `Replicas` entries separated by a colon for all services:

```bash
$ docker service ls --format "{{.ID}}: {{.Mode}} {{.Replicas}}"

0zmvwuiu3vue: replicated 10/10
fm6uf97exkul: global 5/5
```

## Related commands

* [service create](service_create.md)
* [service inspect](service_inspect.md)
* [service logs](service_logs.md)
* [service rm](service_rm.md)
* [service scale](service_scale.md)
* [service ps](service_ps.md)
* [service update](service_update.md)
