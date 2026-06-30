# Branch Protection

Recommended settings for `main` in GitHub repository settings:

1. Enable **Require a pull request before merging**.
2. Enable **Require approvals** and set required approvals to at least `1`.
3. Enable **Require review from Code Owners**. `.github/CODEOWNERS` assigns all paths to `@stonecharioteer`, so owner review is required before a PR is ready to merge.
4. Enable **Require status checks to pass before merging**.
5. Require these checks:
   - `Go checks`
   - `Check PR title`
   - `Check commit messages`
6. Enable **Require branches to be up to date before merging**.
7. Enable **Require conversation resolution before merging**.
8. Keep **Do not allow bypassing the above settings** enabled for normal contributors.

The weekly/manual integration workflow is intentionally not a required merge check because it uses Docker/Testcontainers and is slower than the standard PR checks.
