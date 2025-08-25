# Repository Guidelines

## Project Structure & Module Organization
- Code: Go packages at the top level. Key folders: `cmd/` (Cobra CLI commands), `provider/` (sources like `mangadex/`, `manganato/`, `manganelo/`, `mangapill/`), `converter/` (pdf, cbz, zip, images), `tui/`, `util/`, `config/`, `history/`, `filesystem/`, `version/`.
- Entry point: `main.go` initializes config/logs and executes the CLI.
- Tests: `_test.go` files live next to code (e.g., `provider/*_test.go`, `converter/*_test.go`).
- Assets & scripts: `assets/` for media, `scripts/` for install/run helpers, `completions/` for shell completions. Docker support via `Dockerfile`.

## Build, Test, and Development Commands
- `make build`: Build the `mangal` binary with ldflags.
- `make install`: Install to `GOBIN` for local use.
- `make test` or `go test ./...`: Run all unit tests.
- `go fmt ./...` and `go vet ./...`: Format and basic static checks.
- Docker: `docker build -t mangal .` then `docker run --rm -ti -e TERM=xterm-256color -v $(PWD)/mangal/downloads:/downloads -v $(PWD)/mangal/config:/config mangal`.

## Coding Style & Naming Conventions
- Language: Go 1.18+. Always run `gofmt` (tabs, standard import grouping).
- Packages: lower-case names; no underscores. Exported identifiers PascalCase; unexported camelCase. Test files: `*_test.go`.
- Commands follow Cobra patterns under `cmd/`. Prefer small, cohesive packages and clear error wrapping with `%w`.

## Testing Guidelines
- Framework: `go test`; some suites use goconvey. Name tests `TestXxx(t *testing.T)` and prefer table-driven cases.
- Run `go test -race -cover ./...` locally. Add/adjust tests when changing behavior; keep coverage stable.

## Commit & Pull Request Guidelines
- Commits: Use Conventional Commits (e.g., `feat:`, `fix:`, `chore:`, `docs:`). Keep messages imperative and scoped.
- PRs: Provide a clear summary, linked issues, and relevant screenshots/terminal output for UX/behavior changes. Update docs/help text when flags or defaults change. Ensure `go fmt ./...` and `make test` pass.

## Security & Configuration Tips
- Configuration: TOML at `mangal where --config` (e.g., `~/.config/mangal/mangal.toml`). Override via `MANGAL_CONFIG_PATH`.
- Secrets: Never commit tokens or API keys. Prefer env vars (`mangal env` shows supported vars).
- Providers: Be mindful of site rate limits and user agents when adding or modifying scrapers.

