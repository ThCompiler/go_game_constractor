BIN_DIR=./bin
BIN=scg
BIN_WINDOWS=scg.exe
BIN_DEBUG=$(BIN).debug
GCFLAGS_DEBUG="all=-N -l"
SYSTEMD_DIR=~/.config/systemd/user
NOTIFY_SCRIPT_DIR=./scripts
NOTIFY_SCRIPT=notify.sh
NOTIFY_SCRIPT_INSTALL_DIR=~
INSTALL_LOCATION=~/bin
WINDOWS_OS=windows
LINUX_OS=linux
MAC_OS=darwin
ARCH=amd64
VERSION=local


.PHONY: build build-debug test install install-systemd clean release install-notify bin-dir

build: bin-dir
	if [ -z "$(shell git status --porcelain)" ]; then \
		go build -o $(BIN_DIR)/$(BIN) github.com/ThCompiler/go_game_constractor/scg/cmd ; \
		git checkout -- ./version.go; \
	else \
		echo Working directory not clean, commit changes; \
	fi

build-linux: bin-dir
	if [ -z "$(shell git status --porcelain)" ]; then \
		sed -i "s|${VERSION}|${VERSION} $(LINUX_OS)/$(ARCH)|" ./version.go; \
		GOOS=$(LINUX_OS) GOARCH=$(ARCH) go build -o $(BIN_DIR)/$(BIN) github.com/ThCompiler/go_game_constractor/scg/cmd; \
		tar -czvf $(BIN_DIR)/$(BIN).$(LINUX_OS)-$(ARCH).tar.gz $(BIN_DIR)/$(BIN); \
		git checkout -- ./version.go; \
		rm $(BIN_DIR)/$(BIN); \
	else \
		echo Working directory not clean, commit changes; \
	fi

build-darwin: bin-dir
	if [ -z "$(shell git status --porcelain)" ]; then \
		sed -i "s|${VERSION}|${VERSION} $(MAC_OS)/$(ARCH)|" ./version.go; \
		GOOS=$(MAC_OS) GOARCH=$(ARCH) go build -o $(BIN_DIR)/$(BIN) github.com/ThCompiler/go_game_constractor/scg/cmd; \
		tar -czvf $(BIN_DIR)/$(BIN).$(MAC_OS)-$(ARCH).tar.gz $(BIN_DIR)/$(BIN); \
		git checkout -- ./version.go; \
		rm $(BIN_DIR)/$(BIN); \
	else \
		echo Working directory not clean, commit changes; \
	fi

build-windows: bin-dir
	if [ -z "$(shell git status --porcelain)" ]; then \
		sed -i "s|${VERSION}|${VERSION} $(WINDOWS_OS)/$(ARCH)|" ./version.go; \
		GOOS=$(WINDOWS_OS) GOARCH=$(ARCH) go build -o $(BIN_DIR)/$(BIN_WINDOWS) github.com/ThCompiler/go_game_constractor/scg/cmd; \
		zip -9 -y $(BIN_DIR)/$(BIN).$(WINDOWS_OS)-$(ARCH).zip $(BIN_DIR)/$(BIN_WINDOWS); \
		git checkout -- ./version.go; \
		rm $(BIN_DIR)/$(BIN_WINDOWS); \
	else \
		echo Working directory not clean, commit changes; \
	fi

bin-dir:
	mkdir -p $(BIN_DIR)

release: build
	git tag $VERSION; \
	git push origin $VERSION