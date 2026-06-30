# Contributing

Thanks for contributing to GoForGo!

## Before opening a PR

Please run the pre-commit checks locally. We use `uvx` so you do not need to install `pre-commit` globally:

```bash
uvx pre-commit install --hook-type pre-commit --hook-type commit-msg
uvx pre-commit run --all-files
```

The hooks check formatting, module tidiness, linting, vulnerability scanning, Go tests for buildable packages, CLI build, exercise integrity, conventional commit messages, and disallowed AI attribution footers.

The repository intentionally stores incomplete learner exercises under `exercises/` and standalone reference programs under `solutions/`. CI and pre-commit therefore test the buildable application packages rather than running `go test ./...` across those content directories.

## Commit messages and PR titles

Use Conventional Commits for commit messages and PR titles, for example:

```text
feat: add exercise validation check
fix: handle watcher cleanup errors
ci: add GitHub Actions checks
```

Allowed types are `build`, `chore`, `ci`, `docs`, `feat`, `fix`, `perf`, `refactor`, `revert`, `style`, and `test`.

Do not include `Generated with Claude Code` or `Co-authored-by: Claude` attribution in commit messages.

## Maintainer notes

A PR is ready to merge only after all required checks pass and `@stonecharioteer` has reviewed it. See [docs/branch-protection.md](docs/branch-protection.md) for the recommended branch protection settings for `main`.
