ARG GO_IMAGE
ARG BUILD_IMAGE=resin/rpi-raspbian:stretch
FROM ${GO_IMAGE} as golang

FROM ${BUILD_IMAGE}

RUN apt-get update && apt-get install -y curl devscripts equivs git

ARG GO_VERSION
ENV GOPATH /go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin
ENV DOCKER_BUILDTAGS apparmor pkcs11 seccomp selinux
ENV RUNC_BUILDTAGS apparmor seccomp selinux

ARG COMMON_FILES
COPY ${COMMON_FILES} /root/build-deb/debian
RUN mk-build-deps -t "apt-get -o Debug::pkgProblemResolver=yes --no-install-recommends -y" -i /root/build-deb/debian/control

COPY sources/ /sources

ENV DISTRO raspbian
ENV SUITE stretch

COPY --from=golang /usr/local/go /usr/local/go

WORKDIR /root/build-deb
COPY build-deb /root/build-deb/build-deb

ENTRYPOINT ["/root/build-deb/build-deb"]
