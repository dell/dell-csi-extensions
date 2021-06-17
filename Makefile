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
PROTOC_VER := 3.14.0
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
PROTOC_GEN_GO_PKG := github.com/golang/protobuf/protoc-gen-go
PROTOC_GEN_GO := protoc-gen-go
$(PROTOC_GEN_GO): PROTOBUF_PKG := $(dir $(PROTOC_GEN_GO_PKG))
$(PROTOC_GEN_GO): PROTOBUF_VERSION := v1.5.2
$(PROTOC_GEN_GO):
	mkdir -p $(dir $(GOPATH)/src/$(PROTOBUF_PKG))
	test -d $(GOPATH)/src/$(PROTOBUF_PKG)/.git || git clone https://$(PROTOBUF_PKG) $(GOPATH)/src/$(PROTOBUF_PKG)
	(cd $(GOPATH)/src/$(PROTOBUF_PKG) && \
		(test "$$(git describe --tags | head -1)" = "$(PROTOBUF_VERSION)" || \
			(git fetch && git checkout tags/$(PROTOBUF_VERSION))))
	(cd $(GOPATH)/src/$(PROTOBUF_PKG) && go get -v -d $$(go list -f '{{ .ImportPath }}' ./...) && \
	go build -o "$(PWD)/$@" $(PROTOC_GEN_GO_PKG))

########################################################################
##                              PATH                                  ##
########################################################################

# Update PATH with the current directory. This enables the protoc
# binary to discover the protoc-gen-go binary, built inside this
# directory.
export PATH := $(shell pwd):$(PATH)

replication: $(PROTOC) $(PROTOC_GEN_GO)
	$(PWD)/$(PROTOC) -I $(PROTO_INCLUDE) --go_out=plugins=grpc:"$@" --proto_path="$@" ./"$@"/"$@".proto
	(cd "$@"; go mod tidy; go build "$@".pb.go)

podmon: $(PROTOC) $(PROTOC_GEN_GO)
	$(PWD)/$(PROTOC) -I $(PROTO_INCLUDE) --go_out=plugins=grpc:"$@" --proto_path="$@" ./"$@"/"$@".proto
	(cd "$@"; go mod tidy; go build "$@".pb.go)

volumeGroupSnapshot: $(PROTOC) $(PROTOC_GEN_GO)
	$(PWD)/$(PROTOC) -I $(PROTO_INCLUDE) --go_out=plugins=grpc:"$@" --proto_path="$@" ./"$@"/"$@".proto
	(cd "$@"; go mod tidy; go build "$@".pb.go)

clean:
	rm -rf ./replication/replication.pb.go ./podmon/podmon.pb.go ./volumeGroupSnapshot/volumeGroupSnapshot.pb.go

clobber: clean
	rm -rf "$(PROTOC)" "$(PROTOC_GEN_GO)" "$(PROTOC_TMP_DIR)"


