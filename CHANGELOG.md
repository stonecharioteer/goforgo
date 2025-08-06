# Changelog

All notable changes to GoForGo will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased] - 2025-08-06T11:19:00+05:30

### Added
- **Exercise Validation Infrastructure**: Created comprehensive validation script `scripts/check_exercises.sh`
  - Automated checking of all exercise triplets (exercise + solution + TOML)
  - Category-wise breakdown and completion statistics
  - Detection of missing components and orphaned files
  - Verified 122 complete exercise sets with 100% completion rate
- **Professional TUI Interface**: Replaced manual table formatting with lipgloss table widget
  - Automatic column alignment and consistent spacing
  - Rich column-specific colors (difficulty by level, status by completion)
  - Dynamic column sizing based on all exercises with 10% padding
  - Perfect alignment that prevents table resizing during scrolling
- **Shell Integration & Automation**: Added CLI `--oneline` flag for machine-readable output
  - Pipe-friendly format: `name|category|difficulty|status|title|time`
  - Perfect for shell scripts and automation tools
  - Maintains backward compatibility with existing CLI formatting
- **Dark Terminal Optimization**: Improved color choices for better visibility
  - Light text colors for excellent contrast against dark backgrounds
  - Color-coded difficulty levels (green→blue→orange→red→purple)
  - Proper selection highlighting with purple theme

### Fixed
- **Difficulty Display Issues**: Fixed "unknown" difficulty display in TUI
  - Direct TOML mapping for reliable difficulty values
  - Fallback logic for edge cases
  - All 122 exercises now show proper difficulty levels
- **Column Width Consistency**: Eliminated table resizing during scrolling
  - Column widths calculated from ALL exercises, not just visible ones
  - Dynamic sizing with 10% padding for optimal readability
  - Consistent table layout throughout navigation

### Changed
- **Exercise Count**: Validated and confirmed 122 complete exercise sets (up from 121+)
- **TUI Architecture**: Migrated from manual formatting to professional table widget
- **Documentation**: Updated CLAUDE.md, TODO.md, and GAMEPLAN.md with latest achievements
  - Reflects current status with 122 exercises and 100% validation rate
  - Documents new TUI capabilities and shell automation features
  - Added Phase 3.5 completion for TUI enhancements

## [0.3.0] - 2025-08-05T15:30:00+05:30

### Added
- **Complete Exercise Validation System**: Achieved 121+ complete exercise sets
  - Created 20 missing TOML metadata files
  - Created 2 missing solution files  
  - Created 3 missing exercise files for existing solutions
  - Established three-component rule: every exercise needs .go + .toml + solution
- **Centralized Counting Architecture**: Single source of truth for exercise counts
  - Added GetTotalExerciseCount(), GetCompletedExerciseCount(), GetProgressStats()
  - Updated TUI to use centralized methods instead of local counting
  - Updated CLI list command to use ExerciseManager.GetProgressStats()
  - Fixed discrepancy between init (121) and list (101) commands

### Fixed
- **Exercise Component Integrity**: Ensured all exercises have complete triplets
- **Counting Consistency**: Eliminated discrepancies between commands
- **Dynamic Exercise Loading**: Made system directory-agnostic

## [0.2.0] - 2025-08-04T12:00:00+05:30

### Added
- **Professional UI/UX Polish**: Production-ready interface with animations
  - Animated splash screen with 8-frame logo animation
  - Color-cycling startup sequence with loading dots
  - Beautiful ASCII art transitions and smooth timing
- **Uniform Visual Styling**: Consistent decorative borders across all TUI pages
- **Enhanced User Experience**: Progressive hints system and smart progress tracking

### Fixed
- **Border Width Calculations**: Prevented border cutoff issues
- **Progress Display**: Shows accurate completion percentage
- **Responsive Design**: Adapts to different terminal sizes

## [0.1.0] - 2025-08-03T10:00:00+05:30

### Added
- **Core CLI Framework**: Complete command structure with Cobra
  - `init`, `run`, `watch`, `hint`, `list`, `reset` subcommands
  - Basic argument parsing and validation
- **Exercise Management System**: TOML-based exercise configuration
  - Exercise metadata structure (name, category, difficulty, hints)
  - Exercise loading and validation logic
  - Progress tracking and state management
- **Bubble Tea TUI Interface**: Interactive terminal interface
  - Progress bar with current exercise info
  - Real-time compilation feedback
  - Navigation between exercises
- **File Watching System**: fsnotify integration for Go file changes
  - Debounced compilation triggers
  - Smart filtering (ignore temp files, build artifacts)

### Technical
- **Project Structure**: Established internal/ layout for Go project
- **Dependencies**: Added Cobra, Bubble Tea, fsnotify, TOML parser
- **Build System**: Set up build target to ./bin/goforgo