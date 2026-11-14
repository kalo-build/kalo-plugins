#!/bin/sh
# Stage a recipe into dist/recipes/<name>/ ready to run or pack:
#   recipes/<name>/ + artifacts/<plugin>.wasm (built from plugins/) + protocol/v0.1-alpha/schemas
# and fill each @@DIGEST_<plugin>@@ placeholder with the sha256 of the wasm just built.
# Usage: scripts/stage-recipe.sh openapi-sdk      (prints the staged folder)
set -eu
root=$(cd "$(dirname "$0")/.." && pwd)
name=${1:?usage: stage-recipe.sh <recipe-name>}
src="$root/recipes/$name"
[ -f "$src/recipe.yaml" ] || { echo "no recipe $name" >&2; exit 1; }
out="$root/dist/recipes/$name"
rm -rf "$out"
mkdir -p "$out/artifacts" "$out/protocol/v0.1-alpha"
cp -R "$src/." "$out/"
cp -R "$root/protocol/v0.1-alpha/schemas" "$out/protocol/v0.1-alpha/schemas"
plugins=$(grep -rhoE '@@DIGEST_[a-z0-9-]+@@' "$src/descriptors" | sed 's/@@DIGEST_//; s/@@//' | sort -u)
for p in $plugins; do
  "$root/scripts/ci.sh" "$p" >&2
  wasm="$root/dist/wasm/$p.wasm"
  cp "$wasm" "$out/artifacts/$p.wasm"
  sum=$(sha256sum "$wasm" | cut -d' ' -f1)
  grep -rl "@@DIGEST_${p}@@" "$out/descriptors" | while read -r f; do
    sed "s/@@DIGEST_${p}@@/$sum/g" "$f" > "$f.tmp" && mv "$f.tmp" "$f"
  done
done
echo "$out"
