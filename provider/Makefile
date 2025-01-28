include ./make/config.mk

PACT_GO_VERSION=2.0.8
PACT_DOWNLOAD_DIR=/tmp
ifeq ($(OS),Windows_NT)
	PACT_DOWNLOAD_DIR=$$TMP
endif

install:
	go install github.com/pact-foundation/pact-go/v2@v$(PACT_GO_VERSION)
	pact-go -l DEBUG install --libDir $(PACT_DOWNLOAD_DIR);

install_cli:
	@if [ ! -d pact/bin ]; then\
		echo "=== Installing Pact CLI dependencies ===";\
		curl -fsSL https://raw.githubusercontent.com/pact-foundation/pact-ruby-standalone/master/install.sh | bash;\
  fi

run-provider:
	@go run cmd/main.go

unit:
	@echo "=== 🔨 Running Unit tests 🔨 ==="
	go test ./... -v

provider: export PACT_TEST := true
provider:
	@echo "=== 🔨 Running Provider Pact tests 🔨 ==="
	go test -count=1 -tags=contracts github.com/mesirendon/contract-testing/provider/cmd/ -run "TestUserServicePact" -v

check-provider-deploy:
	@echo "=== 🔬 Checking if we can deploy provider 🔬 ==="
	pact/bin/pact-broker can-i-deploy \
		--pacticipant $(PROVIDER_NAME) \
		--broker-base-url ${PACT_BROKER_PROTO}://$(PACT_BROKER_URL) \
		--broker-username $(PACT_BROKER_USERNAME) \
		--broker-password $(PACT_BROKER_PASSWORD) \
		--version ${VERSION_COMMIT} \
		--to-environment production

record-provider-deploy:
	@echo "=== 💽 Recording deployment of provider 💽 ==="
	pact/bin/pact-broker record-deployment \
		--pacticipant $(PROVIDER_NAME) \
		--broker-base-url ${PACT_BROKER_PROTO}://$(PACT_BROKER_URL) \
		--broker-username $(PACT_BROKER_USERNAME) \
		--broker-password $(PACT_BROKER_PASSWORD) \
		--version ${VERSION_COMMIT} \
		--environment production

.PHONY: unit provider publish
