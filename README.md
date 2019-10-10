# Terraform Provider for healthchecks.io

[![CircleCI](https://circleci.com/gh/kristofferahl/terraform-provider-healthchecksio/tree/master.svg?style=svg)](https://circleci.com/gh/kristofferahl/terraform-provider-healthchecksio/tree/master)

## Usage

### Provider configuration

| Property | Description                 | Environment variable   | Type   | Required |
|----------|-----------------------------|------------------------|--------|----------|
| api_key  | The healthchecks.io API Key | HEALTHCHECKSIO_API_KEY | string | true     |


### Example

```terraform
provider "healthchecksio" {
  api_key = var.healthchecks_io_api_key
  version = "~> 1.3"
}

variable "healthchecks_io_api_key" {
  type        = string
  description = "API Key. tfvars can't be used here, to keep secrets out of code first set environment TF_VAR_healthchecks_io_api_key"
}

resource "healthchecksio_check" "test" {
  name = "test-check"

  tags = [
    "go",
    "gophers",
    "unite",
  ]

  grace = 120 # in seconds
  schedule = "0,30 2 * * *"
  timezone = "Asia/Tokyo"

  channels = [
    data.healthchecksio_channel.pagerduty.id,
  ]
}

data "healthchecksio_channel" "pagerduty" {
  kind = "pd"
}
```

More examples can be found in the [examples directory](./examples).

### Import

Checks can be imported using the uuid, e.g.

```
$ terraform import healthchecksio_check.my_first_check 760ca858-576a-432b-8b1f-378049d7ce96
```

## Development

### Pre-requisites
- [Go](https://golang.org/) 1.13 or later
- [Terraform](https://www.terraform.io/) v0.12 or later
- [Healthchecks.io](https://healthchecks.io/) account and an API key
- [goreleaser](https://goreleaser.com/) 0.110.0


### Help

```bash
./run help
```

### Running commands - local

```bash
./run <command> [<arg1> <arg2> ...]
```

### Running commands - in docker

Most commands can also be executed using docker. Simply run the commands like below.

```bash
./run docker <command> [<arg1> <arg2> ...]
```

### Running unit tests

```bash
./run test
./run docker test
```

### Running integration tests

```bash
./run test-integration
./run docker test-integration
```

### Building the provider

```bash
./run build
./run docker build
```

### Running examples

```bash
./run examples [<example>]
./run docker examples [<example>]
```

### Releasing a new version

```bash
./run release
./run docker release
```

## Contributors
- [masutaka](https://github.com/masutaka)
- [kristofferahl](https://github.com/kristofferahl)
- [rossmckelvie](https://github.com/rossmckelvie)
