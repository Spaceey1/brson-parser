OS      ?= $(shell go env GOOS)
ARCH    ?= $(shell go env GOARCH)

EXT :=
ifeq ($(OS),windows)
	EXT := .exe
endif

build:
	GOOS=$(OS) GOARCH=$(ARCH) GOEXPERIMENT=jsonv2 go build -o brson-parser$(EXT) main.go

