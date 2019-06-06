# WIP: Currently this is only support macOS.
# 2019-06-06: this does not seem to be working for me, though I did update examples using a working config from a project

OS_NAME := $(shell uname -s)
ARCH := $(shell uname -m)
DOWNLOAD_DIR := $(shell mktemp -d)
PLUGIN_DIR := .terraform/plugins/$(shell echo $(OS_NAME) | tr A-Z a-z)_amd64 # hard coding

PROVIDER_HEALTHCHECKSIO_VERSION := 1.2.0
PROVIDER_HEALTHCHECKSIO_ARCHIVE_URL := https://github.com/kristofferahl/terraform-provider-healthchecksio/releases/download/v$(PROVIDER_HEALTHCHECKSIO_VERSION)/terraform-provider-healthchecksio_$(PROVIDER_HEALTHCHECKSIO_VERSION)_$(OS_NAME)_$(ARCH).tar.gz
PROVIDER_HEALTHCHECKSIO_DEST_NAME := $(PLUGIN_DIR)/terraform-provider-healthchecksio_v$(PROVIDER_HEALTHCHECKSIO_VERSION)

all: init

.PHONY: init
init: $(PROVIDER_HEALTHCHECKSIO_DEST_NAME)

$(PROVIDER_HEALTHCHECKSIO_DEST_NAME):
	curl -sL $(PROVIDER_HEALTHCHECKSIO_ARCHIVE_URL) | tar xz -C $(DOWNLOAD_DIR)
	mkdir -p $(PLUGIN_DIR)
	cp $(DOWNLOAD_DIR)/terraform-provider-healthchecksio_v$(PROVIDER_HEALTHCHECKSIO_VERSION) $(PROVIDER_HEALTHCHECKSIO_DEST_NAME)
	terraform init
	rm -rf $(DOWNLOAD_DIR)

.PHONY: clean
clean:
	rm -rf $(PLUGIN_DIR)
