ARG GO_IMAGE
ARG CONTAINERD_SHIM_PROCESS_IMAGE
ARG BUILD_IMAGE=centos:7
FROM ${GO_IMAGE} as golang
FROM ${CONTAINERD_SHIM_PROCESS_IMAGE} as shim-process

FROM ${BUILD_IMAGE}
ENV DISTRO centos
ENV SUITE 7
ENV GOPATH=/go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin
ENV AUTO_GOPATH 1
ENV DOCKER_BUILDTAGS pkcs11 seccomp selinux
ENV RUNC_BUILDTAGS seccomp selinux
RUN yum install -y rpm-build rpmlint
COPY SPECS /root/rpmbuild/SPECS
# Overwrite repo that was failing on aarch64
RUN sed -i 's/altarch/centos/g' /etc/yum.repos.d/CentOS-Sources.repo
RUN yum-builddep -y /root/rpmbuild/SPECS/*.spec
COPY --from=golang /usr/local/go /usr/local/go/
COPY --from=shim-process /bin/containerd-shim-process-v1 /containerd-shim-process-v1
WORKDIR /root/rpmbuild
ENTRYPOINT ["/bin/rpmbuild"]
