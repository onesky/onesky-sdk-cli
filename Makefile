NAME = onesky-sdk-cli
NAME_BIN=onesky
SHELL = /bin/bash
LOCAL_DIR = $(shell pwd)
LOCAL_BIN = $(LOCAL_DIR)/bin
INSTALL_DIR	= /usr/local/bin

all:
	@echo "################################################################################################"
	@echo "Posible jobs: setup test, build, install, uninstall, build_all, build_mac, build_win, build_nix"
	@echo "  setup == 'make test && make build && make install'"
	@echo "  test - run tests"
	@echo "  build - build binary for current platform"
	@echo "  install - install builded binary"
	@echo "  uninstall - uninstall binary from system"
	@echo "  build_all - create builds for mac, windows and linux"
	@echo "  build_mac, build_win, build_nix - create binaries for specific platform"
	@echo "!by default, the builded binaries are located in local 'bin/' dorrectory"
	@echo "################################################################################################"
	@echo "Use chain 'make test && make build && make install' or 'make setup'"
	@echo ""

install:
	@echo " *** Install binary"
	install $(LOCAL_BIN)/$(NAME_BIN) $(INSTALL_DIR)/$(NAME_BIN)

setup: test build install

uninstall:
	@rm -r $(INSTALL_DIR)/$(NAME_BIN)

build:
	@echo "###############################"
	@echo "# Compile build"
	@echo "###############################"

	$(MAKE) clean
	mkdir -p $(LOCAL_BIN)

	@go version
	GO111MODULE=on go build -v -o $(LOCAL_BIN)/$(NAME_BIN) src/onesky.go
	@echo " *** Done ***"
	@echo "See: " $(LOCAL_BIN)/$(NAME_BIN)
	@echo ""

build_mac:
	@echo "###############################"
	@echo "# Compile MacOS build"
	@echo "###############################"

	$(MAKE) clean
	mkdir -p $(LOCAL_BIN)

	@go version
	GOOS=darwin GOARCH=amd64 go build -v -o $(LOCAL_BIN)/$(NAME_BIN) src/onesky.go
	@echo " *** Build is compiled ***"
	@echo "See: " $(LOCAL_BIN)/$(NAME_BIN)
	@echo ""

build_win:
	@echo "###############################"
	@echo "# Compile Windows build"
	@echo "###############################"

	$(MAKE) clean
	@mkdir -p $(LOCAL_BIN)

	@go version
	GOOS=windows GOARCH=amd64 go build -v -o $(LOCAL_BIN)/$(NAME_BIN) src/onesky.go
	@echo " *** Build is compiled ***"
	@echo "See: " $(LOCAL_BIN)/$(NAME_BIN)
	@echo ""

build_nix:
	@echo "###############################"
	@echo "# Compile Linux build"
	@echo "###############################"

	$(MAKE) clean
	@mkdir -p $(LOCAL_BIN)

	@go version
	GOOS=linux GOARCH=amd64 go build -v -o $(LOCAL_BIN)/$(NAME_BIN) src/onesky.go
	@echo " *** Build is compiled ***"
	@echo "See: " $(LOCAL_BIN)/$(NAME_BIN)
	@echo ""

build_all:
	@echo "###############################"
	@echo "# Compile cross-platfom builds"
	@echo "###############################"

	mkdir -p $(LOCAL_BIN)

	@go version
	go build -v -o $(LOCAL_BIN)/$(NAME_BIN) src/onesky.go

	GOOS=linux GOARCH=amd64 GO111MODULE=on go build -v -o $(LOCAL_BIN)/$(NAME_BIN)_nix src/onesky.go
	GOOS=windows GOARCH=amd64 GO111MODULE=on go build -v -o $(LOCAL_BIN)/$(NAME_BIN)_mac src/onesky.go
	GOOS=darwin GOARCH=amd64 GO111MODULE=on go build -v -o $(LOCAL_BIN)/$(NAME_BIN)_win src/onesky.go

	@echo " *** Done ***"
	@echo "bin/:"
	@ls -lh bin/
	@echo ""

test:
	@echo "###############################"
	@echo "#**** 	TESTING BUILD	 ****#"
	@echo "###############################"

	@echo " ==> Testing..."
	go test -v ./...

	@echo " *** Build is OK ***"
	@echo ""


clean:
	@rm -rf $(LOCAL_BIN)/$(NAME_BIN)

clean_all:
	@rm -rf $(LOCAL_BIN)

purge:
	@rm -rf $(LOCAL_BIN)
	$(MAKE) uninstall
