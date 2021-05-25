# Wolfpack Makefile is used to drive the build and installation of Wolfpack
# this is meant to be used with a local copy of code repository.

tests:
	@echo "Launching Tests"

clean:
	@echo "Cleaning up build junk"

build:
	@echo "Building from source"
	go build

install:
	@echo "Installing from source"
	go install

release:
	@echo "Creating Release Binaries"
	go get -u github.com/mitchellh/gox
	mkdir -p build
	gox -output="build/{{.OS}}_{{.Arch}}/{{.Dir}}"