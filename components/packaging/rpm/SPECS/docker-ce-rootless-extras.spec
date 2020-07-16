%global debug_package %{nil}

Name: docker-ce-rootless-extras
Version: %{_version}
Release: %{_release}%{?dist}
Epoch: 0
Source0: engine.tgz
Summary: Rootless support for Docker
Group: Tools/Docker
License: ASL 2.0
URL: https://docs.docker.com/engine/security/rootless/
Vendor: Docker
Packager: Docker <support@docker.com>

Requires: docker-ce
# slirp4netns >= 0.4 is available in the all supported versions of CentOS and Fedora.
Requires: slirp4netns >= 0.4
# fuse-overlayfs >= 0.7 is available in the all supported versions of CentOS and Fedora.
Requires: fuse-overlayfs >= 0.7

BuildRequires: bash

# conflicting packages
Conflicts: rootlesskit

%description
Rootless support for Docker.
Use dockerd-rootless.sh to run the daemon.
Use dockerd-rootless-setuptool.sh to setup systemd for dockerd-rootless.sh .
This package contains RootlessKit, but does not contain VPNKit.
Either VPNKit or slirp4netns (>= 0.4.0) needs to be installed separately.

%prep
%setup -q -c -n src -a 0

%build

export DOCKER_GITCOMMIT=%{_gitcommit_engine}
mkdir -p /go/src/github.com/docker
ln -s ${RPM_BUILD_DIR}/src/engine /go/src/github.com/docker/docker
TMP_GOPATH="/go" ${RPM_BUILD_DIR}/src/engine/hack/dockerfile/install/install.sh rootlesskit dynamic

%check
/usr/local/bin/rootlesskit -v

%install
install -D -p -m 0755 engine/contrib/dockerd-rootless.sh ${RPM_BUILD_ROOT}%{_bindir}/dockerd-rootless.sh
install -D -p -m 0755 engine/contrib/dockerd-rootless-setuptool.sh ${RPM_BUILD_ROOT}%{_bindir}/dockerd-rootless-setuptool.sh
install -D -p -m 0755 /usr/local/bin/rootlesskit ${RPM_BUILD_ROOT}%{_bindir}/rootlesskit
install -D -p -m 0755 /usr/local/bin/rootlesskit-docker-proxy ${RPM_BUILD_ROOT}%{_bindir}/rootlesskit-docker-proxy

%files
%{_bindir}/dockerd-rootless.sh
%{_bindir}/dockerd-rootless-setuptool.sh
%{_bindir}/rootlesskit
%{_bindir}/rootlesskit-docker-proxy

%post

%preun

%postun

%changelog
