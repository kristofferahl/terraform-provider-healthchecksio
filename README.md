# Terraform Provider for healthchecks.io

## Provider configuration

| Property | Description                 | Environment variable   | Type   | Required |
|----------|-----------------------------|------------------------|--------|----------|
| api_key  | The healthchecks.io API Key | HEALTHCHECKSIO_API_KEY | string | true     |


## Example

```
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

## Import

Checks can be imported using the uuid, e.g.

```
$ terraform import healthchecksio_check.my_first_check 760ca858-576a-432b-8b1f-378049d7ce96
```
