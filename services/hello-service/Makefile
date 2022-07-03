BASE = $(CURDIR)
MODULE = hello-service

.PHONY: all $(MODULE)
all: $(MODULE)

$(MODULE):
	@go build -v -o $(BASE)/bin/$@

.PHONY: clean list
clean:
	@rm -rfv bin
	@exit 0

list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs
