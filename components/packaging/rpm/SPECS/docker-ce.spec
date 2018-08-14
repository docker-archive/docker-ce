%global debug_package %{nil}

Name: docker-ce
Version: %{_version}
Release: %{_release}%{?dist}
Epoch: %{getenv:EPOCH}
Source0: containerd-proxy.tgz
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

%prep
%setup -q -c -n src

%build
mkdir -p /go/src/github.com/crosbymichael/
ls %{_topdir}/BUILD/src
ln -s %{_topdir}/BUILD/src/containerd-proxy /go/src/github.com/crosbymichael/containerd-proxy
go build -v -o /build/dockerd github.com/crosbymichael/containerd-proxy

%install
install -D -m 0755 /build/dockerd $RPM_BUILD_ROOT/%{_bindir}/dockerd

%files
/%{_bindir}/dockerd

%post
if ! getent group docker > /dev/null; then
    groupadd --system docker
fi

%changelog
