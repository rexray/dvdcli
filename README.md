# dvdcli [![Build Status](http://travis-ci.org/emccode/dvdcli.svg?branch=master)](https://travis-ci.org/emccode/dvdcli) [![Coverage Status](http://coveralls.io/repos/emccode/dvdcli/badge.svg?branch=master&service=github)](https://coveralls.io/github/emccode/dvdcli?branch=master)
`dvdcli` is a Docker Volume Driver client CLI.  The CLI is used enable any application to call the CLI to perform external Volume Management to a Linux host.  This project exposes the `Docker` Volume Driver eco-system for additional use cases.

Docker Volume Driver
--------------------

As of `Docker` 1.7, there was a Volume Driver API defined that allows `Docker` to work with external storage providers.  The API documentation is available [here](https://github.com/docker/docker/blob/master/docs/extend/index.md).  The existing API has five main features, `Create`, `Remove`, `Mount`, `Unmount`, `Path`.  In order to leverage these features, a Volume Driver is created that manages and orchestrates the API calls to specific storage platforms.  It is these Volume Drivers that this project is using in order to enable Volume Management.  See official plugins [here](https://github.com/docker/docker/blob/master/docs/extend/plugins.md).


# Installation
Installing `dvdcli` couldn't be easier.

```bash
curl -sSL https://dl.bintray.com/emccode/dvdcli/install | sh -
```

# Downloading
There are also pre-built binaries at the following locations:

Repository | Version | Description
---------- | ------- | -----------
[unstable](https://dl.bintray.com/emccode/dvdcli/unstable) | [ ![Download](https://api.bintray.com/packages/emccode/dvdcli/unstable/images/download.svg) ](https://dl.bintray.com/emccode/dvdcli/unstable/latest/) | The most up-to-date, bleeding-edge, and often unstable dvdcli binaries.
[staged](https://dl.bintray.com/emccode/dvdcli/staged)   | [ ![Download](https://api.bintray.com/packages/emccode/dvdcli/staged/images/download.svg) ](https://dl.bintray.com/emccode/dvdcli/staged/latest/) | The most up-to-date, release candidate dvdcli binaries.
[stable](https://dl.bintray.com/emccode/dvdcli/stable)   | [ ![Download](https://api.bintray.com/packages/emccode/dvdcli/stable/images/download.svg) ](https://dl.bintray.com/emccode/dvdcli/stable/latest/) | The most up-to-date, stable dvdcli binaries.

CLI
---
The CLI takes advantage of native `Docker` packages which means the functionality available from the CLI mimics what the `Docker` daemon does to manage Volumes.  In addition, the SPEC relating to how the API communicates, ie looking up socket files or HTTP endpoints, is respected entirely.

A typical command from Docker to take advantages of Volume Drivers would be something like `docker run -ti --volume-driver=rexray -v volume:/to busybox`.

For `dvdcli`, the command is very similar.
`dvdcli mount --volumedriver=rexray --volumename=test`.  The main difference is that in Docker's case the mount path is specified as a target inside of a container.  `dvdcli` leaves the mount at the host level, and it is then up to the consumer to make the path available in whatever way necessary.

Examples
--------
General note about the functionality.  Based on mimicing `Docker` functionality, a create is called at the beginning of each of these operations.  The existing Volume Driver spec expects idempotent operations.

#### Mount a volume
```
dvdcli mount --volumedriver=rexray --volumename=test123456789  \
  --volumeopts=size=5 --volumeopts=iops=150 --volumeopts=volumetype=io1 \
  --volumeopts=newFsType=ext4 --volumeopts=overwritefs=true
```

#### Unmount a volume
```
dvdcli unmount --volumedriver=rexray --volumename=test
```

#### Create a volume
```
dvdcli create --volumedriver=rexray --volumename=test123456789 \
  --volumeopts=size=5 --volumeopts=iops=150 --volumeopts=volumetype=io1 \
  --volumeopts=newFsType=ext4 --volumeopts=overwritefs=true
```

#### Remove a volume
```
dvdcli remove --volumedriver=rexray --volumename=test
```

#### Return path of a mounted volume
```
dvdcli path --volumedriver=rexray --volumename=test
```

#### Extra Options
option|description
------|-----------
size|Size in GB
IOPS|IOPS
volumeType|Type of Volume or Storage Pool
newFsType|FS Type for Volume if filesystem unknown
overwriteFs|Overwrite existing known filesystem
volumeName|Create from an existing volume name
volumeID|Creat from an existing volume ID
snapshotName|Create from an existing snapshot name
snapshotID|Create from an existing snapshot ID



EMC {code} - REX-Ray
-------
The `REX-Ray` project is a good example of a service that can expose a valid Volume Driver endpoint that can be used and is available [here](https://github.com/emccode/dvdcli).  The options mentioned above are dependent on the Volume Driver.  `REX-Ray` does implement options as listed above.

# Licensing
---------
Licensed under the Apache License, Version 2.0 (the “License”); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an “AS IS” BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

# Support
-------
Please file bugs and issues at the Github issues page. For more general discussions you can contact the EMC Code team at <a href="https://groups.google.com/forum/#!forum/emccode-users">Google Groups</a> or tagged with **EMC** on <a href="https://stackoverflow.com">Stackoverflow.com</a>. We also are available for real-time feedback on the EMC {code} Slack channel [here](http://community.emccode.com/).  The code and documentation are released with no warranties or SLAs and are intended to be supported through a community driven process.
