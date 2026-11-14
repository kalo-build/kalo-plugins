#!/bin/sh
# Build dist/pack/recipe-<name>-<version>.tar.gz and its .sha256 (sha256sum format).
# Usage: scripts/pack-recipe.sh openapi-sdk
# The tarball holds the recipe folder contents at its root: recipe.yaml, descriptors/,
# protocol/, artifacts/*.wasm. Entries are sorted, owner 0, mtime fixed, so the same
# wasm gives the same tarball. Never uploaded or fetched by this script.
set -eu
root=$(cd "$(dirname "$0")/.." && pwd)
name=${1:?usage: pack-recipe.sh <recipe-name>}
stage=$("$root/scripts/stage-recipe.sh" "$name" | tail -1)
version=$(sed -n 's/^version: *//p' "$root/recipes/$name/recipe.yaml" | head -1)
asset="recipe-$name-$version.tar.gz"
mkdir -p "$root/dist/pack"
(cd "$stage" && find . -type f | sed 's#^\./##' | LC_ALL=C sort |
  tar --owner=0 --group=0 --numeric-owner --mtime='2026-01-01 00:00:00 UTC' -cf - -T - |
  gzip -n -9 > "$root/dist/pack/$asset")
(cd "$root/dist/pack" && sha256sum "$asset" > "$asset.sha256" && cat "$asset.sha256" && ls -l "$asset")
