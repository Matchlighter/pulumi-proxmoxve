PROJECT_NAME := proxmoxve Package
include build/common.mk

PACK             := proxmoxve
ORG              := matchlighter
PACKDIR          := sdk
PROJECT          := github.com/${ORG}/pulumi-${PACK}
NODE_MODULE_NAME := @matchlighter/pulumi-${PACK}
TF_NAME          := ${PACK}
PROVIDER_PATH    := provider
VERSION_PATH     := ${PROVIDER_PATH}/pkg/version.Version

TFGEN           := pulumi-tfgen-${PACK}
PROVIDER        := pulumi-resource-${PACK}
VERSION         := $(shell pulumictl get version)

TESTPARALLELISM := 4

WORKING_DIR     := $(shell pwd)

build:: install_plugins provider build_sdks install_sdks
only_build:: build

tfgen:: install_plugins
	(cd provider && go build -o $(WORKING_DIR)/bin/${TFGEN} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${TFGEN})
	$(WORKING_DIR)/bin/${TFGEN} schema --out provider/cmd/${PROVIDER}
	(cd provider && VERSION=$(VERSION) go generate cmd/${PROVIDER}/main.go)

provider:: tfgen install_plugins # build the provider binary
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${PROVIDER})

# Build SDKs:
build_sdks:: install_plugins provider build_nodejs build_python build_go build_dotnet

build_nodejs:: VERSION := $(shell pulumictl get version --language javascript)
build_nodejs:: install_plugins tfgen # build the node sdk
	$(WORKING_DIR)/bin/$(TFGEN) nodejs --overlays provider/overlays/nodejs --out sdk/nodejs/

	cd sdk/nodejs/ && \
		yarn install && \
		yarn run tsc && \
		cp ../../README.md package.json yarn.lock ./bin/ && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" ./bin/package.json

build_python:: PYPI_VERSION := $(shell pulumictl get version --language python)
build_python:: install_plugins tfgen # build the python sdk
	$(WORKING_DIR)/bin/$(TFGEN) python --overlays provider/overlays/python --out sdk/python/
	cd sdk/python/ && \
		cp ../../README.md . && \
		python3 setup.py clean --all 2>/dev/null && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e "s/\$${VERSION}/$(PYPI_VERSION)/g" -e "s/\$${PLUGIN_VERSION}/$(VERSION)/g" ./bin/setup.py && \
		rm ./bin/setup.py.bak && \
		cd ./bin && python3 setup.py build sdist

build_go:: install_plugins tfgen # build the go sdk
	$(WORKING_DIR)/bin/$(TFGEN) go --overlays provider/overlays/go --out sdk/go/

build_dotnet:: DOTNET_VERSION := $(shell pulumictl get version --language dotnet)
build_dotnet:: install_plugins tfgen # build the dotnet sdk
	pulumictl get version --language dotnet
	$(WORKING_DIR)/bin/$(TFGEN) dotnet --overlays provider/overlays/dotnet --out sdk/dotnet/
	cd sdk/dotnet/ && \
		echo "${DOTNET_VERSION}" >version.txt && \
		dotnet build /p:Version=${DOTNET_VERSION}

build_go:: install_plugins tfgen # build the go sdk
	$(WORKING_DIR)/bin/$(TFGEN) go --overlays provider/overlays/go --out sdk/go/

# Misc:
cleanup:: # cleans up the temporary directory
	rm -r $(WORKING_DIR)/bin
	rm -f provider/cmd/${PROVIDER}/schema.go

install_plugins::
	[ -x $(shell which pulumi) ] || curl -fsSL https://get.pulumi.com | sh

# Install SDKs:
install_sdks:: install_python_sdk install_nodejs_sdk install_dotnet_sdk

install_dotnet_sdk::
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::
	cd ${PACKDIR}/python/bin && $(PIP) install --user -e .

install_go_sdk::

install_nodejs_sdk::
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

# Testing:
test_all::
	cd examples && $(GO_TEST) .

test_fast::
	cd examples && $(GO_TEST_FAST) .

.PHONY: publish_tgz
publish_tgz:
	$(call STEP_MESSAGE)
	./scripts/publish_tgz.sh

.PHONY: publish_packages
publish_packages:
	$(call STEP_MESSAGE)
	$$(go env GOPATH)/src/github.com/pulumi/scripts/ci/publish-tfgen-package .
	# $$(go env GOPATH)/src/github.com/pulumi/scripts/ci/build-package-docs.sh ${PACK}

.PHONY: check_clean_worktree
check_clean_worktree:
	$$(go env GOPATH)/src/github.com/pulumi/scripts/ci/check-worktree-is-clean.sh

# The travis_* targets are entrypoints for CI.
.PHONY: travis_cron travis_push travis_pull_request travis_api
travis_cron: all
travis_push: only_build check_clean_worktree publish_tgz only_test publish_packages
travis_pull_request: all check_clean_worktree
travis_api: all
