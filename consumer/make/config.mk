SHELL = /bin/bash

export PATH := $(PWD)/pact/bin:$(PATH)
export PATH
export CONSUMER_NAME = GoConsumerService
export PROVIDER_NAME = GoProviderService
export PACT_DIR = $(PWD)/pacts
export LOG_DIR = $(PWD)/log
