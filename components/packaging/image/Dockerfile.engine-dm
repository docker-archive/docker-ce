# Common builder
ARG GO_IMAGE
ARG BASE_IMAGE=centos:7
FROM ${GO_IMAGE} as golang

FROM ${BASE_IMAGE} as builder
ENV GOPATH=/go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin
ENV AUTO_GOPATH 1
COPY --from=golang /usr/local/go /usr/local/go
COPY hack/dockerfile/install/tini.installer /
COPY hack/dockerfile/install/proxy.installer /
RUN yum install -y \
    bash \
    ca-certificates \
    cmake \
    gcc \
    git \
    glibc-static \
    libtool \
    make
RUN grep "_COMMIT=" /*.installer  |cut -f2- -d: > /binaries-commits

# dockerd
FROM builder as dockerd-builder
RUN yum install -y \
    btrfs-progs-devel \
    device-mapper-devel \
    libseccomp-devel \
    selinux-policy-devel \
    systemd-devel
WORKDIR /go/src/github.com/docker/docker
COPY . /go/src/github.com/docker/docker
ARG VERSION
ARG GITCOMMIT
ARG BUILDTIME
ARG PLATFORM
ARG PRODUCT
ARG DEFAULT_PRODUCT_LICENSE
ENV VERSION ${VERSION}
ENV GITCOMMIT ${GITCOMMIT}
ENV BUILDTIME ${BUILDTIME}
ENV PLATFORM ${PLATFORM}
ENV PRODUCT ${PRODUCT}
ENV DEFAULT_PRODUCT_LICENSE ${DEFAULT_PRODUCT_LICENSE}
# TODO The way we set the version could easily be simplified not to depend on hack/...
RUN bash ./hack/make/.go-autogen
RUN go build -o /dockerd \
    -tags 'autogen apparmor seccomp selinux journald' \
    -i \
    -buildmode=pie \
    -a -ldflags '-w'\
    github.com/docker/docker/cmd/dockerd

# docker-proxy
# TODO if libnetwork folds into the docker tree this can be combined above
FROM builder as proxy-builder
RUN git clone https://github.com/docker/libnetwork.git /go/src/github.com/docker/libnetwork
WORKDIR /go/src/github.com/docker/libnetwork
RUN . /binaries-commits && \
    git checkout -q "$LIBNETWORK_COMMIT" && \
    go build -buildmode=pie -ldflags="-w" \
        -o /docker-proxy \
        github.com/docker/libnetwork/cmd/proxy

# docker-init - TODO move this out, last time we bumped was 2016!
FROM builder as init-builder
RUN git clone https://github.com/krallin/tini.git /tini
WORKDIR /tini
RUN . /binaries-commits && \
    git checkout -q "$TINI_COMMIT" && \
    cmake . && make tini-static && \
    cp tini-static /docker-init

# Final docker image
FROM scratch
ARG VERSION
ARG GITCOMMIT
ARG BUILDTIME
ARG PLATFORM
ARG ENGINE_IMAGE
COPY --from=dockerd-builder /dockerd /bin/
COPY --from=proxy-builder /docker-proxy /bin/
COPY --from=init-builder /docker-init /bin/

LABEL \
    org.opencontainers.image.authors="Docker Inc." \
    org.opencontainers.image.created="${BUILDTIME}" \
    org.opencontainers.image.documentation="https://docs.docker.com/" \
    org.opencontainers.image.licenses="Apache-2.0" \
    org.opencontainers.image.revision="${GITCOMMIT}" \
    org.opencontainers.image.url="https://www.docker.com/products/docker-engine" \
    org.opencontainers.image.vendor="Docker Inc." \
    org.opencontainers.image.version="${VERSION}" \
    com.docker.distribution_based_engine="{\"platform\":\"${PLATFORM}\",\"engine_image\":\"${ENGINE_IMAGE}\",\"containerd_min_version\":\"1.2.0-beta.1\",\"runtime\":\"host_install\"}"

ENTRYPOINT ["/bin/dockerd"]
