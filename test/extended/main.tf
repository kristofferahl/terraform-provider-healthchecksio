resource "healthchecksio_check" "extended" {
  name    = "test-extended-check"
  desc    = "An extended check"
  methods = "POST"

  tags = [
    "go",
    "gophers",
    "unite",
  ]

  grace    = 120 # seconds
  schedule = "0,30 2 * * *"
  timezone = "Europe/Stockholm"
}
