# Changelog

All notable changes to GoForGo will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- **GitHub checks and local pre-commit hooks** *(2026-06-30 10:26:54 IST)*: Added CI for formatting, module tidiness, linting, focused tests, CLI build, exercise integrity, and conventional PR titles, plus matching `uvx pre-commit` setup documentation.
- **Repository quality automation** *(2026-06-30 11:50:14 IST)*: Added Dependabot, advisory govulncheck scanning, PR/issue templates, release builds, branch-protection guidance, CODEOWNERS review requirements, explicit golangci-lint configuration, conventional commit message checks, and Claude attribution rejection in commit-msg hooks.
- **Go exercise integrity checker** *(2026-06-30 11:50:14 IST)*: Replaced the shell implementation with a Go validator that checks triplets, orphaned solutions, TOML metadata, category minimums, and duplicate exercise ordering.

### Fixed
- **Lint compliance** *(2026-06-30 10:26:54 IST)*: Cleaned up lint findings in command output handling, resource cleanup, TUI key handling, and unused code so the new lint check passes.
- **Portable exercise checker** *(2026-06-30 10:26:54 IST)*: Updated `scripts/check_exercises.sh` to run on older Bash versions without associative arrays or `bc`.
- **Integration test gating** *(2026-06-30 11:50:14 IST)*: Moved Testcontainers-dependent validation coverage behind the `integration` build tag and added a separate manual/weekly workflow.
- **Generic exercise ordering** *(2026-06-30 11:50:14 IST)*: Fixed a duplicate order value in the generics category surfaced by the stricter exercise validator.

## [0.9.4] - 2026-04-07

### Changed
- **Daily update-check caching** *(2026-04-07 12:05:00 local)*: GitHub tag checks are now cached for 24 hours in the user cache directory to avoid hitting the API on every command run while still surfacing updates promptly.

## [0.9.3] - 2026-04-07

### Added
- **TUI update notice** *(2026-04-07 11:45:00 local)*: watch mode now shows update availability in the footer so users get upgrade guidance inside the interactive interface too.
- **`goforgo self-update` command** *(2026-04-07 11:45:00 local)*: Added a self-update command that checks tag versions and runs `go install ...@latest` after user confirmation (`--yes` to skip prompt, `--check` for check-only).

## [0.9.2] - 2026-04-07

### Added
- **Version update check via GitHub tags** *(2026-04-07 11:35:00 local)*: CLI now checks the latest repository tag at startup and shows a non-blocking update notice when a newer semantic version is available, including a `go install ...@latest` command.
- **Update-check opt-out flag** *(2026-04-07 11:35:00 local)*: Added `--no-update-check` global flag to skip network version checks.

## [0.9.1] - 2026-04-07

### Fixed
- **Memory creep in watch mode** *(2026-04-07 11:15:00 local)*: eliminated duplicate in-flight file watcher wait commands in TUI to prevent goroutine buildup during long sessions.
- **Watcher shutdown race** *(2026-04-07 11:15:00 local)*: reworked watcher lifecycle with coordinated shutdown (`done` channel + `WaitGroup` + `sync.Once`) to avoid send-on-closed-channel/data-race behavior.
- **Exercise reload growth** *(2026-04-07 11:15:00 local)*: `ExerciseManager.LoadExercises()` now clears the in-memory slice before rescanning to prevent duplicate accumulation.
- **Runner output memory pressure** *(2026-04-07 11:15:00 local)*: capped stdout/stderr capture buffers with truncation notices to prevent unbounded output allocations.
- **Validation concurrency safety** *(2026-04-07 11:15:00 local)*: protected shared validation maps in orchestrator parallel paths to avoid concurrent write races.

## [0.9.0] - 2026-03-03

### Added
- **`goforgo solve` command**: Copy reference solutions over exercises for a range (`solve 5` or `solve 3-7`), marking them complete. Gated behind `dev` build tag — only available via `just dev-build`.
- **`goforgo sync` command**: Re-validate every exercise and update progress to match actual state. Passes get marked complete; failures get unmarked. Includes a styled lipgloss progress bar with green fill, percentage, and current exercise name.
- **`goforgo clean` command**: Remove build artifacts (`go.mod`, `go.sum`, compiled binaries) that accumulate in exercise directories after running exercises.
- **TUI sync via `r` key**: Pressing `r` in the list view triggers a full sync — validates all exercises, updates progress, and refreshes the list in-place. Footer updated with `r=sync` hint.
- **`ExerciseManager.UnmarkExerciseCompleted()`**: Inverse of `MarkExerciseCompleted()` — removes completion status and persists to disk. Used by sync to un-complete exercises that no longer pass.
- **Auto-advance mode**: Press `a` in the main view to toggle. On successful exercise completion, shows a brief crossfade success screen then automatically moves to the next exercise after 1 second.
- **Embedded exercises**: All exercises and solutions are now embedded in the binary via `go:embed`. `go install` works out of the box — no repo clone needed.
- **`goforgo update` command**: Sync local exercises with the embedded content, preserving completed solutions while adding new exercises and removing stale ones.
- **Color-coded file paths**: Exercise detail view shows category in blue and filename in orange for quick visual parsing.
- **Vim-style navigation in list view**: `{count}j/k` motions, `gg`/`G` to jump to top/bottom, `H/M/L` screen positioning, `Ctrl+u/d` half-page scroll.
- **GitHub Dark color scheme**: Full palette overhaul optimized for dark terminals — purple headers, green success, red errors, blue links, orange hints.

