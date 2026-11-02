# Untrack build output: bin/plugin.wasm and dist/ (ut-docs#3599)

## What shipped
- `.gitignore` now ignores `/bin/` and `/dist/`, the same as the sibling plugins.
- `bin/plugin.wasm` and four stale `dist/` files (the 1.0.0 and 1.1.0 tarballs and their `.sha256`) are no longer tracked. A local `scripts/build.sh` or `package.sh` no longer dirties the tree, and a locally built binary can't slip into a commit.
- `scripts/check-version-bump.sh` is now byte-identical with ut-plugin-payment-demo's. It drops `'bin/*'` from `SHIPPED_PATTERNS` and the old "this repo commits bin/" header. `check-version-bump.test.sh` matches payment-demo's except for the plugin id: case 12b is gone, and the test fixture now gitignores `/bin/ /dist/`. `check-version-bump.body.sh` is unchanged.
- There is no version bump. Once `bin/*` is out of the patterns, this PR touches no shipped file. CI and release.yml run `scripts/build.sh` before `package.sh`, so the release bundle is built exactly as before.

## Review
The review was independent and used a different model from the author. It ran in a fresh clone of the change.
- These checks passed: `GOOS=wasip1 GOARCH=wasm go vet ./...`, `gofmt -l .` (empty), `build.sh`, `validate.sh` and `package.sh`. The tarball still holds `manifest.json`, `assets/`, `README.md`, `bin/plugin.wasm` and `LICENSE`.
- `check-version-bump.test.sh` passed every case. The "entries mirror" case now takes its gitignored-build-artifact branch.
- The guard ran on the commit and reported "no shipped file changed".
- `git status` is clean after build and package.
- Nothing in the README, the workflows or `publish.sh`/`approve.sh` assumes tracked build output.
- Dropping `bin/*` loses no coverage. An untracked `bin/` can never appear in a diff, and its inputs (src/, go.mod, go.sum, build.sh) are still shipped patterns.

Findings: two nits, both accepted with no change.
- `docs/code-reviews/2026-10-28-version-bump-guard-4917.md` still says `bin/` is committed. It is a dated record, so it is left as it was; this record notes the change.
- `check-version-bump.sh` still mentions the compiled `bin/plugin.wasm`. That is accurate.

Not verified locally: the real release workflow. CI's `validate` job runs the same build, validate and package steps on a clean checkout.

## Verdict
Safe to merge.
