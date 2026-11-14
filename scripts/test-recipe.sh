#!/bin/sh
# Run a recipe on its fixture through the real host and diff against the golden output.
# Usage: KALO=/path/to/kalo scripts/test-recipe.sh openapi-sdk
# KALO must be a kalo-cli build from the commit in scripts/kalo-cli-ref (needs spec run --recipe-dir).
set -eu
root=$(cd "$(dirname "$0")/.." && pwd)
name=${1:?usage: test-recipe.sh <recipe-name>}
kalo=${KALO:-kalo}
staged=$("$root/scripts/stage-recipe.sh" "$name" 2>/dev/null | tail -1)
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT
"$kalo" spec run --recipe-dir "$staged" --input "$root/recipes/$name/testdata/openapi.yaml" --out "$tmp/out" >/dev/null
diff -r "$root/recipes/$name/testdata/golden" "$tmp/out" && echo "recipe $name: golden OK"
