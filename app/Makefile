
.PHONY: browser-test

test:
	go test go/lib/lib_test.go go/lib/lib.go go/lib/helm_engine.go

build:
	mkdir -p dist
	rm -f dist/*
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" dist/wasm_exec.js
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o ./dist/lib.wasm go/main/main.go

browser-test:
	node browser-test/run.js
