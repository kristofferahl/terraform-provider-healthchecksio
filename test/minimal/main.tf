provider "healthchecksio" {
  version = "~> 1.5.0"
}

resource "healthchecksio_check" "bare_minimum" {
  name    = "test-bare-minimum-check"
  timeout = 86400
}

resource "healthchecksio_check" "extended" {
  name    = "test-extended-check"
  desc    = "An extended check"
  timeout = 86400
}
