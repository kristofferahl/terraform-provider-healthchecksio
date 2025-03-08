---
page_title: "Development"
---

# Development

## Pre-requisites

- [Go](https://golang.org/) 1.19 or later
- [Terraform](https://www.terraform.io/) v0.13 or later
- [Healthchecks.io](https://healthchecks.io/) account and an API key
- [goreleaser](https://goreleaser.com/) 2.7.0

## Help

```bash
./run help
```

## Running commands - local

```bash
./run <command> [<arg1> <arg2> ...]
```

## Running commands - in dev mode

This is useful to test the provider during development.
Using `./run dev ...` will:

- build the provider (go build .)
- setup overrides for provider installation (see dev.tfrc)
- run the command as specified

```bash
./run dev <command> [<arg1> <arg2> ...]
```

## Running commands - in docker

Most commands can also be executed using docker. Simply run the commands like below.

```bash
./run docker <command> [<arg1> <arg2> ...]
```

## Building the provider

```bash
./run build
./run docker build
```

## Running unit tests

```bash
./run test
./run docker test
```

## Running integration tests

**NOTE**: Depends on the output of the build command

```bash
./run test-integration
./run docker test-integration
```

## Running examples

**NOTE**: Depends on the output of the build command

```bash
./run examples [<example>]
./run docker examples [<example>]
```

## Preparing a release

```bash
./run prepare-release
```

This will prompt you for the next version. Make sure the version you choose follows Semantic Versioning and is prefixed with `v` (example version: `v1.0.0`).

## Automatic release pipeline

When pushing a tag that match the terraform provider tag conventions, the automatic release pipeline will be triggered.
See [Publishing Providers](https://www.terraform.io/docs/registry/providers/publishing.html) for more information.

## Releasing a new version manually

**NOTE** In general, manual releases should be avoided in favour of using the automatic release pipeline.

**NOTE** Only tags that follow semantic versioning should be released to the public.

```bash
./run release
./run docker release
```
