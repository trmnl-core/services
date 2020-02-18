# Johnnytest1 Service

This is the Johnnytest1 service

Generated with

```
micro new johnnytest1 --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.johnnytest1
- Type: srv
- Alias: johnnytest1

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./johnnytest1-srv
```

Build a docker image
```
make docker
```