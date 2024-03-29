
# The binaries to build (just the basenames).
BINARY_NAMES := seabattle

# Where to push the docker image.
REGISTRY ?= docker.pkg.github.com/teratron/seabattle

# This version-strategy uses git tags to set the version string
#VERSION ?= $(shell git describe --tags --always --dirty)

# This version-strategy uses a manual value to set the version string
VERSION ?= 0.0.1

# Directories which hold app source (not vendored)
SRC_DIRS := cmd pkg

ALL_PLATFORMS := linux/amd64 linux/arm linux/arm64 linux/ppc64le linux/s390x windows/amd64

# Used internally. Users should pass GOOS and/or GOARCH.
OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))
TAG := $(VERSION)__$(OS)_$(ARCH)

BASE_IMAGE ?= gcr.io/distroless/static
BUILD_IMAGE ?= golang:1.17-alpine

BIN_EXTENSION :=
	ifeq ($(OS), windows)
		BIN_EXTENSION := .exe
	endif

deps: ## install deps
	go get -u gopkg.in/yaml.v2
	go get -u github.com/go-sql-driver/mysql
	go mod tidy
	go mod vendor

all: build ## builds binaries for one platform ($OS/$ARCH)

# For the following OS/ARCH expansions, we transform OS/ARCH into OS_ARCH
# because make pattern rules don't match with embedded '/' characters.
build-%:
	@$(MAKE) build                        \
	    --no-print-directory              \
	    GOOS=$(firstword $(subst _, ,$*)) \
	    GOARCH=$(lastword $(subst _, ,$*))

container-%:
	@$(MAKE) container                    \
	    --no-print-directory              \
	    GOOS=$(firstword $(subst _, ,$*)) \
	    GOARCH=$(lastword $(subst _, ,$*))

push-%:
	@$(MAKE) push                         \
	    --no-print-directory              \
	    GOOS=$(firstword $(subst _, ,$*)) \
	    GOARCH=$(lastword $(subst _, ,$*))

all-build: $(addprefix build-, $(subst /,_, $(ALL_PLATFORMS))) ## builds binaries for all platforms

all-container: $(addprefix container-, $(subst /,_, $(ALL_PLATFORMS))) ## builds containers for all platforms

all-push: $(addprefix push-, $(subst /,_, $(ALL_PLATFORMS))) ## pushes containers for all platforms to the defined registry

# The following structure defeats Go's (intentional) behavior to always touch
# result files, even if they have not changed. This will still run `go` but
# will not trigger further work if nothing has actually changed.
OUTBINS = $(foreach bin,$(BINARY_NAMES),bin/$(OS)_$(ARCH)/$(bin)$(BIN_EXTENSION))

build: $(OUTBINS)

# Directories that we need created to build/test.
BUILD_DIRS := bin/$(OS)_$(ARCH)     \
              .go/bin/$(OS)_$(ARCH) \
              .go/cache

# Each outbin target is just a facade for the respective stampfile target.
# This `eval` establishes the dependencies for each.
$(foreach outbin, $(OUTBINS), $(eval $(outbin): .go/$(outbin).stamp))

# This is the target definition for all outbins.
$(OUTBINS):
	@true

# Each stampfile target can reference an $(OUTBIN) variable.
$(foreach outbin,$(OUTBINS),$(eval $(strip  \
    .go/$(outbin).stamp: OUTBIN = $(outbin) \
)))

# This is the target definition for all stampfiles.
# This will build the binary under ./.go and update the real binary iff needed.
STAMPS = $(foreach outbin,$(OUTBINS),.go/$(outbin).stamp)
.PHONY: $(STAMPS)
$(STAMPS): go-build
	@echo "binary: $(OUTBIN)"
	@if ! cmp -s .go/$(OUTBIN) $(OUTBIN); then  \
	    mv .go/$(OUTBIN) $(OUTBIN);             \
	    date >$@;                               \
	fi

go-build: $(BUILD_DIRS) ## this runs the actual `go build` which updates all binaries
	@echo
	@echo "building for $(OS)/$(ARCH)"
	@docker run                                                 \
	    -i                                                      \
	    --rm                                                    \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $$(pwd):/src                                         \
	    -w /src                                                 \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin                \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin/$(OS)_$(ARCH)  \
	    -v $$(pwd)/.go/cache:/.cache                            \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    $(BUILD_IMAGE)                                          \
	    /bin/sh -c "                                            \
	        ARCH=$(ARCH)                                        \
	        OS=$(OS)                                            \
	        VERSION=$(VERSION)                                  \
	        ./build/build.sh                                    \
	    "

# Example: make shell CMD="-c 'date > datefile'"
shell: $(BUILD_DIRS) ## launches a shell in the containerized build environment
	@echo "launching a shell in the containerized build environment"
	@docker run                                                 \
	    -ti                                                     \
	    --rm                                                    \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $$(pwd):/src                                         \
	    -w /src                                                 \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin                \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin/$(OS)_$(ARCH)  \
	    -v $$(pwd)/.go/cache:/.cache                            \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    $(BUILD_IMAGE)                                          \
	    /bin/sh $(CMD)

CONTAINER_DOTFILES = $(foreach bin,$(BINARY_NAMES),.container-$(subst /,_,$(REGISTRY)/$(bin))-$(TAG))

