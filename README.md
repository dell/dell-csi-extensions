# csiextensions

This repo holds definitions of additional grpc protocols we want our CSI drivers to support.
These are defined as .proto files, and compiled using the protoc program to generate a go file.

Multiple protocols can be supported, each in a separate top level directory, forming its own go package.
The initial one is podmon, which defines additional driver calls that the podmon program uses.

Additional calls for the replication effort are anticipated, and could be put in a directory like
replication.

The top level Makefile can be used to compile the protocol definitions in .proto files to .go source files,
and then build the go source files..
It is anticipated that CSI drivers would then include this repo to reference the package(s) they require
as defined in the in the .go files.

The gocsi package will be extended to allow the registration of additional Servers. Podmon defines a
PodmonServer, similar to the ControllerServer, NodeServer, and IdentityServer coming from the CSI spec.

