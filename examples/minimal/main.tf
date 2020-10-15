provider "healthchecksio" {
  version = "~> 1.6.0"
}

resource "healthchecksio_check" "bare_minimum" {
  name    = "test-bare-minimum-check"
  timeout = 86400
}

output "ping_url" {
  value = healthchecksio_check.bare_minimum.ping_url
}

output "pause_url" {
  value = healthchecksio_check.bare_minimum.pause_url
}
