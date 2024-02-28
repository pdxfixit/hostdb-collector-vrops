# HostDB Collector for vRealize Operations Manager

Queries the vROps REST API, to get data about all of the virtual machines in each vCenter instance, and sends that data to HostDB.

## Getting Started

This section will describe the process of developing the collector.
Please see [Deployment](#deployment) for notes on how the collector is used when deployed.

### Prerequisites

The vROps collector requires a few things to operate:

* An instance of vROps to query.
* A HostDB instance to write to.

For development, you'll also need:

* Docker
* Golang

### Installing

The collector is a golang binary, and after compilation, it can be run on any Linux x86 system. No installation necessary.

## Running tests

This should be as simple as `make test`. It will execute `go fmt`, `go vet`, `golint`, `errcheck` and `go test`.

## Builds

The binary is compiled and containerized, and pushed to registry for deployment.

`registry.pdxfixit.com/hostdb-collector-vrops`

## Deployment

The vROps collector runs from Kubernetes as a cron job, defined in the [hostdb-server Helm chart](https://github.com/pdxfixit/hostdb-server-chart/blob/master/hostdb-server/templates/collector-vrops.yaml).

## Debugging

Set the environment variable `HOSTDB_COLLECTOR_VROPS_COLLECTOR_DEBUG` to true, and the collector will output additional detail, *including secrets*.
In addition, if the variable `HOSTDB_COLLECTOR_VROPS_COLLECTOR_SAMPLE_DATA` is set to true, the collector will output all collected data to file, instead of sending it to HostDB.

## Built With

Build will run tests, compile the golang binary, create a container including the binary, and upload that container image to the registry for use.

## Authors & Support

- Email: info@pdxfixit.com
