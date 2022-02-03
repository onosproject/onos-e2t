<!--
SPDX-FileCopyrightText: 2020 Open Networking Foundation <info@opennetworking.org>

SPDX-License-Identifier: Apache-2.0
-->
# onos-e2t
E2 AP Termination module for ONOS SD-RAN (µONOS Architecture)

## Overview

The E2 Termination (E2T) acts as an intelligent proxy and adapter for managing the interactions betwen SD-RAN components and E2 nodes. The southbound of E2T implements the E2AP specification (ASN.1 over SCTP), and the northbound implements the onos-e2t API as specified within the [onos-api]. Messages traveling southbound through E2T nodes are converted from Protobuf to ASN.1, and those received from the environment are converted from ASN.1 to Protobuf before they're propagated up through the northbound API. This process can be extended for service models with plugins.

The E2 termination is shipped as a [Docker] image and deployed with [Helm]. To build the Docker image, run `make images`.

[onos-api]: https://github.com/onosproject/onos-api
[Docker]: https://www.docker.com/
[Helm]: https://helm.sh
