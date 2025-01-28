SHELL = /bin/bash

export PATH := $(PWD)/pact/bin:$(PATH)
export PATH
export CONSUMER_NAME = GoConsumerService
export PROVIDER_NAME = GoProviderService
export PACT_DIR = $(PWD)/pacts
export LOG_DIR = $(PWD)/log
export PACT_BROKER_PROTO = http
export PACT_BROKER_URL = localhost:8080
export PACT_BROKER_USERNAME = broker_username
export PACT_BROKER_PASSWORD = broker_password
export VERSION_COMMIT?=$(shell git rev-parse HEAD)
export VERSION_BRANCH?=$(shell git rev-parse --abbrev-ref HEAD)
