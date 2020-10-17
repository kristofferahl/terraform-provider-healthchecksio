---
page_title: "Provider: healthchecksio"
---

# Provider: healthchecksio

A provider used to manage [Healthchecks.io](https://healthchecks.io/) resources.

## Example Usage

```hcl
provider "healthchecksio" {
  api_key = var.healthchecksio_api_key
  version = "~> 1.6.0"
}

variable "healthchecksio_api_key" {
  type        = string
  description = "API Key. To keep secrets out of code first set environment variable TF_VAR_healthchecksio_api_key"
}

resource "healthchecksio_check" "test" {
  name = "test-check"
  desc = "A description for the check"

  tags = [
    "go",
    "gophers",
    "unite",
  ]

  grace = 120 # in seconds
  schedule = "0,30 2 * * *"
  timezone = "Asia/Tokyo"
}

output "ping_url" {
  value = healthchecksio_check.test.ping_url
}
```

## Argument Reference

| Property | Description                 | Environment variable   | Type   | Required |
|----------|-----------------------------|------------------------|--------|----------|
| api_key  | The healthchecks.io API Key | HEALTHCHECKSIO_API_KEY | string | true     |

## Authentication

Your requests to the Healthchecks.io API must be authenticated using an API key. [Read more on how to generate an API key here.](https://healthchecks.io/docs/api/)
