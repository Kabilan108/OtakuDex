MAKEFLAGS += --silent

ldflags := -X 'github.com/kabilan108/OtakuDex/constant.BuiltAt=$(shell date -u)'
ldflags += -X 'github.com/kabilan108/OtakuDex/constant.BuiltBy=$(shell whoami)'
ldflags += -X 'github.com/kabilan108/OtakuDex/constant.Revision=$(shell git rev-parse --short HEAD)'
ldflags += -s
ldflags += -w

build_flags := -ldflags=${ldflags}

all: help

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build        Build the OtakuDex binary"
	@echo "  install      Install the OtakuDex binary"
	@echo "  uninstall    Uninstall the OtakuDex binary"
	@echo "  test         Run the tests"
	@echo "  help         Show this help message"
	@echo ""

install:
	@go install "$(build_flags)"


build:
	@mkdir -p ./build
	@go build -o ./build/ "$(build_flags)"

test:
	@go test ./...

uninstall:
	@rm -f $(shell which OtakuDex)
