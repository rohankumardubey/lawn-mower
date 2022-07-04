BASE = $(CURDIR)
MODULE = catalog-service

.PHONY: all $(MODULE)
all: $(MODULE)

$(MODULE):
	@go build -v -o $(BASE)/bin/$@

.PHONY: test coverage
test-all:
	@go test -v ./...

coverage:
	@go test -coverprofile=coverage.out -v ./...
	@go tool cover -html=coverage.out


.PHONY: clean list
clean:
	@rm -rfv bin
	@exit 0

list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs
