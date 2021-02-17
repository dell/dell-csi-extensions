# csiextensions

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](docs/CODE_OF_CONDUCT.md)
[![License](https://img.shields.io/github/license/dell/dell-csi-extensions)](LICENSE)

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

## Table of Contents

- [Code of Conduct](./docs/CODE_OF_CONDUCT.md)
- [Getting Started Guide](./docs/GETTING_STARTED_GUIDE.md)
- [Branching Strategy](./docs/BRANCHING.md)
- [Contributing Guide](./docs/CONTRIBUTING.md)
- [Maintainers](./docs/MAINTAINERS.md)
- [About](#About)

## Support

Don’t hesitate to ask! Contact the team and community on our [support](./docs/SUPPORT.md) page.
Open an issue if you found a bug on [Github Issues](https://github.com/dell/dell-csi-extensions/issues).

## Versioning

This project is adhering to [Semantic Versioning](https://semver.org/).

## About

Dell CSI Extensions is 100% open source and community-driven. All components are available
under [Apache 2 License](https://www.apache.org/licenses/LICENSE-2.0.html) on
GitHub.