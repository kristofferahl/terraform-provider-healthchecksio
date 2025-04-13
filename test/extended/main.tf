resource "healthchecksio_check" "extended" {
  name    = "test-extended-check"
  slug    = "test-extended"
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

  channels = [
    data.healthchecksio_channel.email.id
  ]
}

data "healthchecksio_channel" "email" {
  kind = "email"
}
