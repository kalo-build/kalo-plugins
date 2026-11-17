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
  wasm="$root/dist/wasm/$p.wasm"
  lock=$(grep -E "^$p " "$src/artifacts.lock" 2>/dev/null || true)
  if [ -n "$lock" ]; then
    # Released binary pinned in recipes/<name>/artifacts.lock: download once, fail closed on the digest.
    url=$(echo "$lock" | cut -d' ' -f2); want=$(echo "$lock" | cut -d' ' -f3)
    mkdir -p "$root/dist/wasm"
    if [ "$(sha256sum "$wasm" 2>/dev/null | cut -d' ' -f1)" != "$want" ]; then
      echo "== $p: download $url" >&2
      curl -fsSL --retry 3 -m 300 -o "$wasm.part" "$url"
      got=$(sha256sum "$wasm.part" | cut -d' ' -f1)
      [ "$got" = "$want" ] || { rm -f "$wasm.part"; echo "$p: sha256 $got, artifacts.lock says $want" >&2; exit 1; }
      mv "$wasm.part" "$wasm"
    fi
    echo "$want  dist/wasm/$p.wasm (artifacts.lock)" >&2
  else
    "$root/scripts/ci.sh" "$p" >&2
  fi
  cp "$wasm" "$out/artifacts/$p.wasm"
  sum=$(sha256sum "$wasm" | cut -d' ' -f1)
  grep -rl "@@DIGEST_${p}@@" "$out/descriptors" | while read -r f; do
    sed "s/@@DIGEST_${p}@@/$sum/g" "$f" > "$f.tmp" && mv "$f.tmp" "$f"
  done
done
echo "$out"
