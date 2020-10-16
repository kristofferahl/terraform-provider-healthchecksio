# Data Source: healthchecksio_channel

The channel data source allows a check to be associated with pre-defined channels. Read mor at [Healthchecks.io](https://healthchecks.io/docs/configuring_notifications/).

## Example Usage

```hcl
data "healthchecksio_channel" "email" {
  kind = "email"
}

data "healthchecksio_channel" "slack" {
  kind = "slack"
}

data "healthchecksio_channel" "pagerduty" {
  kind = "pd"
}

resource "healthchecksio_check" "test" {
  name  = "test-check"
  desc  = "A test check"
  grace = 120

  channels = [
    data.healthchecksio_channel.email.id,
    data.healthchecksio_channel.slack.id,
    data.healthchecksio_channel.pagerduty.id,
  ]
}
```

## Argument Reference

* `kind` - (Required) Kind of channel
* `name` - (Optional) Name of the channel to search for

## Attribute Reference

* `id` - ID of the channel
