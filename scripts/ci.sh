#!/bin/sh
# Build every plugin to WASM and run every test. Same script locally and in CI.
# Usage: scripts/ci.sh            (all plugins)
#        scripts/ci.sh openapi-sdk-ts
set -eu
root=$(cd "$(dirname "$0")/.." && pwd)
mkdir -p "$root/dist/wasm"
# The TypeScript golden tests run tsc --strict on the output and fail in CI without it.
if [ -n "${CI:-}" ] && ! command -v tsc >/dev/null 2>&1; then
  echo "== installing typescript@5.9.3 (Apache-2.0) for the tsc --strict checks"
  npm install -g --no-fund --no-audit typescript@5.9.3 >/dev/null
fi
plugins=${1:-$(cd "$root/plugins" && ls -d */ | tr -d /)}
for p in $plugins; do
  dir="$root/plugins/$p"
  [ -f "$dir/go.mod" ] || { echo "skip $p: no go.mod"; continue; }
  echo "== $p: go vet + test"
  (cd "$dir" && go vet ./... && go test ./...)
  echo "== $p: wasm"
  (cd "$dir" && GOOS=wasip1 GOARCH=wasm go build -buildvcs=false -trimpath \
     -ldflags='-s -w -buildid=' -o "$root/dist/wasm/$p.wasm" ./cmd/plugin)
  sha256sum "$root/dist/wasm/$p.wasm" | sed "s#$root/##"
done
