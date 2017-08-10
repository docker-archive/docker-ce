---
title: "export"
description: "The export command description and usage"
keywords: "export, file, system, container"
---

<!-- This file is maintained within the docker/cli Github
     repository at https://github.com/docker/cli/. Make all
     pull requests against that repo. If you see this file in
     another repository, consider it read-only there, as it will
     periodically be overwritten by the definitive file. Pull
     requests which include edits to this file in other repositories
     will be rejected.
-->

# export

```markdown
Usage:  docker export [OPTIONS] CONTAINER

Export a container's filesystem as a tar archive

Options:
      --help            Print usage
  -o, --output string   Write to a file, instead of STDOUT
```

## Description

The `docker export` command does not export the contents of volumes associated
with the container. If a volume is mounted on top of an existing directory in
the container, `docker export` will export the contents of the *underlying*
directory, not the contents of the volume.

Refer to [Backup, restore, or migrate data volumes](https://docs.docker.com/engine/tutorials/dockervolumes/#backup-restore-or-migrate-data-volumes)
in the user guide for examples on exporting data in a volume.

## Examples

Each of these commands has the same result.

```bash
$ docker export red_panda > latest.tar
```

```bash
$ docker export --output="latest.tar" red_panda
```
