# dvdcli
`dvdcli` is a Docker Volume Driver client CLI.  The CLI is used enable any application to call the CLI to perform external Volume Management to a Linux host.  This project exposes the `Docker` Volume Driver eco-system for additional use cases.

Docker Volume Driver
--------------------

As of `Docker` 1.7, there was a Volume Driver API defined that allows `Docker` to work with external storage providers.  The API documentation is available [here](https://github.com/docker/docker/blob/master/docs/extend/index.md).  The existing API has five main features, `Create`, `Remove`, `Mount`, `Unmount`, `Path`.  In order to leverage these features, a Volume Driver is created that manages and orchestrates the API calls to specific storage platforms.  It is these Volume Drivers that this project is using in order to enable Volume Management.  See official plugins [here](https://github.com/docker/docker/blob/master/docs/extend/plugins.md).

CLI
---
The CLI takes advantage of native `Docker` packages which means the functionality available from the CLI mimics what the `Docker` daemon does to manage Volumes.  In addition, the SPEC relating to how the API communicates, ie looking up socket files or HTTP endpoints, is respected entirely.

A typical command from Docker to take advantages of Volume Drivers would be something like `docker run -ti --volume-driver=rexray -v volume:/to busybox`.

For `dvdcli`, the command is very similar.
`dvdcli mount --volumedriver=rexray --volumename=test`.  The main difference is that in Docker's case the mount path is specified as a target inside of a container.  `dvdcli` leaves the mount at the host level, and it is then up to the consumer to make the path available in whatever way necessary.

Examples
--------
General note about the functionality.  Based on mimicing `Docker` functionality, a create is called at the beginning of each of these operations.  The existing Volume Driver spec expects idempotent operations.

`dvdcli mount --volumedriver=rexray --volumename=test`

`dvdcli unmount --volumedriver=rexray --volumename=test`

`dvdcli create --volumedriver=rexray --volumename=test`

`dvdcli remove --volumedriver=rexray --volumename=test`

`dvdcli path --volumedriver=rexray --volumename=test`

EMC {code} - REX-Ray
-------
The `REX-Ray` project is a good example of a service that can expose a valid Volume Driver endpoint that can be used and is available [here](https://github.com/emccode/rexray).

# Licensing
---------
Licensed under the Apache License, Version 2.0 (the “License”); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an “AS IS” BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

# Support
-------
Please file bugs and issues at the Github issues page. For more general discussions you can contact the EMC Code team at <a href="https://groups.google.com/forum/#!forum/emccode-users">Google Groups</a> or tagged with **EMC** on <a href="https://stackoverflow.com">Stackoverflow.com</a>. We also are available for real-time feedback on the EMC {code} Slack channel [here](http://community.emccode.com/).  The code and documentation are released with no warranties or SLAs and are intended to be supported through a community driven process.
