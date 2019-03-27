ARG GO_IMAGE
ARG BUILD_IMAGE=fedora:29
FROM ${GO_IMAGE} as golang

FROM ${BUILD_IMAGE}
ENV DISTRO fedora
ENV SUITE 29
ENV GOPATH /go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin
ENV AUTO_GOPATH 1
ENV DOCKER_BUILDTAGS seccomp selinux
ENV RUNC_BUILDTAGS seccomp selinux
RUN dnf install -y rpm-build rpmlint dnf-plugins-core
COPY SPECS /root/rpmbuild/SPECS
RUN dnf builddep -y /root/rpmbuild/SPECS/*.spec
COPY --from=golang /usr/local/go /usr/local/go
WORKDIR /root/rpmbuild
ENTRYPOINT ["/bin/rpmbuild"]
