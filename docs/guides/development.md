---
page_title: "Development"
---

# Development

## Pre-requisites
- [Go](https://golang.org/) 1.13 or later
- [Terraform](https://www.terraform.io/) v0.12 or later
- [Healthchecks.io](https://healthchecks.io/) account and an API key
- [goreleaser](https://goreleaser.com/) 0.110.0


## Help

```bash
./run help
```

## Running commands - local

```bash
./run <command> [<arg1> <arg2> ...]
```

## Running commands - in docker

Most commands can also be executed using docker. Simply run the commands like below.

```bash
./run docker <command> [<arg1> <arg2> ...]
```

## Running unit tests

```bash
./run test
./run docker test
```

## Running integration tests

```bash
./run test-integration
./run docker test-integration
```

## Building the provider

```bash
./run build
./run docker build
```

## Running examples

```bash
./run examples [<example>]
./run docker examples [<example>]
```

## Releasing a new version

```bash
./run release
./run docker release
```
