provider "healthchecksio" {
  # api_key = "{your api key}"
}

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

resource "healthchecksio_check" "bare_minimum" {
  name    = "test-bare-minimum-check"
  timeout = 86400
}
