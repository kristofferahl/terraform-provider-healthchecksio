variable "healthchecksio_api_key" {
  type        = string
  description = "healthchecksio API Key"
}

variable "healthchecksio_api_url" {
  type        = string
  description = "healthchecksio API Base URL"
  default     = "https://hc.example.com/api/v1"
}

provider "healthchecksio" {
  api_key = var.healthchecksio_api_key
  api_url = var.healthchecksio_api_url
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
