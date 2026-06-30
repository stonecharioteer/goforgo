package cli

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	goforgo "github.com/stonecharioteer/goforgo"
	"github.com/stonecharioteer/goforgo/internal/exercise"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update exercise source files without overwriting completed work",
	Long: `Update exercise source files (like .toml configuration files) from the binary
without overwriting Go files for exercises you've already completed. This allows you
to get updates to exercise metadata while preserving your solutions.

This command will:
- Update all .toml configuration files
- Update .go files only for incomplete exercises
- Preserve your completed exercise solutions
- Keep your progress intact`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := GetWorkingDirectory()
		if err != nil {
			return fmt.Errorf("failed to get working directory: %w", err)
		}

		// Check if this looks like a goforgo directory
		if _, err := os.Stat(filepath.Join(cwd, ".goforgo.toml")); err != nil {
			return fmt.Errorf("this doesn't appear to be a GoForGo directory (no .goforgo.toml found)")
		}

		fmt.Println("🔄 Updating GoForGo exercise files...")

		// Load exercise manager to get completion status.
		// Progress is loaded from .goforgo-progress.toml independently of
		// exercise scanning, so completed exercises are available even when
		// the directory structure has changed between versions.
		em := exercise.NewExerciseManager(cwd)
		if err := em.LoadExercises(); err != nil {
			fmt.Println("⚠️  Could not load current exercises (structure may have changed)")
			fmt.Println("   Completed exercises will still be preserved from progress file.")
		}
		completedExercises := em.GetCompletedExercises()
		fmt.Printf("📊 Found %d completed exercises to preserve\n", len(completedExercises))

		// Update exercise files selectively
		if err := updateExerciseFiles(cwd, completedExercises); err != nil {
			return fmt.Errorf("failed to update exercise files: %w", err)
		}

		fmt.Println("✅ GoForGo exercises updated successfully!")
		fmt.Println("💡 Your completed exercise solutions have been preserved.")
		return nil
	},
}

func updateExerciseFiles(baseDir string, completedExercises map[string]bool) error {
	fmt.Println("📂 Updating from embedded exercises...")

	// Walk through embedded exercises and update selectively
	err := fs.WalkDir(goforgo.Content, "exercises", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel("exercises", path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(baseDir, "exercises", relPath)

		if d.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		// Determine if we should update this file
		shouldUpdate := shouldUpdateFile("", destPath, relPath, completedExercises)

		if shouldUpdate {
			fmt.Printf("  📝 Updating: %s\n", relPath)
			content, readErr := goforgo.Content.ReadFile(path)
			if readErr != nil {
				return readErr
			}
			return os.WriteFile(destPath, content, 0644)
		} else {
			fmt.Printf("  ⏭️  Preserving: %s (exercise completed)\n", relPath)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to update exercises: %w", err)
	}

	// Update solutions directory too (these are reference solutions)
	fmt.Println("📂 Updating solutions...")
	err = fs.WalkDir(goforgo.Content, "solutions", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel("solutions", path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(baseDir, "solutions", relPath)

		if d.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		fmt.Printf("  📝 Updating solution: %s\n", relPath)
		content, readErr := goforgo.Content.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		return os.WriteFile(destPath, content, 0644)
	})
	if err != nil {
		return fmt.Errorf("failed to update solutions: %w", err)
	}

	// Remove stale files that no longer exist in embedded content
	fmt.Println("🧹 Cleaning up removed exercises...")
	removed, err := removeStaleFiles(baseDir)
	if err != nil {
		return fmt.Errorf("failed to clean up stale files: %w", err)
	}
	if removed > 0 {
		fmt.Printf("  🗑️  Removed %d stale files\n", removed)
	}

	return nil
}

// collectEmbeddedPaths builds a set of all relative file paths in an embedded directory.
func collectEmbeddedPaths(root string) (map[string]bool, error) {
	paths := make(map[string]bool)
	err := fs.WalkDir(goforgo.Content, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			relPath, relErr := filepath.Rel(root, path)
			if relErr != nil {
				return relErr
			}
			paths[relPath] = true
		}
		return nil
	})
	return paths, err
}

// removeStaleFiles removes files from on-disk exercises/ and solutions/ that
// no longer exist in the embedded content, then prunes empty directories.
func removeStaleFiles(baseDir string) (int, error) {
	removed := 0

	for _, dir := range []string{"exercises", "solutions"} {
		embeddedPaths, err := collectEmbeddedPaths(dir)
		if err != nil {
			return removed, fmt.Errorf("failed to collect embedded %s paths: %w", dir, err)
		}

		diskDir := filepath.Join(baseDir, dir)
		if _, err := os.Stat(diskDir); os.IsNotExist(err) {
			continue
		}

		// Remove files not in the embedded set
		err = filepath.Walk(diskDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			relPath, relErr := filepath.Rel(diskDir, path)
			if relErr != nil {
				return relErr
			}
			if !embeddedPaths[relPath] {
				fmt.Printf("  🗑️  Removing stale file: %s/%s\n", dir, relPath)
				if rmErr := os.Remove(path); rmErr != nil {
					return rmErr
				}
				removed++
			}
			return nil
		})
		if err != nil {
			return removed, err
		}

		// Prune empty directories (walk in reverse depth order)
		_ = pruneEmptyDirs(diskDir)
	}

	return removed, nil
}

// pruneEmptyDirs removes empty directories under root, bottom-up.
func pruneEmptyDirs(root string) error {
	// Collect directories bottom-up
	var dirs []string
	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() && path != root {
			dirs = append(dirs, path)
		}
		return nil
	})

	// Remove in reverse order (deepest first)
	for i := len(dirs) - 1; i >= 0; i-- {
		entries, err := os.ReadDir(dirs[i])
		if err != nil {
			continue
		}
		if len(entries) == 0 {
			fmt.Printf("  🗑️  Removing empty directory: %s\n", dirs[i])
			_ = os.Remove(dirs[i])
		}
	}
	return nil
}

func shouldUpdateFile(srcPath, destPath, relPath string, completedExercises map[string]bool) bool {
	// Always update .toml files (configuration/metadata)
	if strings.HasSuffix(relPath, ".toml") {
		return true
	}

	// Always update non-.go files (README, scripts, etc.)
	if !strings.HasSuffix(relPath, ".go") {
		return true
	}

	// For .go files, extract exercise name from path and check completion
	exerciseName := extractExerciseName(relPath)
	if exerciseName == "" {
		// Can't determine exercise name, update it to be safe
		return true
	}

	// If the exercise is completed, don't update the .go file (preserve user's work)
	if completedExercises[exerciseName] {
		return false
	}

	// Exercise is not completed, safe to update
	return true
}

func extractExerciseName(relPath string) string {
	// Extract exercise name from path like "04_control_flow/range_loops.go"
	// Returns "range_loops"
	if strings.HasSuffix(relPath, ".go") {
		base := filepath.Base(relPath)
		return strings.TrimSuffix(base, ".go")
	}
	return ""
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
