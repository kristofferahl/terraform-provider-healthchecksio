# Resource: healthchecksio_check

The check resource allows a check to be created in [Healthchecks.io](https://healthchecks.io).

## Example Usage

```hcl
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
}
```

## Argument Reference

* `name` - (Required) Name of the check
* `tags` - (Optional) Tags associated with the check
* `timeout` - (Optional) Timeout period of the check
* `grace` - (Optional) Grace period for the check
* `schedule` - (Optional) A cron expression defining the check's schedule
* `timezone` - (Optional) Timezone used for the schedule
* `channels` - (Optional) Channels integrated with the check
* `desc` - (Optional) Description of the check

## Attribute Reference

* `ping_url` - Ping URL associated with this check
* `pause_url` - Pause URL associated with this check
