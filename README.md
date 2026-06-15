# Ebitengine2026

Ebitengine Game Jam 2026 prototype.

## Run locally

```sh
go run .
```

## Run in a browser

```sh
go run github.com/hajimehoshi/wasmserve@latest .
```

Then open <http://localhost:8080/>.

## Build GitHub Pages files

```sh
env GOOS=js GOARCH=wasm go build -o docs/main.wasm .
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" docs/wasm_exec.js
```

The files under `docs/` are intended to be served by GitHub Pages.
