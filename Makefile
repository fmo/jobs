APP_NAME := fmoj
VERSION := v1.0.0

PLATFORMS := \
	linux/amd64 \
	linux/arm64 \
	darwin/amd64 \
	darwin/arm64 \
	windows/amd64

BUILD_DIR := build

all: clean build

build:
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
		OS=$${platform%%/*}; \
		ARCH=$${platform##*/}; \
		EXT=$${OS=="windows" && echo ".exe" || echo ""}; \
		OUT=$(BUILD_DIR)/$(APP_NAME)_$(VERSION)_$${OS}_$${ARCH}$${EXT}; \
		echo "Building $$OUT"; \
		GOOS=$$OS GOARCH=$$ARCH go build -o $$OUT .; \
	done

clean:
	@rm -rf $(BUILD_DIR)

.PHONY: all build clean
