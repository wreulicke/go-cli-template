MAKEFILE_DIR:=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))
build/gotestsum:
	mkdir -p build
	GOBIN=$(MAKEFILE_DIR)/build go install gotest.tools/gotestsum@v1.12.1

.PHONY: test
test: build/gotestsum
	mkdir -p build
	build/gotestsum \
		--format standard-verbose \
		--jsonfile build/reports.json \
		--junitfile build/report.xml ./... \
		-- -race -coverprofile=build/coverage.out
