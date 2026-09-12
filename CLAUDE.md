# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

`qx` (Quick Execute) is a small Go CLI that turns a natural-language query into a shell command via the Groq API, then copies the result to the clipboard (`pbcopy`). It also installs an optional `qq()` shell function (zsh/bash) that prefills the terminal command line with the generated command.

## Commands

```bash
make build      # go build -> bin/qx
make install    # build + cp bin/qx /usr/local/bin/qx
make dev        # go run ./cmd/qx/main.go
make test       # go test -v ./...
make deps       # go mod download && go mod tidy
make clean      # rm -rf bin/ && go clean
```

Run a single test: `go test -v ./internal/cmd/ -run TestName`

Local manual testing requires `QX_GROQ_KEY` to be set in the environment (see `qx set-key` below) — without it, `callGroqAPI` returns an onboarding error message instead of calling the API.

## Architecture

- `cmd/qx/main.go` — entrypoint, just calls `cmd.Execute()`.
- `internal/cmd/root.go` — **custom arg-dispatch, not standard Cobra routing.** `Execute()` inspects `os.Args` directly before touching Cobra: it special-cases `--version`/`-v`, `--raw`/`-r`, and otherwise treats any first argument that isn't a known subcommand (`set-key`, `qq-install`, `qq-uninstall`, `help`, `-h`, `--help`) as a natural-language query and routes straight to `processQuery`/`processQueryWithRaw`, bypassing `rootCmd.Execute()` entirely. When adding a new subcommand, you must add its name to the exclusion check in `Execute()` (root.go) or it will be swallowed as a query string.
- `internal/cmd/groq.go` — `callGroqAPI` is the only integration point with Groq (`openai/gpt-oss-120b`, chat completions endpoint). The system prompt instructs the model to return only a raw shell command with no formatting.
- `internal/cmd/set-api-key.go`, `qq-install.go`, `qq-uninstall.go` — all directly edit the user's shell rc file (`~/.zshrc` or `~/.bashrc`, chosen via the `SHELL` env var) as plain text: appending/replacing an `export QX_GROQ_KEY=...` line, or inserting/removing a `qq()` function between `# >>> qx tool >>>` / `# <<< qx tool <<<` markers. Any change to the qq() function body must keep those exact markers since `qq-uninstall` matches on them verbatim.
- Every subcommand registers itself onto the shared `rootCmd` via an `init()` calling `rootCmd.AddCommand(...)`.
- Version is injected at build/release time via ldflags into `internal/cmd.Version` (see `Makefile` and `.goreleaser.yml`).

## Release process

Push a `v*` tag (e.g. `v0.1.13`) to `main` to cut a release. `.github/workflows/release.yml` runs GoReleaser (`goreleaser/goreleaser-action`) against `.goreleaser.yml`, which builds darwin/linux/windows × amd64/arm64 binaries, creates the GitHub Release, and pushes an updated formula to the `Samuel787/homebrew-tap` Homebrew tap in one step. This requires a `HOMEBREW_TAP_TOKEN` repo secret — a PAT with write access to `Samuel787/homebrew-tap` — since the default `GITHUB_TOKEN` can't push to a different repo. `.github/workflows/build.yml` separately runs `make build`/`make test` on push/PR across macOS/Linux/Windows.
