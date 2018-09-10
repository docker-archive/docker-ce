ARG GO_IMAGE
ARG CONTAINERD_SHIM_PROCESS_IMAGE
FROM ${GO_IMAGE} as golang
FROM ${CONTAINERD_SHIM_PROCESS_IMAGE} as shim-process

FROM ubuntu:bionic

RUN apt-get update && apt-get install -y curl devscripts equivs git

ARG GO_VERSION
ENV GOPATH /go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin
ENV DOCKER_BUILDTAGS apparmor pkcs11 seccomp selinux
ENV RUNC_BUILDTAGS apparmor seccomp selinux

ARG COMMON_FILES
COPY ${COMMON_FILES} /root/build-deb/debian
RUN mk-build-deps -t "apt-get -o Debug::pkgProblemResolver=yes --no-install-recommends -y" -i /root/build-deb/debian/control

# Copy our sources and untar them
COPY sources/ /sources
RUN mkdir -p /go/src/github.com/docker/ && tar -xzf /sources/cli.tgz -C /go/src/github.com/docker/
RUN mkdir -p /go/src/github.com/crosbymichael && tar -xzf /sources/containerd-proxy.tgz -C /go/src/github.com/crosbymichael

RUN ln -snf /go/src/github.com/docker/cli /root/build-deb/cli

ENV DISTRO ubuntu
ENV SUITE bionic

COPY --from=golang /usr/local/go /usr/local/go
COPY --from=shim-process /bin/containerd-shim-process-v1 /containerd-shim-process-v1

WORKDIR /root/build-deb
COPY build-deb /root/build-deb/build-deb

ENTRYPOINT ["/root/build-deb/build-deb"]
