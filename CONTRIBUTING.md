# Contributing

Thanks for contributing to GoForGo!

## Before opening a PR

Please run the pre-commit checks locally. We use `uvx` so you do not need to install `pre-commit` globally:

```bash
uvx pre-commit install --hook-type pre-commit --hook-type commit-msg
uvx pre-commit run --all-files
```

The hooks check formatting, module tidiness, linting, focused Go tests, CLI build, exercise integrity, and conventional commit messages.

## Commit messages and PR titles

Use Conventional Commits for commit messages and PR titles, for example:

```text
feat: add exercise validation check
fix: handle watcher cleanup errors
ci: add GitHub Actions checks
```

Allowed types are `build`, `chore`, `ci`, `docs`, `feat`, `fix`, `perf`, `refactor`, `revert`, `style`, and `test`.
