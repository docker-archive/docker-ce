%global debug_package %{nil}


Name: docker-ce
Version: %{_version}
Release: %{_release}%{?dist}
Epoch: 3
Source0: docker.service
Source1: docker.socket
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
Requires: containerd.io

BuildRequires: which
BuildRequires: make
BuildRequires: gcc
BuildRequires: pkgconfig(systemd)

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

%build

%install
# Install containerd-proxy as dockerd
install -D -m 0755 /sources/dockerd $RPM_BUILD_ROOT/%{_bindir}/dockerd-ce
install -D -m 0755 /sources/docker-proxy $RPM_BUILD_ROOT/%{_bindir}/docker-proxy
install -D -m 0755 /sources/docker-init $RPM_BUILD_ROOT/%{_bindir}/docker-init
install -D -m 0644 %{_topdir}/SOURCES/docker.service $RPM_BUILD_ROOT/%{_unitdir}/docker.service
install -D -m 0644 %{_topdir}/SOURCES/docker.socket $RPM_BUILD_ROOT/%{_unitdir}/docker.socket
install -D -m 0644 %{_topdir}/SOURCES/distribution_based_engine.json $RPM_BUILD_ROOT/var/lib/docker-engine/distribution_based_engine-ce.json

%files
/%{_bindir}/dockerd-ce
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
dbefile=/var/lib/docker-engine/distribution_based_engine.json
URL=https://docs.docker.com/releasenote
if [ -f "${dbefile}" ] && sed -e 's/.*"platform"[ \t]*:[ \t]*"\([^"]*\)".*/\1/g' "${dbefile}"| grep -v -i community > /dev/null; then
    echo
    echo
    echo
    echo "Warning: Your engine has been activated to Docker Engine - Enterprise but you are still using Community packages"
    echo "You can use the 'docker engine update' command to update your system, or switch to using the Enterprise packages."
    echo "See $URL for more details."
    echo
    echo
    echo
else
    rm -f %{_bindir}/dockerd
    update-alternatives --install %{_bindir}/dockerd dockerd %{_bindir}/dockerd-ce 1 \
        --slave "${dbefile}" distribution_based_engine.json /var/lib/docker-engine/distribution_based_engine-ce.json
fi


%preun
%systemd_preun docker
update-alternatives --remove dockerd %{_bindir}/dockerd || true

%postun
%systemd_postun_with_restart docker

%posttrans
if [ $1 -ge 0 ] ; then
    dbefile=/var/lib/docker-engine/distribution_based_engine.json
    URL=https://docs.docker.com/releasenote
    if [ -f "${dbefile}" ] && sed -e 's/.*"platform"[ \t]*:[ \t]*"\([^"]*\)".*/\1/g' "${dbefile}"| grep -v -i community > /dev/null; then
        echo
        echo
        echo
        echo "Warning: Your engine has been activated to Docker Engine - Enterprise but you are still using Community packages"
        echo "You can use the 'docker engine update' command to update your system, or switch to using the Enterprise packages."
        echo "See $URL for more details."
        echo
        echo
        echo
    else
        rm -f %{_bindir}/dockerd
        update-alternatives --install %{_bindir}/dockerd dockerd %{_bindir}/dockerd-ce 1 \
            --slave "${dbefile}" distribution_based_engine.json /var/lib/docker-engine/distribution_based_engine-ce.json
    fi
    # package upgrade scenario, after new files are installed

    # check if docker was running before upgrade
    if [ -f %{_localstatedir}/lib/rpm-state/docker-is-active ]; then
        systemctl start docker > /dev/null 2>&1 || :
        rm -f %{_localstatedir}/lib/rpm-state/docker-is-active > /dev/null 2>&1 || :
    fi
fi

%changelog
