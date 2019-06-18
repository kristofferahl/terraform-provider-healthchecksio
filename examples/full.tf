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

  grace    = 120 # seconds
  schedule = "0,30 2 * * *"
  timezone = "Asia/Tokyo"

  channels = [
    data.healthchecksio_channel.email.id,
    data.healthchecksio_channel.slack.id,
    data.healthchecksio_channel.pagerduty.id,
  ]
}

data "healthchecksio_channel" "email" {
  kind = "email"
}

data "healthchecksio_channel" "slack" {
  kind = "slack"
}

data "healthchecksio_channel" "pagerduty" {
  kind = "pd"
}
