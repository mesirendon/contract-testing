include ./make/config.mk

PACT_GO_VERSION=2.0.8
PACT_DOWNLOAD_DIR=/tmp
ifeq ($(OS),Windows_NT)
	PACT_DOWNLOAD_DIR=$$TMP
endif

install:
	go install github.com/pact-foundation/pact-go/v2@v$(PACT_GO_VERSION)
	pact-go -l DEBUG install --libDir $(PACT_DOWNLOAD_DIR);

run-consumer:
	@go run consumer/cmd/main.go

consumer: export PACT_TEST := true
consumer:
	@echo "=== 🔨 Running Consumer Pact tests 🔨 ==="
	go test -tags=contracts -count=1 github.com/mesirendon/contract-testing/consumer/internal/services/users -run 'TestUserClientPact' -v

.PHONY: consumer
