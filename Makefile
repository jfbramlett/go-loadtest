MAKE_SYSTEM ?= ../plat0.base-layer/build/make-system
include $(MAKE_SYSTEM)/build_controls.mk

NAME=plat0-common-sidecar-sample
PROTOBUF=false

go-unit-test: $(GO_FILES) $(PB_SCHEMA_TARGET) go-mods
	@echo "\n MAKE-SYSTEM: running unit tests.\n"
	cd pkg && go test ./...

go-integration-test: $(GO_FILES) $(PB_SCHEMA_TARGET) go-mods start-integration-test-docker-containers go-run-integration-test stop-docker-containers
	@echo "\n MAKE-SYSTEM: completed local integration tests (i.e _test.go(s) in cmd/) .\n"

go-component-test: $(GO_FILES) $(PB_SCHEMA_TARGET) go-mods start-component-test-docker-containers go-run-component-test stop-docker-containers
	@echo "\n MAKE-SYSTEM: completed local component tests (i.e _test.go(s) in test/) .\n"

go-run-integration-test:
	@echo "\n MAKE-SYSTEM: running local integration tests (i.e _test.go(s) in cmd/) .\n"
	@if test -d ./cmd ;then cd cmd && go test ./...;else echo " reMAKE-SYSTEM: no integration tests defined... oh the Humanity!!!\n";fi

go-run-component-test:
	@echo "\n MAKE-SYSTEM: running local component tests (i.e _test.go(s) in test/) .\n"
	@if test -d ./test ;then cd test && go test ./...;else echo " reMAKE-SYSTEM: no component tests defined... oh the Humanity!!!\n";fi

start-integration-test-docker-containers:
ifneq ($(CI),true)
	@echo "Starting docker containers for tests"
	@if test -f docker-compose.yml ; then SERVICE_IMAGE=$(NAME) docker-compose up --scale $(NAME)=0  -d && sleep 5;fi
else
	@echo "CI Build, not starting containers for tests"
endif

start-component-test-docker-containers:
ifneq ($(CI),true)
	@echo "Starting docker containers for tests"
	@if test -f docker-compose.yml ; then SERVICE_IMAGE=$(NAME) docker-compose up -d && sleep 5;fi
else
	@echo "CI Build, not starting containers for tests"
endif

stop-docker-containers:
ifneq ($(CI),true)
	@echo "Stopping docker containers for tests"
	@if test -f docker-compose.yml ; then docker-compose down;fi
else
	@echo "CI Build, not stopping containers for tests"
endif



include $(MAKE_SYSTEM)/target_go_cmd.mk
