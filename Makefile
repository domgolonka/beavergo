GOBIN=$(shell pwd)/bin
GOFILES=main.go
GONAME=$(shell basename "$(PWD)")
PID=/tmp/go-$(GONAME).pid
## Display this help text
help:
	$(info Available Targets)
	@awk '/^[a-zA-Z\-\_0-9]+:/ {                    \
		nb = sub( /^## /, "", helpMsg );              \
		if(nb == 0) {                                 \
		helpMsg = $$0;                              \
		nb = sub( /^[^:]*:.* ## /, "", helpMsg );   \
		}                                             \
		if (nb)                                       \
		printf "\033[1;31m%-" width "s\033[0m %s\n", $$1, helpMsg;   \
	}                                               \
	{ helpMsg = $$0 }'                              \
	$(MAKEFILE_LIST) | column -ts:

## build the repo
build:
	@echo "Building $(GOFILES) to ./bin"
	go build -o bin/$(GONAME) main.go

## get all the dependencies
get:
	@GOBIN=$(GOBIN) go get main.go
## installs all the dependencies
install:
	go install $(GOFILES)
## runs the build
run:
	go run $(GOFILES)
## This is for development. Restarts after every save
watch:
	@$(MAKE) restart &
	@fswatch -o . -e 'bin/.*' -ignore='Path vendor' | xargs -n1 -I{}  make restart

restart: clear stop clean build start

format: ## Format go code with goimports
	@go get golang.org/x/tools/cmd/goimports
	@goimports -l -w .

format-check: ## Check if the code is formatted
	@go get golang.org/x/tools/cmd/goimports
	@for i in $$(goimports -l .); do echo "[ERROR] Code is not formated run 'make format'" && exit 1; done

check: format-check ## Linting and static analysis
	@if grep -r --include='*.go' -E "fmt.Print|spew.Dump" *; then \
		echo "code contains fmt.Print* or spew.Dump function"; \
		exit 1; \
	fi

	@if test ! -e ./bin/golangci-lint; then \
		curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh; \
	fi
	@./bin/golangci-lint run --timeout 180s -E gosec -E stylecheck -E golint -E goimports

start: ## start defending on localmachine
	@echo "Starting bin/$(GONAME)"
	@./bin/$(GONAME) & echo $$! > $(PID)

stop: ## stop defending on localmachine
	@echo "Stopping bin/$(GONAME) if it's running"
	@-kill `[[ -f $(PID) ]] && cat $(PID)` 2>/dev/null || true

clear: ## clear
	@clear

clean: ## clean
	@echo "Cleaning"
	@GOBIN=$(GOBIN) go clean


.PHONY: format check build get install run watch start stop restart clean