# Common things for containerd functionality

CONTAINERD_PROXY_COMMIT=82ae3d13e91d062dd4853379fe018638023c8da2
CONTAINERD_SHIM_PROCESS_IMAGE=docker.io/docker/containerd-shim-process:ff98a47

# If the docker-containerd.sock is available use that, else use the default containerd.sock
ifeq (,$(wildcard /var/run/docker/containerd/docker-containerd.sock))
CONTAINERD_SOCK:=/var/run/docker/containerd/docker-containerd.sock
else
CONTAINERD_SOCK:=/var/run/containerd/containerd.sock
endif
CTR=docker run \
	--rm -i \
	-v $(CONTAINERD_SOCK):/ours/containerd.sock \
	-v $(CURDIR)/artifacts:/artifacts \
	docker:18.06.0-ce \
	docker-containerd-ctr -a /ours/containerd.sock
