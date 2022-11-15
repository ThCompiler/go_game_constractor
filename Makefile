BIN_DIR=./bin
BIN=scg
BIN_WINDOWS=scg.exe
WINDOWS_OS=windows
LINUX_OS=linux
MAC_OS=darwin
ARCH=amd64
VERSION=local
ARM64=arm64
ARM=arm
I386=386

.PHONY: local-build
local-build: bin-dir
	go build -o $(BIN_DIR)/$(BIN) github.com/ThCompiler/go_game_constractor/scg/cmd

.PHONY: build
build: bin-dir
	if [ -z "$(shell git status --porcelain)" ]; then \
		go build -o $(BIN_DIR)/$(BIN) github.com/ThCompiler/go_game_constractor/scg/cmd ; \
	else \
		echo Working directory not clean, commit changes; \
	fi

.PHONY: build-linux
build-linux: bin-dir
	if [ -z "$(shell git status --porcelain)" ]; then \
  		for TYPE in $(ARCH) $(ARM) $(ARM64) $(I386) ; do \
			sed -i "s|${VERSION}|${VERSION} $(LINUX_OS)/$$TYPE|" ./version.go; \
			GOOS=$(LINUX_OS) GOARCH=$$TYPE go build -o $(BIN_DIR)/$(BIN) github.com/ThCompiler/go_game_constractor/scg/cmd; \
			tar -czvf $(BIN_DIR)/$(BIN).$(LINUX_OS)-$$TYPE.tar.gz $(BIN_DIR)/$(BIN); \
			git checkout -- ./version.go; \
			rm $(BIN_DIR)/$(BIN); \
		done ;\
	else \
		echo Working directory not clean, commit changes; \
	fi

.PHONY: build-darwinr
build-darwin: bin-dir
	if [ -z "$(shell git status --porcelain)" ]; then \
  		for TYPE in $(ARCH) $(ARM64) ; do \
			sed -i "s|${VERSION}|${VERSION} $(MAC_OS)/$$TYPE|" ./version.go; \
			GOOS=$(MAC_OS) GOARCH=$$TYPE go build -o $(BIN_DIR)/$(BIN) github.com/ThCompiler/go_game_constractor/scg/cmd; \
			tar -czvf $(BIN_DIR)/$(BIN).$(MAC_OS)-$$TYPE.tar.gz $(BIN_DIR)/$(BIN); \
			git checkout -- ./version.go; \
			rm $(BIN_DIR)/$(BIN); \
		done ;\
	else \
		echo Working directory not clean, commit changes; \
	fi

.PHONY: build-windows
build-windows: bin-dir
	if [ -z "$(shell git status --porcelain)" ]; then \
		for TYPE in $(ARCH) $(ARM) $(ARM64) $(I386) ; do \
			sed -i "s|${VERSION}|${VERSION} $(WINDOWS_OS)/$$TYPE|" ./version.go; \
			GOOS=$(WINDOWS_OS) GOARCH=$$TYPE go build -o $(BIN_DIR)/$(BIN_WINDOWS) github.com/ThCompiler/go_game_constractor/scg/cmd; \
			zip -9 -y $(BIN_DIR)/$(BIN).$(WINDOWS_OS)-$$TYPE.zip $(BIN_DIR)/$(BIN_WINDOWS); \
			git checkout -- ./version.go; \
			rm $(BIN_DIR)/$(BIN_WINDOWS); \
		done ;\
	else \
		echo Working directory not clean, commit changes; \
	fi

.PHONY: bin-dir
bin-dir:
	mkdir -p $(BIN_DIR)

.PHONY: release
release:
	git tag $(VERSION); \
	git push origin $(VERSION)

.PHONY: clean
changelog:
	sh ./workflow/changes.sh $(VERSION) > CURRENT-CHANGELOG.md

.PHONY: clean
clean:
	echo "Cleaning..."; \
	rm -Rf $(BIN_DIR)