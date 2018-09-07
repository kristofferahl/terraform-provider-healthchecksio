# Terraform Provider for healthchecks.io

## Provider configuration

| Property | Description | Environment variable | Type | Required |
|---|
| api_key | The healthchecks.io API Key | HEALTHCHECKSIO_API_KEY | string | true |


## Example

```
resource "healthchecksio_check" "test" {
  name = "test-check"

  tags = [
    "go",
    "gophers",
    "unite",
  ]

  grace = 120
  schedule = "0,30 2 * * *"
  timezone = "Asia/Tokyo"
}
```

More examples can be found in the [examples directory](./examples).
