all: clean common replication podmon migration

########################################################################
##                             GOLANG                                 ##
########################################################################

# If GOPATH isn't defined then set its default location.
ifeq (,$(strip $(GOPATH)))
GOPATH := $(HOME)/go
else
# If GOPATH is already set then update GOPATH to be its own
# first element.
GOPATH := $(word 1,$(subst :, ,$(GOPATH)))
endif
export GOPATH

########################################################################
##                             PROTOC                                 ##
########################################################################

# Only set PROTOC_VER if it has an empty value.
ifeq (,$(strip $(PROTOC_VER)))
PROTOC_VER := 26.1
endif

PROTOC_OS := $(shell uname -s)
ifeq (Darwin,$(PROTOC_OS))
PROTOC_OS := osx
endif
ifeq (Linux,$(PROTOC_OS))
PROTOC_OS := linux
endif

PROTOC_ARCH := $(shell uname -m)
ifeq (i386,$(PROTOC_ARCH))
PROTOC_ARCH := x86_32
endif
ifeq (arm64osx,$(PROTOC_ARCH)$(PROTOC_OS))
PROTOC_ARCH := aarch_64
endif
ifeq (x86_64,$(PROTOC_ARCH))
PROTOC_ARCH := x86_64
endif

PROTOC := ./protoc
PROTO_INCLUDE=.protoc/include

PROTOC_ZIP := protoc-$(PROTOC_VER)-$(PROTOC_OS)-$(PROTOC_ARCH).zip
ifeq (,$(strip $(PROTOC_OS)))
	PROTOC_ZIP := protoc-$(PROTOC_VER)-win64.zip
endif
PROTOC_URL := https://github.com/google/protobuf/releases/download/v$(PROTOC_VER)/$(PROTOC_ZIP)
PROTOC_TMP_DIR := .protoc
PROTOC_TMP_BIN := $(PROTOC_TMP_DIR)/bin/protoc

$(PROTOC):
	-mkdir -p "$(PROTOC_TMP_DIR)" && \
	  curl -L $(PROTOC_URL) -o "$(PROTOC_TMP_DIR)/$(PROTOC_ZIP)" && \
	  unzip "$(PROTOC_TMP_DIR)/$(PROTOC_ZIP)" -d "$(PROTOC_TMP_DIR)" && \
	  chmod 0755 "$(PROTOC_TMP_BIN)" && \
	  cp -f "$(PROTOC_TMP_BIN)" "$@"
	stat "$@" > /dev/null 2>&1

########################################################################
##                          PROTOC-GEN-GO                             ##
########################################################################

# This is the recipe for getting and installing the go plug-in
# for protoc
PROTOC_GEN_GO := protoc-gen-go
$(PROTOC_GEN_GO):
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

########################################################################
##                              PATH                                  ##
########################################################################

# Update PATH with the current directory. This enables the protoc
# binary to discover the protoc-gen-go binary, built inside this
# directory.
export PATH := $(shell pwd):$(PATH)

.PHONY: replication
replication: $(PROTOC) $(PROTOC_GEN_GO)
	$(PWD)/$(PROTOC) -I $(PROTO_INCLUDE) -I ./common --go_out="$@" --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:"$@" --go-grpc_opt=paths=source_relative  --proto_path="$@" ./"$@"/"$@".proto
	(cd "$@"; go mod tidy; go build "$@".pb.go)

.PHONY: migration
migration: $(PROTOC) $(PROTOC_GEN_GO)
	$(PWD)/$(PROTOC) -I $(PROTO_INCLUDE) -I ./common --go_out="$@" --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:"$@" --go-grpc_opt=paths=source_relative  --proto_path="$@" ./"$@"/"$@".proto
	(cd "$@"; go mod tidy; go build "$@".pb.go)

.PHONY: podmon
podmon: $(PROTOC) $(PROTOC_GEN_GO)
	$(PWD)/$(PROTOC) -I $(PROTO_INCLUDE) --go_out="$@" --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:"$@" --go-grpc_opt=paths=source_relative  --proto_path="$@" ./"$@"/"$@".proto
	(cd "$@"; go mod tidy; go build "$@".pb.go)

.PHONY: common
common: $(PROTOC) $(PROTOC_GEN_GO)
	$(PWD)/$(PROTOC) -I $(PROTO_INCLUDE) --go_out="$@" --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:"$@" --go-grpc_opt=paths=source_relative  --proto_path="$@" ./"$@"/"$@".proto
	(cd "$@"; go mod tidy; go build "$@".pb.go)

clean:
	rm -rf ./replication/replication.pb.go ./replication/replication_grpc.pb.go ./podmon/podmon.pb.go ./podmon/podmon_grpc.pb.go ./migration/migration.pb.go ./migration/migration_grpc.pb.go ./common/common.pb.go

clobber: clean
	rm -rf "$(PROTOC)" "$(PROTOC_GEN_GO)" "$(PROTOC_TMP_DIR)"

.PHONY: actions action-help
actions: ## Run all GitHub Action checks that run on a pull request creation
	@echo "Running all GitHub Action checks for pull request events..."
	@act -l | grep -v ^Stage | grep pull_request | grep -v image_security_scan | awk '{print $$2}' | while read WF; do \
		echo "Running workflow: $${WF}"; \
		act pull_request --no-cache-server --platform ubuntu-latest=ghcr.io/catthehacker/ubuntu:act-latest --job "$${WF}"; \
	done

action-help: ## Echo instructions to run one specific workflow locally
	@echo "GitHub Workflows can be run locally with the following command:"
	@echo "act pull_request --no-cache-server --platform ubuntu-latest=ghcr.io/catthehacker/ubuntu:act-latest --job <jobid>"
	@echo ""
	@echo "Where '<jobid>' is a Job ID returned by the command:"
	@echo "act -l"
	@echo ""
	@echo "NOTE: if act is not installed, it can be downloaded from https://github.com/nektos/act"