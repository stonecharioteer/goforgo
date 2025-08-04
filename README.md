# GoForGo 🚀

[![Go Version](https://img.shields.io/badge/Go-1.24%2B-blue.svg)](https://golang.org/dl/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://img.shields.io/badge/Build-Passing-green.svg)](#building-from-source)

**Interactive Go tutorial CLI inspired by Rustlings** 📚

GoForGo helps you learn Go by fixing broken code exercises with real-time feedback. It features 250+ exercises covering Go fundamentals through advanced topics and popular libraries, all with a beautiful terminal interface.

## ✨ Features

- 🎯 **250+ Interactive Exercises** - From basics to advanced Go concepts
- 👁️ **Real-time File Watching** - Automatic compilation and feedback as you code
- 🎨 **Beautiful TUI** - Clean terminal interface with progress tracking
- 📚 **Progressive Learning** - Structured curriculum with difficulty levels
- 💡 **Smart Hints** - Context-aware hints that adapt to your attempts
- 🧪 **Comprehensive Testing** - Built-in Go testing integration
- 🔧 **Modern Go** - Latest Go 1.24+ features and best practices
- 📦 **Popular Libraries** - Exercises with Charm, Gorilla, Gin, and more

## 🚀 Quick Start

### Option 1: Install from Source (Recommended for now)

```bash
# Clone the repository
git clone https://github.com/stonecharioteer/goforgo.git
cd goforgo

# Build and install
just build  # or: go build -o bin/goforgo ./cmd/goforgo

# Initialize exercises in your learning directory  
mkdir ~/learn-go && cd ~/learn-go
~/path/to/goforgo/bin/goforgo init

# Start learning!
~/path/to/goforgo/bin/goforgo
```

### Option 2: Direct Go Install (Coming Soon)

```bash
go install github.com/stonecharioteer/goforgo/cmd/goforgo@latest
```

## 🎮 Usage

### Initialize Your Learning Environment

```bash
goforgo init
```

This creates:
- `exercises/` - 250+ Go exercises organized by topic
- `solutions/` - Complete reference solutions
- `.goforgo.toml` - Your progress and preferences

### Start Interactive Mode (Default)

```bash
goforgo
```

This launches the interactive mode with:
- ⚡ Real-time file watching and compilation
- 📊 Progress tracking and visual feedback  
- ⌨️ Keyboard shortcuts for navigation
- 💡 Progressive hints and guidance

### Run Specific Exercises

```bash
goforgo run hello              # Run the 'hello' exercise
goforgo run                    # Run next incomplete exercise
goforgo hint variables1        # Show hint for specific exercise
goforgo list                   # List all exercises with progress
goforgo list --all             # Show completed exercises too
```

### Available Commands

| Command | Description |
|---------|-------------|
| `goforgo` | Start interactive watch mode (default) |
| `goforgo init` | Initialize exercises in current directory |
| `goforgo run [exercise]` | Run specific exercise or next incomplete |
| `goforgo hint [exercise]` | Show progressive hints |
| `goforgo list [--all] [--category=...]` | List exercises with filters |
| `goforgo watch` | Explicit watch mode with file monitoring |

## 🏗️ Building from Source

### Prerequisites

- **Go 1.24+** (required for latest language features)
- **Just** (recommended) - Install from [casey/just](https://github.com/casey/just)
- **Git** for version information

### Development Setup

```bash
# Clone repository
git clone https://github.com/stonecharioteer/goforgo.git
cd goforgo

# Install development dependencies
just install-deps

# Build for development  
just dev-build

# Run tests
just test

# Build optimized release binary
just build
```

### Available Just Commands

```bash
just --list                    # Show all available commands

# Building
just build                     # Build optimized binary
just dev-build                 # Fast development build
just build-race                # Build with race detection
just build-release             # Cross-platform release binaries

# Testing & Quality
just test                      # Run all tests
just test-coverage             # Generate coverage report
just bench                     # Run benchmarks
just lint                      # Lint code with golangci-lint
just fmt                       # Format code

# Development
just dev-run                   # Build and test CLI in demo mode
just test-cli                  # Test CLI functionality
just watch                     # Auto-rebuild on changes (requires entr)
just pre-commit                # Full check before committing

# Maintenance
just clean                     # Clean build artifacts
just tidy                      # Tidy Go modules
just info                      # Show project information
```

### Manual Build (without Just)

```bash
# Basic build
mkdir -p bin
go build -o bin/goforgo ./cmd/goforgo

# With version information
go build -ldflags="-X 'github.com/stonecharioteer/goforgo/internal/cli.version=v1.0.0'" -o bin/goforgo ./cmd/goforgo

# Cross-platform builds
GOOS=linux GOARCH=amd64 go build -o dist/goforgo-linux-amd64 ./cmd/goforgo
GOOS=darwin GOARCH=amd64 go build -o dist/goforgo-darwin-amd64 ./cmd/goforgo
GOOS=windows GOARCH=amd64 go build -o dist/goforgo-windows-amd64.exe ./cmd/goforgo
```

## 📚 Exercise Categories

GoForGo includes exercises in these categories:

### Core Go (150+ exercises)
- **01_basics** - Hello world, syntax, comments
- **02_variables** - Types, declarations, zero values  
- **03_functions** - Parameters, returns, methods
- **04_control_flow** - if/else, loops, switch
- **05_data_structures** - Arrays, slices, maps
- **06_structs** - Definition, embedding, methods
- **07_interfaces** - Types, satisfaction, composition
- **08_pointers** - Memory addresses, performance
- **09_error_handling** - Patterns, wrapping, custom errors
- **10_packages** - Modules, imports, visibility

### Concurrency (50+ exercises)  
- **11_goroutines** - Basic concurrency, race conditions
- **12_channels** - Communication, buffering, patterns
- **13_select** - Multiplexing, timeouts, non-blocking
- **14_sync** - Mutex, WaitGroup, atomic operations
- **15_context** - Request scoping, cancellation

### Modern Go (40+ exercises)
- **16_generics** - Type parameters, constraints, inference
- **17_testing** - Units, benchmarks, examples, fuzzing
- **18_reflection** - Type inspection, dynamic calls
- **19_json** - Encoding, decoding, streaming
- **20_http** - Servers, clients, middleware

### Popular Libraries (50+ exercises)
- **21_bubbletea** - TUI applications, Elm architecture
- **22_cobra** - CLI applications, commands, flags
- **23_gin** - Web APIs, middleware, JSON binding
- **24_gorilla** - Advanced HTTP routing and sessions
- **25_gorm** - ORM, migrations, associations

## 🎯 Learning Path

1. **🌱 Beginner** (1-50): Syntax, variables, functions, control flow
2. **🌿 Intermediate** (51-120): Data structures, interfaces, error handling  
3. **🌳 Advanced** (121-200): Concurrency, generics, reflection
4. **🚀 Expert** (201-250): Performance, libraries, real-world projects

Each exercise includes:
- 📝 Clear learning objectives
- ⭐ Difficulty rating (1-5 stars)
- ⏱️ Estimated completion time
- 💡 Progressive hints (3 levels)
- ✅ Automatic validation

## 🤝 Contributing

We welcome contributions! See our [contribution guidelines](CONTRIBUTING.md) for:

- 🐛 **Bug Reports** - Found an issue? Let us know!
- ✨ **Feature Requests** - Ideas for improvements
- 📚 **Exercise Contributions** - Add new exercises
- 🔧 **Code Improvements** - Performance, usability, tests

### Development Workflow

```bash
# 1. Fork and clone
git clone https://github.com/yourusername/goforgo.git
cd goforgo

# 2. Install dependencies  
just install-deps

# 3. Make changes and test
just pre-commit

# 4. Submit PR with clear description
```

## 📖 Documentation

- 📋 [**TODO.md**](TODO.md) - Development roadmap and tasks
- 🎯 [**GAMEPLAN.md**](GAMEPLAN.md) - Project architecture and phases  
- 🧠 [**CLAUDE.md**](CLAUDE.md) - Development insights and learnings
- 📚 [**docs/**](docs/) - Additional documentation

## 🙏 Acknowledgments

- **Rustlings** - Original inspiration for interactive learning
- **Go Team** - Amazing language and tooling
- **Charm** - Beautiful TUI libraries (Bubble Tea, Lip Gloss)
- **Community** - All the Go learning resources and examples

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

---

**Happy Learning! 🎉** 

*Start your Go journey today with `goforgo init`*
