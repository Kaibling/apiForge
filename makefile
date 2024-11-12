CURRENT_TAG := $(shell git describe --tags --abbrev=0)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

# Extract major, minor, patch parts
MAJOR := $(shell echo $(CURRENT_TAG) | cut -d. -f1 | sed 's/v//')
MINOR := $(shell echo $(CURRENT_TAG) | cut -d. -f2)
PATCH := $(shell echo $(CURRENT_TAG) | cut -d. -f3)

# Helper functions to increment versions
define increment_major
v$(shell echo $$(($(MAJOR) + 1))).0.0
endef

define increment_minor
v$(MAJOR).$(shell echo $$(($(MINOR) + 1))).0
endef

define increment_patch
v$(MAJOR).$(MINOR).$(shell echo $$(($(PATCH) + 1)))
endef

# Targets to increment versions and push tags
# .PHONY: major minor patch
lint:
	golangci-lint run

major:
	@git tag $(call increment_major)
	@git push origin $(call increment_major)

minor:
	@git tag $(call increment_minor)
	@git push origin $(call increment_minor)

patch:
	@git tag $(call increment_patch)
	@git push origin $(call increment_patch)

current-tag:
	@echo "Current tag: $(CURRENT_TAG)"
	@echo "Branch: $(BRANCH)"

deps:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.61.0
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install mvdan.cc/gofumpt@latest
	go install github.com/daixiang0/gci@latest

fmt:
	gofumpt -l -w .

vuln:
	govulncheck ./...

gci:
	gci write --skip-generated -s standard -s default .

full-lint: gci fmt lint vuln
