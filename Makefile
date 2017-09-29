CLI_DIR:=$(CURDIR)/components/cli
ENGINE_DIR:=$(CURDIR)/components/engine
PACKAGING_DIR:=$(CURDIR)/components/packaging
MOBY_COMPONENTS_SHA=ab7c118272b02d8672dc0255561d0c4015979780
MOBY_COMPONENTS_URL=https://raw.githubusercontent.com/shykes/moby-extras/$(MOBY_COMPONENTS_SHA)/cmd/moby-components
MOBY_COMPONENTS=.helpers/moby-components-$(MOBY_COMPONENTS_SHA)
VERSION=$(shell cat VERSION)

.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: test-integration-cli
test-integration-cli: $(CLI_DIR)/build/docker ## test integration of cli and engine
	$(MAKE) -C $(ENGINE_DIR) DOCKER_CLI_PATH=$< test-integration-cli

$(CLI_DIR)/build/docker:
	$(MAKE) -C $(CLI_DIR) -f docker.Makefile build

.PHONY: deb
deb: ## build deb packages
	$(MAKE) VERSION=$(VERSION) CLI_DIR=$(CLI_DIR) ENGINE_DIR=$(ENGINE_DIR) -C $(PACKAGING_DIR) deb

.PHONY: rpm
rpm: ## build rpm packages
	$(MAKE) VERSION=$(VERSION) CLI_DIR=$(CLI_DIR) ENGINE_DIR=$(ENGINE_DIR) -C $(PACKAGING_DIR) rpm

.PHONY: static
static: ## build static packages
	$(MAKE) VERSION=$(VERSION) CLI_DIR=$(CLI_DIR) ENGINE_DIR=$(ENGINE_DIR) -C $(PACKAGING_DIR) static

.PHONY: clean
clean: ## clean the build artifacts
	-$(MAKE) -C $(CLI_DIR) clean
	-$(MAKE) -C $(ENGINE_DIR) clean
	-$(MAKE) -C $(PACKAGING_DIR) clean

$(MOBY_COMPONENTS):
	mkdir -p .helpers
	curl -fsSL $(MOBY_COMPONENTS_URL) > $(MOBY_COMPONENTS)
	chmod +x $(MOBY_COMPONENTS)

.PHONY: update-components
update-components: update-components-cli update-components-engine update-components-packaging ## udpate components using moby extra tool

.PHONY: update-components-cli
update-components-cli: $(MOBY_COMPONENTS)
	$(MOBY_COMPONENTS) update cli

.PHONY: update-components-engine
update-components-engine: $(MOBY_COMPONENTS)
	$(MOBY_COMPONENTS) update engine

.PHONY: update-components-packaging
update-components-packaging: $(MOBY_COMPONENTS)
	$(MOBY_COMPONENTS) update packaging
