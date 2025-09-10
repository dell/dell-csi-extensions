# :lock: **Important Notice**
Starting with the release of **Container Storage Modules v1.16.0**, this repository will no longer be maintained as an open source project. Future development will continue under a closed source model. This change reflects our commitment to delivering even greater value to our customers by enabling faster innovation and more deeply integrated features with the Dell storage portfolio.<br>
For existing customers using Dell’s Container Storage Modules, you will continue to receive:
* **Ongoing Support & Community Engagement**<br>
       You will continue to receive high-quality support through Dell Support and our community channels. Your experience of engaging with the Dell community remains unchanged.
* **Streamlined Deployment & Updates**<br>
        Deployment and update processes will remain consistent, ensuring a smooth and familiar experience.
* **Access to Documentation & Resources**<br>
       All documentation and related materials will remain publicly accessible, providing transparency and technical guidance.
* **Continued Access to Current Open Source Version**<br>
       The current open-source version will remain available under its existing license for those who rely on it.

Moving to a closed source model allows Dell’s development team to accelerate feature delivery and enhance integration across our Enterprise Kubernetes Storage solutions ultimately providing a more seamless and robust experience.<br>
We deeply appreciate the contributions of the open source community and remain committed to supporting our customers through this transition.<br>

For questions or access requests, please contact the maintainers via [Dell Support](https://www.dell.com/support/kbdoc/en-in/000188046/container-storage-interface-csi-drivers-and-container-storage-modules-csm-how-to-get-support).

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
