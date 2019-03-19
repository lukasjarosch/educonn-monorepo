VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || cat $(CURDIR)/.version 2> /dev/null || echo v0.0.0)
BUILDVER = module:$(MODULE),version:$(VERSION),build:$(shell date +"%Y%m%d.%H%M%S.%N.%z")
BASE = $(CURDIR)
MODULE = accountsvc

.PHONY: all $(MODULE)
all: version $(MODULE)


$(MODULE):| $(BASE)
	@go build -v -o $(BASE)/bin/$@

$(BASE):
	@mkdir -p $(dir $@)

version:
	@echo Version ${VERSION}
