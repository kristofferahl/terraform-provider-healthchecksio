provider "healthchecksio" {
  version = "~> 1.6.0"
}

resource "healthchecksio_check" "test" {
  name = "test-check"
  desc = "A test check"

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
