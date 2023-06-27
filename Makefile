PKG_NAME=confluence
BINARY_NAME=terraform-provider-confluence
INSTALL_DIR=$(HOME)/.terraform.d/plugins
VERSION=$$(cat VERSION)
TEST?=$$(go list ./...)
GOFMT_FILES?=$$(find . -name '*.go')
provider_path = registry.terraform.io/scottwallacesh/$(PKG_NAME)/$(VERSION)/linux_amd64
all: fmt check test build

check: revive
	~/go/bin/revive

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

build: $(BINARY_NAME)

$(BINARY_NAME):
	go build -v -o $(BINARY_NAME)

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 5m

fmt:
	gofmt -s -w $(GOFMT_FILES)

clean:
	rm -rf bin
	rm -rf site
	rm -f $(BINARY_NAME)

install: $(BINARY_NAME)
	mkdir -p $(INSTALL_DIR)/$(provider_path)
	cp $(BINARY_NAME) $(INSTALL_DIR)/$(provider_path)/$(BINARY_NAME)_v$(VERSION)

uninstall:
	rm -rf $(INSTALL_DIR)/$(provider_path)

revive:
	scripts/get-revive.sh

.PHONY: all build check clean fmt install test testacc uninstall
