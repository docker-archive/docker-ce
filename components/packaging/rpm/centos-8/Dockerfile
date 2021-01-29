ARG GO_IMAGE
ARG DISTRO=centos
ARG SUITE=8
ARG BUILD_IMAGE=${DISTRO}:${SUITE}

FROM ${GO_IMAGE} AS golang

FROM ${BUILD_IMAGE}
ENV GOPROXY=direct
ENV GO111MODULE=off
ENV GOPATH=/go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin
ENV AUTO_GOPATH 1
ENV DOCKER_BUILDTAGS exclude_graphdriver_btrfs seccomp selinux
ENV RUNC_BUILDTAGS seccomp selinux
ARG DISTRO
ARG SUITE
ENV DISTRO=${DISTRO}
ENV SUITE=${SUITE}

# In aarch64 (arm64) images, the altarch repo is specified as repository, but
# failing, so replace the URL.
RUN if [ -f /etc/yum.repos.d/CentOS-Linux-Sources.repo ]; then sed -i 's/altarch/centos/g' /etc/yum.repos.d/CentOS-Linux-Sources.repo; fi

RUN if [ -f /etc/yum.repos.d/CentOS-Linux-PowerTools.repo ]; then sed -i 's/enabled=0/enabled=1/g' /etc/yum.repos.d/CentOS-Linux-PowerTools.repo; fi

# RHEL8 / CentOS 8 changed behavior and no longer "rpm --import" or
# "rpmkeys --import"as part of rpm package's %post scriplet. See
# https://forums.centos.org/viewtopic.php?f=54&t=72574, and
# https://access.redhat.com/solutions/3720351
RUN rpm --import /etc/pki/rpm-gpg/RPM-GPG-KEY-centosofficial
RUN yum install -y rpm-build rpmlint yum-utils
COPY SPECS /root/rpmbuild/SPECS
RUN yum-builddep --define '_without_btrfs 1' -y /root/rpmbuild/SPECS/*.spec
COPY --from=golang /usr/local/go /usr/local/go
WORKDIR /root/rpmbuild
ENTRYPOINT ["/bin/rpmbuild"]
