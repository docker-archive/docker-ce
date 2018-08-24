# Common things for containerd functionality

CONTAINERD_PROXY_COMMIT=afca176732d9416fac1e79fd76ce45afe4cbc41f
CONTAINERD_SHIM_PROCESS_IMAGE=docker.io/docker/containerd-shim-process:ff98a47

# If containerd is running use that socket instead
ifeq ("$(shell systemctl is-active containerd)", "active")
CONTAINERD_SOCK:=/var/run/containerd/containerd.sock
else
CONTAINERD_SOCK:=/var/run/docker/containerd/docker-containerd.sock
endif
CTR=docker run \
	--rm -i \
	-v $(CONTAINERD_SOCK):/ours/containerd.sock \
	-v $(CURDIR)/artifacts:/artifacts \
	docker:18.06.0-ce \
	docker-containerd-ctr -a /ours/containerd.sock
