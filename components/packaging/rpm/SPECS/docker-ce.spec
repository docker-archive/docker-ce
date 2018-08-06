%global debug_package %{nil}

Name: docker-ce
Version: %{_version}
Release: %{_release}%{?dist}
Epoch: %{getenv:EPOCH}
Summary: The open-source application container engine
Group: Tools/Docker
License: ASL 2.0
URL: https://www.docker.com
Vendor: Docker
Packager: Docker <support@docker.com>

Requires: docker-ce-cli

# conflicting packages
Conflicts: docker
Conflicts: docker-io
Conflicts: docker-engine-cs
Conflicts: docker-ee

# Obsolete packages
Obsoletes: docker-ce-selinux
Obsoletes: docker-engine-selinux
Obsoletes: docker-engine

%description
Docker is an open source project to build, ship and run any application as a
lightweight container.

Docker containers are both hardware-agnostic and platform-agnostic. This means
they can run anywhere, from your laptop to the largest EC2 compute instance and
everything in between - and they don't require you to use a particular
language, framework or packaging system. That makes them great building blocks
for deploying and scaling web apps, databases, and backend services without
depending on a particular stack or provider.

%install

%files

%post
if ! getent group docker > /dev/null; then
    groupadd --system docker
fi
# TODO Needs upgrade vs. install logic handling here
if ctr --namespace docker container info dockerd > /dev/null 2>&1 ; then
    docker engine init
fi

%changelog
