all: clean replication podmon volumeGroupSnapshot

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
PROTOC_VER := 3.15.0
endif

PROTOC_OS := $(shell uname -s)
ifeq (Darwin,$(PROTOC_OS))
PROTOC_OS := osx
endif

PROTOC_ARCH := $(shell uname -m)
ifeq (i386,$(PROTOC_ARCH))
PROTOC_ARCH := x86_32
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
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

########################################################################
##                              PATH                                  ##
########################################################################

# Update PATH with the current directory. This enables the protoc
# binary to discover the protoc-gen-go binary, built inside this
# directory.
export PATH := $(shell pwd):$(PATH)

.PHONY: replication
replication: $(PROTOC) $(PROTOC_GEN_GO)
	$(PWD)/$(PROTOC) -I $(PROTO_INCLUDE) --go_out="$@" --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:"$@" --go-grpc_opt=paths=source_relative  --proto_path="$@" ./"$@"/"$@".proto
	(cd "$@"; go mod tidy; go build "$@".pb.go)

.PHONY: podmon
podmon: $(PROTOC) $(PROTOC_GEN_GO)
	$(PWD)/$(PROTOC) -I $(PROTO_INCLUDE) --go_out="$@" --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:"$@" --go-grpc_opt=paths=source_relative  --proto_path="$@" ./"$@"/"$@".proto
	(cd "$@"; go mod tidy; go build "$@".pb.go)

.PHONY: volumeGroupSnapshot
volumeGroupSnapshot: $(PROTOC) $(PROTOC_GEN_GO)
	$(PWD)/$(PROTOC) -I $(PROTO_INCLUDE) --go_out="$@" --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:"$@" --go-grpc_opt=paths=source_relative  --proto_path="$@" ./"$@"/"$@".proto
	(cd "$@"; go mod tidy; go build "$@".pb.go)

clean:
	rm -rf ./replication/replication.pb.go ./replication/replication_grpc.pb.go ./podmon/podmon.pb.go ./podmon/podmon_grpc.pb.go ./volumeGroupSnapshot/volumeGroupSnapshot.pb.go ./volumeGroupSnapshot/volumeGroupSnapshot_grpc.pb.go

clobber: clean
	rm -rf "$(PROTOC)" "$(PROTOC_GEN_GO)" "$(PROTOC_TMP_DIR)"


