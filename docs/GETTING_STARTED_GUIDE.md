<!--
Copyright (c) 2021 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0
-->

# Getting Started Guidelines
This is an interface definition for Dell EMC CSI Driver extensions. The server side implementation of these interfaces 
should be done at the CSI Driver level. CSI Driver clients can reference these interfaces to make these extension 
interface calls to drivers that have implemented the extension. 

The instructions here are for producing the protobuf stubs (if necessary)

## Building New Project
Git clone the repo. 
```shell
git clone https://github.com/dell/dell-csi-extensions.git
```

Build the protos

```shell
make gen-podmon
```

### Notes/Tips

For the Podmon extension it is recommended to use protoc v3.14.0 protobuf compiler.