### Fixed
- **Broken table layout**: Fixed column alignment, width calculation, and border overhead in the list view. Columns now size based on all exercises, not just visible ones.
- **Exercise ordering**: Exercises sort correctly by category + order within category. Swapped control flow (ch 3) before functions (ch 4) since function exercises depend on `if` statements.
- **Stale exercise cleanup**: `goforgo update` now removes exercises that no longer exist in the embedded content instead of only adding new ones.
- **Resilient update**: Update command handles broken exercise directory structures gracefully instead of failing.
- **Preserved solutions on update**: Completed exercise files are no longer overwritten when directory structure changes.
- **Init command**: Removed hardcoded stale directory list that was creating invalid exercise folders.
- **go.mod version**: Runner now generates correct Go version in auto-created go.mod files.
- **Method exercises**: Moved method exercises out of `03_functions` into `08_structs` where they belong.
- **Beginner hints**: Improved hint text to avoid prematurely revealing solutions.

### Changed
- **Exercise count**: 230 exercises across 57 categories (was 184).
- **Curriculum order**: Control flow now precedes functions to match pedagogical dependencies.
- **Dev build tag**: `just dev-build` now passes `-tags dev`, enabling dev-only commands like `solve`.

## [0.8.0] - 2025-08-20

### Added
- **`goforgo update` command**: Update local exercise files from the embedded binary content. Preserves progress while refreshing exercise content.
- **Run mode default**: Changed all exercises to use `run` validation mode for consistent behavior.

## [0.7.0] - 2025-08-07

### Fixed
- **Table width bug**: Fixed table rendering issue that caused misalignment at certain terminal widths.
- **Table alignment**: Fixed column alignment and width calculation for the exercise list.
- Updated preview image/demo recording.

## [0.6.0] - 2025-08-07

### Added
- **Phase 5 third-party libraries**: Complete integration with 12 major Go library categories.
  - Gorilla Mux, Cobra CLI, Bubble Tea TUI, advanced concurrency (golang.org/x/sync)
  - GORM database, Gin web framework, Logrus logging
  - Kafka message streaming, Kubernetes client-go (5 exercises)
  - Hadoop HDFS/MapReduce/YARN, Spark DataFrames/streaming, Elasticsearch
- **38 new exercises** covering the Go ecosystem with production-ready patterns.
- Total exercise count reached **184 complete sets**.

## [0.5.0] - 2025-08-06

### Added
- **Universal validation system**: TestOrchestrator with 7 pluggable validation rules (HTTP, database, process, network, concurrency, metrics, log).
- **Testcontainers integration**: PostgreSQL and Redis containers for realistic exercise validation.
- **ServiceRegistry**: Manages lifecycle of supporting services with health checking and automatic cleanup.
- **Enhanced TOML configuration**: Extended validation section supporting services and rules specifications, backward compatible with legacy modes.

## [0.4.0] - 2025-08-06

### Added
- **Comprehensive real-world patterns**: Microservices (service discovery, circuit breakers, distributed tracing), databases (SQL, connection pooling, NoSQL), gRPC (basics, streaming, interceptors).
- **144+ complete exercises** across 34 categories.
- Expanded to 133 exercises with additional standard library, crypto, networking, encoding, I/O, path, OS, math, sorting, data structures, algorithms, and web categories.

## [0.3.5] - 2025-08-06

### Added
- **Professional TUI table**: Replaced manual formatting with lipgloss table widget for perfect alignment.
- **Column-specific colors**: Difficulty colored by level (green/blue/orange/red/purple), status by completion.
- **Dynamic column sizing**: Consistent widths based on all exercises with 10% padding.
- **CLI `--oneline` flag**: Pipe-friendly output format for shell scripts and automation.

### Fixed
- **Difficulty display**: Fixed "unknown" difficulty in TUI — direct TOML mapping with fallback logic.
- **Column width consistency**: Eliminated table resizing during scrolling.

## [0.3.0] - 2025-08-05

### Added
- **Complete exercise validation system**: 121+ complete exercise sets verified.
  - Created 20 missing TOML metadata files, 2 missing solution files, 3 missing exercise files.
  - Established three-component rule: every exercise needs `.go` + `.toml` + solution.
- **Centralized counting**: Single source of truth via `GetTotalExerciseCount()`, `GetCompletedExerciseCount()`, `GetProgressStats()`.

### Fixed
- **Counting consistency**: Fixed discrepancy between init (121) and list (101) commands.

## [0.2.0] - 2025-08-04

### Added
- **Animated splash screen**: 8-frame logo animation with color-cycling startup sequence.
- **Uniform visual styling**: Consistent decorative borders across all TUI pages.
- **Progressive hints**: 3-level hint system that adapts to attempt count.
- **TODO comment validation**: Exercises must have all TODO comments resolved to pass.

### Fixed
- **Border width calculations**: Prevented border cutoff issues.
- **Progress display**: Accurate completion percentage.
- **Responsive design**: Adapts to different terminal sizes.

## [0.1.0] - 2025-08-04

### Added
- **Core CLI framework**: Cobra-based command structure with `init`, `run`, `watch`, `hint`, `list`, `reset` subcommands.
- **Exercise management system**: TOML-based exercise configuration with metadata, hints, and validation modes.
- **Bubble Tea TUI**: Interactive terminal interface with progress bar, real-time compilation feedback, and exercise navigation.
- **File watching**: fsnotify integration with debounced compilation triggers and smart filtering.
- **Progress tracking**: Persistent `.goforgo-progress.toml` with completion state and current exercise tracking.
- **10 exercises** in `01_basics` and `02_variables` categories.

### Technical
- Established `internal/` package layout (cli, exercise, runner, tui, watcher).
- Dependencies: Cobra, Bubble Tea, Lip Gloss, fsnotify, BurntSushi/toml.
- Build system with Justfile targeting `./bin/goforgo`.
