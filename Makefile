CLI_DIR:=$(CURDIR)/components/cli
ENGINE_DIR:=$(CURDIR)/components/engine
PACKAGING_DIR:=$(CURDIR)/components/packaging

help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test-integration-cli: $(CLI_DIR)/build/docker ## test integration of cli and engine
	$(MAKE) -C $(ENGINE_DIR) DOCKER_CLI_PATH=$< test-integration-cli

$(CLI_DIR)/build/docker:
	$(MAKE) -C $(CLI_DIR) -f docker.Makefile build

deb: ## build deb packages
	$(MAKE) CLI_DIR=$(CLI_DIR) ENGINE_DIR=$(ENGINE_DIR) -C $(PACKAGING_DIR) deb

rpm: ## build rpm packages
	$(MAKE) CLI_DIR=$(CLI_DIR) ENGINE_DIR=$(ENGINE_DIR) -C $(PACKAGING_DIR) rpm

tgz: ## build static packages
	$(MAKE) CLI_DIR=$(CLI_DIR) ENGINE_DIR=$(ENGINE_DIR) -C $(PACKAGING_DIR) tgz
