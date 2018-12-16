APP_NAME = todo
BUILD_DIR = output
RELEASE_DIR = release
INSTALL_DIR = /usr/local/bin
DARWIN = darwin
LINUX = linux
VERSION = `git tag |sort -Vr |head -1`

build:
	go mod vendor
	@mkdir -p $(BUILD_DIR)
	@mkdir -p $(RELEASE_DIR)
	@GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(DARWIN)/$(APP_NAME) -ldflags "-s -w"
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(LINUX)/$(APP_NAME) -ldflags "-s -w"
	@echo "build success!"
	#@tar -cvzf $(RELEASE_DIR)/$(VERSION).tar.gz -C $(BUILD_DIR)/$(DARWIN) .
	#@tar -cvzf $(RELEASE_DIR)/$(VERSION).tar.xz -C $(BUILD_DIR)/$(LINUX) .
	zip -j $(RELEASE_DIR)/$(VERSION).zip $(BUILD_DIR)/$(DARWIN)/*
	tar -cvzf $(RELEASE_DIR)/$(VERSION).tar.gz -C $(BUILD_DIR)/$(LINUX) .
	@echo "release success!"

