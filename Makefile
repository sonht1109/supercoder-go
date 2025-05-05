VERSION ?= "v0.0.1"
MAIN ?= cmd/main.go
DIR ?= bundles
VERSION_DIR = github.com/sonht1109/supercoder-go/internal/config.Version=$(VERSION)

build:
	@echo "Building SuperCoder version $(VERSION)"

	# Linux
	GOOS=linux GOARCH=amd64 go build -o $(DIR)/supercoder-linux-amd64 -ldflags "-X $(VERSION_DIR)" $(MAIN)
	GOOS=linux GOARCH=arm64 go build -o $(DIR)/supercoder-linux-arm64 -ldflags "-X $(VERSION_DIR)" $(MAIN)

	# macOS
	GOOS=darwin GOARCH=amd64 go build -o $(DIR)/supercoder-darwin-amd64 -ldflags "-X $(VERSION_DIR)" $(MAIN)
	GOOS=darwin GOARCH=arm64 go build -o $(DIR)/supercoder-darwin-arm64 -ldflags "-X $(VERSION_DIR)" $(MAIN)

	# Windows
	GOOS=windows GOARCH=amd64 go build -o $(DIR)/supercoder-windows-amd64.exe -ldflags "-X $(VERSION_DIR)" $(MAIN)
	GOOS=windows GOARCH=arm64 go build -o $(DIR)/supercoder-windows-arm64.exe -ldflags "-X $(VERSION_DIR)" $(MAIN)