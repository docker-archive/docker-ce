# Docker CE

This repository hosts open source components of Docker Community Edition
(CE) products. The `master` branch serves to unify the upstream components
on a regular basis. Long-lived release branches host the code that goes
into a product version for the lifetime of the product.

## Unifying upstream sources

The `master` branch is a combination of components adapted from different
upstream git repos into a unified directory structure. Git history is
preserved for each component.

You can view the upstream git repos in the
[components.conf](components.conf) file. Each component is isolated into
its own directory under the [components](components) directory.

## Updates to `master` branch

Main development of new features should be directed towards the upstream
repos. The `master` branch of this repo will periodically pull in new
changes from upstream to provide a point for integration.

## Branching for release

When a new release is started for Docker CE, it will branch from `master`.
