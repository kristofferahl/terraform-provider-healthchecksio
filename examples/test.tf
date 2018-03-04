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
}
