# Dell CSI Extensions

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](docs/CODE_OF_CONDUCT.md)
[![License](https://img.shields.io/github/license/dell/dell-csi-extensions)](LICENSE)

This repository holds definitions of additional `grpc` protocols which can be implemented by CSI drivers to build upon existing
functionality supported via the Container Storage Interface ([CSI](https://github.com/container-storage-interface/spec)).
These are defined as .proto files, and compiled using the protoc program to generate a go file.

Current set of supported protocols:

* podmon -  Extensions to CSI to enable CSI drivers to report host connectivity status
* replication - Extensions to CSI to enable CSI drivers to leverage storage array based
 replication within Kubernetes clusters

You can utilize the provided `Makefile` to compile protocol definitions in .proto files to go source files and then build
the source files.

* `podmon` to compile and update files for podmon
* `replication` to compile and update files for replication.

Each protocol is available as a `go module` and can be included separately in CSI drivers.

If you are using `gocsi` for developing CSI driver, these extensions APIs can be used in CSI drivers by registering additional grpc servers while creating the `Storage Plugin` object.
This facility is available in the forked version of gocsi [here](https://github.com/dell/gocsi).

For using `podmon`, include "github.com/dell/dell-csi-extensions/podmon" in your go module dependencies.  
For using `replication`, include "github.com/dell/dell-csi-extensions/replication" in your go module dependencies.

## Table of Contents

* [Code of Conduct](https://github.com/dell/csm/blob/main/docs/CODE_OF_CONDUCT.md)
* [Maintainer Guide](https://github.com/dell/csm/blob/main/docs/MAINTAINER_GUIDE.md)
* [Committer Guide](https://github.com/dell/csm/blob/main/docs/COMMITTER_GUIDE.md)
* [Contributing Guide](https://github.com/dell/csm/blob/main/docs/CONTRIBUTING.md)
* [List of Adopters](https://github.com/dell/csm/blob/main/docs/ADOPTERS.md)
* [Support](https://github.com/dell/csm/blob/main/docs/SUPPORT.md)
* [Security](https://github.com/dell/csm/blob/main/docs/SECURITY.md)
* [About](#about)

## Support

Don’t hesitate to ask! Contact the team and community on our [support](./docs/SUPPORT.md) page.
Open an issue if you found a bug on [Github Issues](https://github.com/dell/dell-csi-extensions/issues).

## Versioning

This project is adhering to [Semantic Versioning](https://semver.org/).

## About

Dell CSI Extensions is 100% open source and community-driven. All components are available
under [Apache 2 License](https://www.apache.org/licenses/LICENSE-2.0.html) on
GitHub.
