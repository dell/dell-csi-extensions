gen-podmon: protoc
	go clean ./...
	export PATH="$$PATH:./bin" && $(PROTOC) -I $(PROTO_INCLUDE) --go_out=plugins=grpc:podmon --proto_path podmon podmon.proto
	(cd podmon; go build podmon.pb.go)

protoc:
ifeq (, $(shell ls .proto 2> /dev/null))
	@{ \
	mkdir .proto ;\
	cd .proto ;\
	curl -Lo protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/protoc-3.14.0-win64.zip ;\
	unzip protoc.zip ;\
	rm protoc.zip ;\
	}
PROTOC=.proto/bin/protoc
PROTO_INCLUDE=.proto/include
else
PROTOC=.proto/bin/protoc
PROTO_INCLUDE=.proto/include
endif
