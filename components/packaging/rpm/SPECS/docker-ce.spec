%global debug_package %{nil}


Name: docker-ce
Version: %{_version}
Release: %{_release}%{?dist}
Epoch: 3
Source0: engine.tgz
Source1: docker.service
Source2: docker.socket
Summary: The open-source application container engine
Group: Tools/Docker
License: ASL 2.0
URL: https://www.docker.com
Vendor: Docker
Packager: Docker <support@docker.com>

Requires: docker-ce-cli
Requires: container-selinux >= 2.9
Requires: libseccomp >= 2.3
Requires: systemd-units
Requires: iptables
Requires: libcgroup
Requires: containerd.io
Requires: tar
Requires: xz

# Resolves: rhbz#1165615
Requires: device-mapper-libs >= 1.02.90-1

BuildRequires: bash
BuildRequires: btrfs-progs-devel
BuildRequires: ca-certificates
BuildRequires: cmake
BuildRequires: device-mapper-devel
BuildRequires: gcc
BuildRequires: git
BuildRequires: glibc-static
BuildRequires: libseccomp-devel
BuildRequires: libselinux-devel
BuildRequires: libtool
BuildRequires: libtool-ltdl-devel
BuildRequires: make
BuildRequires: pkgconfig
BuildRequires: pkgconfig(systemd)
BuildRequires: selinux-policy-devel
BuildRequires: systemd-devel
BuildRequires: tar
BuildRequires: which

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
Docker is is a product for you to build, ship and run any application as a
lightweight container.

Docker containers are both hardware-agnostic and platform-agnostic. This means
they can run anywhere, from your laptop to the largest cloud compute instance and
everything in between - and they don't require you to use a particular
language, framework or packaging system. That makes them great building blocks
for deploying and scaling web apps, databases, and backend services without
depending on a particular stack or provider.

%prep
%setup -q -c -n src -a 0

%build
export DOCKER_GITCOMMIT=%{_gitcommit}
mkdir -p /go/src/github.com/docker
ln -s /root/rpmbuild/BUILD/src/engine /go/src/github.com/docker/docker

pushd engine
for component in tini "proxy dynamic";do
    TMP_GOPATH="/go" hack/dockerfile/install/install.sh $component
done
VERSION=%{_origversion} PRODUCT=docker hack/make.sh dynbinary
popd

%check
engine/bundles/dynbinary-daemon/dockerd -v

%install
# install daemon binary
install -D -p -m 0755 $(readlink -f engine/bundles/dynbinary-daemon/dockerd) $RPM_BUILD_ROOT/%{_bindir}/dockerd

# install proxy
install -D -p -m 0755 /usr/local/bin/docker-proxy $RPM_BUILD_ROOT/%{_bindir}/docker-proxy

# install tini
install -D -p -m 755 /usr/local/bin/docker-init $RPM_BUILD_ROOT/%{_bindir}/docker-init

# install systemd scripts
install -D -m 0644 %{_topdir}/SOURCES/docker.service $RPM_BUILD_ROOT/%{_unitdir}/docker.service
install -D -m 0644 %{_topdir}/SOURCES/docker.socket $RPM_BUILD_ROOT/%{_unitdir}/docker.socket

# install json for docker engine activate / upgrade
install -D -m 0644 %{_topdir}/SOURCES/distribution_based_engine.json $RPM_BUILD_ROOT/var/lib/docker-engine/distribution_based_engine-ce.json

%files
/%{_bindir}/dockerd
/%{_bindir}/docker-proxy
/%{_bindir}/docker-init
/%{_unitdir}/docker.service
/%{_unitdir}/docker.socket
/var/lib/docker-engine/distribution_based_engine-ce.json

%pre
if [ $1 -gt 0 ] ; then
    # package upgrade scenario, before new files are installed

    # clear any old state
    rm -f %{_localstatedir}/lib/rpm-state/docker-is-active > /dev/null 2>&1 || :

    # check if docker service is running
    if systemctl is-active docker > /dev/null 2>&1; then
        systemctl stop docker > /dev/null 2>&1 || :
        touch %{_localstatedir}/lib/rpm-state/docker-is-active > /dev/null 2>&1 || :
    fi
fi

%post
%systemd_post docker
if ! getent group docker > /dev/null; then
    groupadd --system docker
fi


%preun
%systemd_preun docker

%postun
%systemd_postun_with_restart docker

%posttrans
if [ $1 -ge 0 ] ; then
    # package upgrade scenario, after new files are installed

    # check if docker was running before upgrade
    if [ -f %{_localstatedir}/lib/rpm-state/docker-is-active ]; then
        systemctl start docker > /dev/null 2>&1 || :
        rm -f %{_localstatedir}/lib/rpm-state/docker-is-active > /dev/null 2>&1 || :
    fi
fi

%changelog
