%global debug_package %{nil}

Name: docker-ce-cli
Version: %{_version}
Release: %{_release}%{?dist}
Epoch: %{getenv:EPOCH}
Summary: The open-source application container engine
Group: Tools/Docker
License: ASL 2.0
Source0: cli.tgz
URL: https://www.docker.com
Vendor: Docker
Packager: Docker <support@docker.com>

# required packages on install
Requires: /bin/sh
Requires: containerd

# conflicting packages
Conflicts: docker
Conflicts: docker-io
Conflicts: docker-engine-cs
Conflicts: docker-ee
Conflicts: docker-ee-cli

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
mkdir -p /go/src/github.com/docker
rm -f /go/src/github.com/docker/cli
ln -s /root/rpmbuild/BUILD/src/cli /go/src/github.com/docker/cli
pushd /go/src/github.com/docker/cli
DISABLE_WARN_OUTSIDE_CONTAINER=1 make VERSION=%{_origversion} GITCOMMIT=%{_gitcommit} dynbinary manpages # cli
popd

# %check
# cli/build/docker -v

%install
# install binary
install -d $RPM_BUILD_ROOT/%{_bindir}
install -p -m 755 cli/build/docker $RPM_BUILD_ROOT/%{_bindir}/docker

# add bash, zsh, and fish completions
install -d $RPM_BUILD_ROOT/usr/share/bash-completion/completions
install -d $RPM_BUILD_ROOT/usr/share/zsh/vendor-completions
install -d $RPM_BUILD_ROOT/usr/share/fish/vendor_completions.d
install -p -m 644 cli/contrib/completion/bash/docker $RPM_BUILD_ROOT/usr/share/bash-completion/completions/docker
install -p -m 644 cli/contrib/completion/zsh/_docker $RPM_BUILD_ROOT/usr/share/zsh/vendor-completions/_docker
install -p -m 644 cli/contrib/completion/fish/docker.fish $RPM_BUILD_ROOT/usr/share/fish/vendor_completions.d/docker.fish

# install manpages
install -d %{buildroot}%{_mandir}/man1
install -p -m 644 cli/man/man1/*.1 $RPM_BUILD_ROOT/%{_mandir}/man1
install -d %{buildroot}%{_mandir}/man5
install -p -m 644 cli/man/man5/*.5 $RPM_BUILD_ROOT/%{_mandir}/man5
install -d %{buildroot}%{_mandir}/man8
install -p -m 644 cli/man/man8/*.8 $RPM_BUILD_ROOT/%{_mandir}/man8

mkdir -p build-docs
for cli_file in LICENSE MAINTAINERS NOTICE README.md; do
    cp "cli/$cli_file" "build-docs/$cli_file"
done

# list files owned by the package here
%files
%doc build-docs/LICENSE build-docs/MAINTAINERS build-docs/NOTICE build-docs/README.md
/%{_bindir}/docker
/usr/share/bash-completion/completions/docker
/usr/share/zsh/vendor-completions/_docker
/usr/share/fish/vendor_completions.d/docker.fish
%doc
/%{_mandir}/man1/*
/%{_mandir}/man5/*
/%{_mandir}/man8/*


%post
if ! getent group docker > /dev/null; then
    groupadd --system docker
fi

%changelog
