provider "healthchecksio" {
  version = "~> 1.3"
}

resource "healthchecksio_check" "bare_minimum" {
  name    = "test-bare-minimum-check"
  timeout = 86400
}
