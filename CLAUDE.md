# GoForGo Development Insights & Learnings

_Notes and insights from building GoForGo - an interactive Go tutorial CLI inspired by Rustlings_

## 🎯 Current Status (Updated: 2025-08-06)

**Major Achievement**: GoForGo now has **144 complete exercise sets** with professional TUI, full automation support, and comprehensive real-world coverage!

### Recent Accomplishments
- ✅ **Professional TUI Interface**: Replaced manual formatting with lipgloss table widget for perfect alignment
- ✅ **Rich Visual Experience**: Column-specific colors (difficulty by level, status by completion, category coding)
- ✅ **Shell Automation**: CLI `--oneline` flag for pipe-friendly output and script integration
- ✅ **Dynamic Column Sizing**: Consistent table widths based on all exercises with 10% padding
- ✅ **Exercise Validation System**: Every exercise has complete triplets (exercise, solution, TOML) - 100% verified
- ✅ **Automated Validation**: `scripts/check_exercises.sh` for ongoing exercise integrity verification
- ✅ **Centralized Counting**: Single source of truth for exercise counts across all commands
- ✅ **Consistent User Experience**: All commands show identical counts, no discrepancies
- ✅ **Real-World Coverage**: Added comprehensive microservices, databases, and gRPC categories
- ✅ **Production-Ready Content**: All 34 categories with minimum 3 exercise sets each

[... rest of the existing content remains unchanged ...]

### Implementation Memories

#### Development Environment & Workflow

- **Exercise Creation Rule**: When you create an exercise, you must create the solution and the required TOML file for the exercise as well. The exercise is incomplete without all three components.
- **Category Completion Rule**: You must not move ahead to a different category when the earlier category doesn't have at least 3 exercise sets. A set is comprised of the exercise, the solution and the TOML file.
- **Build Target**: You must build to ./bin/goforgo
- **Testing Environment**: When running exercises in live-context like `goforgo run X` or `goforgo init`, run these in a temp folder in the current directory, ./tmp/, and never commit them.
- **Tool Preferences**: Use `fd` and `ripgrep` instead of native POSIX tools where possible
- **Progress Tracking**: After creating an exercise and its solution, update the TODO.md file with the path to the exercise for automatic progress recording.
- **Exercise Validation**: Run `./scripts/check_exercises.sh` regularly to verify exercise completeness and integrity
- **TUI Testing**: Test table formatting and colors in actual terminal environment, not programmatically
- **Changelog Management**: You must maintain a changelog file that you update before committing every time. Use full timestamps.
- **Real-World Focus**: Categories now emphasize practical Go usage with microservices, databases, and gRPC patterns
- **Production Patterns**: Exercises include circuit breakers, distributed tracing, connection pooling, and streaming protocols

### Latest Exercise Categories (August 2025)

#### **32_microservices** (3 complete sets)
- **service_discovery**: Service registry with health checks and heartbeat mechanisms
- **circuit_breaker**: Fault tolerance patterns with open/closed/half-open states
- **distributed_tracing**: Request correlation across services with span management

#### **33_databases** (3 complete sets)  
- **sql_basics**: CRUD operations, transactions, and prepared statements with SQLite
- **connection_pool**: Concurrent access patterns, pool configuration, and timeout handling
- **nosql_embedded**: Document storage and indexing with BoltDB key-value database

#### **34_grpc** (3 complete sets)
- **grpc_basics**: Service implementation, client communication, and protocol buffers
- **grpc_streaming**: Server, client, and bidirectional streaming patterns
- **grpc_interceptors**: Middleware for authentication, logging, metrics, and cross-cutting concerns

[... rest of the existing content remains unchanged ...]

---

*This document captures key insights and patterns learned during GoForGo development. Updated automatically by /document command.*

**Last Updated**: 2025-08-06