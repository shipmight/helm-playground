
test:
	go test go/lib/lib.go go/lib/lib_test.go

build:
	mkdir -p dist
	rm -f dist/*
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" dist/wasm_exec.js
	GOOS=js GOARCH=wasm go build -o ./dist/lib.wasm go/main/main.go

browser-test:
	echo hi
