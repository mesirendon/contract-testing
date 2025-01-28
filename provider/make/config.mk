SHELL = /bin/bash

export PATH
export PROVIDER_NAME = GoProviderService
export PACT_BROKER_PROTO = http
export PACT_BROKER_URL = localhost:8080
export PACT_BROKER_USERNAME = broker_username
export PACT_BROKER_PASSWORD = broker_password
export VERSION_COMMIT?=$(shell git rev-parse HEAD)
export VERSION_BRANCH?=$(shell git rev-parse --abbrev-ref HEAD)
