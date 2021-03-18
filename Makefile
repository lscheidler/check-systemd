NAME := check-systemd
VERSION := $(shell egrep "version *=" version.go | cut -d '"' -f 2)
GITHUB_VERSION := v$(VERSION)

all: build

fmt:
	go fmt $(shell find -name \*.go |xargs dirname|sort -u)

lint:
	golint $(shell find -name \*.go |xargs dirname|sort -u)

vet:
	go vet $(shell find -name \*.go |xargs dirname|sort -u)

build: fmt vet
	go build -asmflags -trimpath -o build/linux_amd64/$(NAME)

zip:
	mkdir dist
	zip -j dist/$(NAME)_$(GITHUB_VERSION)_linux_amd64.zip build/linux_amd64/$(NAME)

dist: clean build zip
	cd dist && sha512sum *.zip > $(NAME)_$(GITHUB_VERSION)_SHA512SUM.txt

clean:
	rm -rf build dist

sign:
	gpg --armor --sign --detach-sig dist/$(NAME)_$(GITHUB_VERSION)_linux_amd64.zip
	gpg --armor --sign --detach-sig dist/$(NAME)_$(GITHUB_VERSION)_SHA512SUM.txt

release:
	@echo "| File | Sign  | SHA512SUM |"
	@echo "|------|-------|-----------|"
	@echo "| [$(NAME)_$(GITHUB_VERSION)_linux_amd64.zip](../../releases/download/$(GITHUB_VERSION)/$(NAME)_$(GITHUB_VERSION)_linux_amd64.zip) | [$(NAME)_$(GITHUB_VERSION)_linux_amd64.zip.asc](../../releases/download/$(GITHUB_VERSION)/$(NAME)_$(GITHUB_VERSION)_linux_amd64.zip.asc) | $(shell sha512sum dist/$(NAME)_$(GITHUB_VERSION)_linux_amd64.zip | cut -d " " -f 1) |"
