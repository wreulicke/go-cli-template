ARG:=
MAKEFILE_DIR:=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

GOTESTSUM_VERSION:=v1.13.0 # renovate: datasource=github-releases depName=gotestyourself/gotestsum

build/gotestsum:
	mkdir -p build
	GOBIN=$(MAKEFILE_DIR)/build go install gotest.tools/gotestsum@${GOTESTSUM_VERSION}

# Run tests
# if you want to update snapshot, run `make test ARG=-update`
.PHONY: test
test: build/gotestsum
	mkdir -p build
	build/gotestsum \
		--format standard-verbose \
		--jsonfile build/reports.json \
		--junitfile build/reports.xml \
		--  ./... -race -coverprofile=build/coverage.out ${ARG}