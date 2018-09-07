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

  # timeout = 300
  schedule = "0,30 2 * * *"
  timezone = "Asia/Tokyo"
}
