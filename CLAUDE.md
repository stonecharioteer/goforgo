# GoForGo Development Insights & Learnings

_Notes and insights from building GoForGo - an interactive Go tutorial CLI inspired by Rustlings_

[... existing content remains unchanged ...]

## Implementation Memories

### Development Environment & Workflow

- **When running the exercises in a live-context, like trying to run `goforgo run X` or `goforgo init` you must run these in a temp folder in the current directory, ./tmp/, and never commit them.**
- **You must build to ./bin/goforgo,**
- Use `fd` and `ripgrep` instead of the native posix tools where possible
- After creating an exercise and its solution, update the TODO.md file with the path to the exercise so that your progress is recorded automatically.
- **You must add at least 3 exercises per category**

[... rest of existing content remains unchanged ...]