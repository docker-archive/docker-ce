%global debug_package %{nil}

Name: docker-scan-plugin
Version: %{_scan_rpm_version}
Release: %{_release}%{?dist}
Epoch: 0
Source0: scan-cli-plugin.tgz
Summary: Docker Scan plugin for the Docker CLI
Group: Tools/Docker
License: ASL 2.0
URL: https://github.com/docker/scan-cli-plugin/
Vendor: Docker
Packager: Docker <support@docker.com>

Requires: docker-ce-cli

# TODO change once we support scan-plugin on other architectures
BuildArch: x86_64
BuildRequires: bash

%description
Docker Scan plugin for the Docker CLI.

%prep
%setup -q -c -n src -a 0

%build
pushd ${RPM_BUILD_DIR}/src/scan-cli-plugin
bash -c 'GOPROXY="https://proxy.golang.org" TAG_NAME="%{_scan_version}" COMMIT="%{_scan_gitcommit}" PLATFORM_BINARY=docker-scan make native-build'
popd


%check
# FIXME: --version currently doesn't work as it makes a connection to the daemon
#${RPM_BUILD_ROOT}%{_libexecdir}/docker/cli-plugins/docker-scan scan --accept-license --version
${RPM_BUILD_ROOT}%{_libexecdir}/docker/cli-plugins/docker-scan --help

%install
pushd ${RPM_BUILD_DIR}/src/scan-cli-plugin
install -D -p -m 0755 bin/docker-scan ${RPM_BUILD_ROOT}%{_libexecdir}/docker/cli-plugins/docker-scan
popd

%files
%{_libexecdir}/docker/cli-plugins/docker-scan

%post

%preun

%postun

%changelog
