COMMIT_HASH=`git rev-parse HEAD`

.PHONY: wasm itchio-wasm


wasm:
	GOARCH=wasm GOOS=js go build -ldflags="-s -w" -trimpath -o _web/main.wasm ./cmd/game

itchio-wasm: wasm
	cd _web && \
		mkdir -p ../bin && \
		rm -f ../bin/tunefire.zip && \
		zip ../bin/tunefire.zip -r main.wasm index.html wasm_exec.js
