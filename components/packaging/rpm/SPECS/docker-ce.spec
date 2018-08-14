%global debug_package %{nil}

Name: docker-ce
Version: %{_version}
Release: %{_release}%{?dist}
Epoch: 2
Source0: containerd-proxy.tgz
Source1: containerd-shim-process.tar
Summary: The open-source application container engine
Group: Tools/Docker
License: ASL 2.0
URL: https://www.docker.com
Vendor: Docker
Packager: Docker <support@docker.com>

Requires: docker-ce-cli
# Should be required as well by docker-ce-cli but let's just be thorough
Requires: containerd.io

BuildRequires: gcc

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
# dockerd proxy compilation
mkdir -p /go/src/github.com/crosbymichael/
ls %{_topdir}/BUILD/src
ln -s %{_topdir}/BUILD/src/containerd-proxy /go/src/github.com/crosbymichael/containerd-proxy
go build -v -o /build/dockerd github.com/crosbymichael/containerd-proxy

%install
install -D -m 0755 /build/dockerd $RPM_BUILD_ROOT/%{_bindir}/dockerd
# TODO: Use containerd-offline-installer to actually install this as ExecStartPre systemd step
install -D -m 0644 %{_topdir}/SOURCES/containerd-shim-process.tar $RPM_BUILD_ROOT/%{_sharedstatedir}/containerd/containerd-shim-process.tar

%files
/%{_bindir}/dockerd
/%{_sharedstatedir}/containerd/containerd-shim-process.tar

%post
if ! getent group docker > /dev/null; then
    groupadd --system docker
fi

%changelog
