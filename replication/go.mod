module github.com/dell/dell-csi-extensions/replication

go 1.16

replace github.com/dell/dell-csi-extensions/common v1.0.0 => ../common

require (
	github.com/dell/dell-csi-extensions/common v1.0.0
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
)
