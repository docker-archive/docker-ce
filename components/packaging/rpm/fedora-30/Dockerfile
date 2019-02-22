ARG GO_IMAGE
ARG ENGINE_IMAGE
ARG BUILD_IMAGE=fedora:30
FROM ${GO_IMAGE} as golang
FROM ${ENGINE_IMAGE} as engine

FROM ${BUILD_IMAGE}
ENV DISTRO fedora
ENV SUITE 30
ENV GOPATH /go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin
ENV AUTO_GOPATH 1
ENV DOCKER_BUILDTAGS pkcs11 seccomp selinux
ENV RUNC_BUILDTAGS seccomp selinux
RUN dnf install -y rpm-build rpmlint dnf-plugins-core
COPY SPECS /root/rpmbuild/SPECS
RUN dnf builddep -y /root/rpmbuild/SPECS/*.spec
COPY --from=golang /usr/local/go /usr/local/go/
COPY --from=engine /bin/dockerd /sources/
COPY --from=engine /bin/docker-proxy /sources/
COPY --from=engine /bin/docker-init /sources/
WORKDIR /root/rpmbuild
ENTRYPOINT ["/bin/rpmbuild"]