container containers: $(CONTAINER_DOTFILES) ## build(s) container(s) for one platform ($OS/$ARCH)
	@for bin in $(BINARY_NAMES);                       \
		do echo "container: $(REGISTRY)/$$bin:$(TAG)"; \
	done

# Each container-dotfile target can reference a $(BIN) variable.
# This is done in 2 steps to enable target-specific variables.
$(foreach bin,$(BINARY_NAMES),$(eval $(strip                                 \
    .container-$(subst /,_,$(REGISTRY)/$(bin))-$(TAG): BIN = $(bin)  \
)))

$(foreach bin,$(BINARY_NAMES),$(eval                                                                                   \
    .container-$(subst /,_,$(REGISTRY)/$(bin))-$(TAG): bin/$(OS)_$(ARCH)/$(bin)$(BIN_EXTENSION) in.Dockerfile  \
))

# This is the target definition for all container-dotfiles.
# These are used to track build state in hidden files.
$(CONTAINER_DOTFILES):
	@sed                                          \
	    -e 's|{ARG_BIN}|$(BIN)$(BIN_EXTENSION)|g' \
	    -e 's|{ARG_ARCH}|$(ARCH)|g'               \
	    -e 's|{ARG_OS}|$(OS)|g'                   \
	    -e 's|{ARG_FROM}|$(BASE_IMAGE)|g'         \
	    in.Dockerfile > .dockerfile-$(BIN)-$(OS)_$(ARCH)
	@docker build -t $(REGISTRY)/$(BIN):$(TAG) -t $(REGISTRY)/$(BIN):latest -f .dockerfile-$(BIN)-$(OS)_$(ARCH) .
	@docker images -q $(REGISTRY)/$(BIN):$(TAG) > $@
	@echo

push: $(CONTAINER_DOTFILES) ## pushes the container for one platform ($OS/$ARCH) to the defined registry
	@for bin in $(BINARY_NAMES); do                    \
	    docker push $(REGISTRY)/$$bin:$(TAG);  \
	done

manifest-list: all-push ## builds a manifest list of containers for all platforms
	@for bin in $(BINARY_NAMES); do                                   \
	    platforms=$$(echo $(ALL_PLATFORMS) | sed 's/ /,/g');  \
	    manifest-tool                                         \
	        --username=oauth2accesstoken                      \
	        --password=$$(gcloud auth print-access-token)     \
	        push from-args                                    \
	        --platforms "$$platforms"                         \
	        --template $(REGISTRY)/$$bin:$(VERSION)__OS_ARCH  \
	        --target $(REGISTRY)/$$bin:$(VERSION)

version: ## outputs the version string
	@echo $(VERSION)

test: $(BUILD_DIRS) ## runs tests, as defined in ./build/test.sh
	@docker run                                                 \
	    -i                                                      \
	    --rm                                                    \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $$(pwd):/src                                         \
	    -w /src                                                 \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin                \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin/$(OS)_$(ARCH)  \
	    -v $$(pwd)/.go/cache:/.cache                            \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    $(BUILD_IMAGE)                                          \
	    /bin/sh -c "                                            \
	        ARCH=$(ARCH)                                        \
	        OS=$(OS)                                            \
	        VERSION=$(VERSION)                                  \
	        ./build/test.sh $(SRC_DIRS)                         \
	    "

$(BUILD_DIRS):
	@mkdir -p $@

clean: container-clean bin-clean ## removes built binaries and temporary files

container-clean: ## removes temporary files
	rm -rf .container-* .dockerfile-*

bin-clean: ## removes built binaries
	rm -rf .go bin

fmt: ## go fmt
	$(call print-target)
	go fmt ./...

.PHONY: help
help:
	@echo "VARIABLES:"
	@echo "  BINS = $(BINARY_NAMES)"
	@echo "  OS = $(OS)"
	@echo "  ARCH = $(ARCH)"
	@echo "  REGISTRY = $(REGISTRY)"
	@echo
	@echo "TARGETS:"
	@awk '                                             \
		BEGIN {FS = ":.*?## "}                         \
		/^[a-zA-Z_-]+:.*?## /                          \
		{printf "\033[36m%-16s\033[0m %s\n", $$1, $$2} \
	'                                                  \
	$(MAKEFILE_LIST)

.DEFAULT_GOAL := help

# help: @HELP prints this message
#help:
#	@echo "VARIABLES:"
#	@echo "  BINS = $(BINARY_NAMES)"
#	@echo "  OS = $(OS)"
#	@echo "  ARCH = $(ARCH)"
#	@echo "  REGISTRY = $(REGISTRY)"
#	@echo
#	@echo "TARGETS:"
#	@grep -E '^.*: *@HELP' $(MAKEFILE_LIST)       \
#		| awk '                                   \
#			BEGIN {FS = ": *@HELP"};              \
#			{ printf "  %-20s %s\n", $$1, $$2 };  \
#		'

#.PHONY: help2
#help2:
#	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST)   \
#	| sort                                                \
#	| awk '                                               \
#		BEGIN {FS = ":.*?## "};                           \
#		{printf "\033[36m%-30s\033[0m %s\n", $$1, $$2};   \
#	'
#.DEFAULT_GOAL := help2

define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef
