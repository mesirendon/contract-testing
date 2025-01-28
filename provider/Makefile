include ./make/config.mk

PACT_GO_VERSION=2.0.8
PACT_DOWNLOAD_DIR=/tmp
ifeq ($(OS),Windows_NT)
	PACT_DOWNLOAD_DIR=$$TMP
endif

install:
	go install github.com/pact-foundation/pact-go/v2@v$(PACT_GO_VERSION)
	pact-go -l DEBUG install --libDir $(PACT_DOWNLOAD_DIR);

run-provider:
	@go run cmd/main.go

unit:
	@echo "=== 🔨 Running Unit tests 🔨 ==="
	go test ./... -v

provider: export PACT_TEST := true
provider:
	@echo "=== 🔨 Running Provider Pact tests 🔨 ==="
	go test -count=1 -tags=contracts github.com/mesirendon/contract-testing/provider/cmd/ -run "TestUserServicePact" -v


.PHONY: unit provider publish
