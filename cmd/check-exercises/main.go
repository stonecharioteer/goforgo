package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/BurntSushi/toml"
)

type metadata struct {
	Exercise exerciseInfo `toml:"exercise"`
}

type exerciseInfo struct {
	Name       string `toml:"name"`
	Category   string `toml:"category"`
	Order      int    `toml:"order"`
	Difficulty int    `toml:"difficulty"`
}

type report struct {
	totalExercises      int
	completeSets        int
	missingSolutions    []string
	missingTOMLs        []string
	orphanedSolutions   []string
	metadataErrors      []string
	categoryCounts      map[string]int
	categoryOrderByPath map[string]map[int]string
}

func main() {
	report, err := checkExercises("exercises", "solutions")
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Exercise check failed: %v\n", err)
		os.Exit(1)
	}

	printReport(report)

	if !report.ok() {
		os.Exit(1)
	}
}

func checkExercises(exercisesDir, solutionsDir string) (*report, error) {
	report := &report{
		categoryCounts:      make(map[string]int),
		categoryOrderByPath: make(map[string]map[int]string),
	}

	if err := filepath.WalkDir(exercisesDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".go" || strings.HasSuffix(strings.TrimSuffix(filepath.Base(path), ".go"), "_test") {
			return nil
		}

		relPath, err := filepath.Rel(exercisesDir, path)
		if err != nil {
			return err
		}
		category := filepath.Dir(relPath)
		name := strings.TrimSuffix(filepath.Base(path), ".go")

		report.totalExercises++

		solutionFile := filepath.Join(solutionsDir, category, name+".go")
		tomlFile := filepath.Join(exercisesDir, category, name+".toml")

		hasSolution := fileExists(solutionFile)
		hasTOML := fileExists(tomlFile)

		if !hasSolution {
			report.missingSolutions = append(report.missingSolutions, relPath)
		}
		if !hasTOML {
			report.missingTOMLs = append(report.missingTOMLs, relPath)
		} else {
			report.validateMetadata(tomlFile, category, name)
		}

		if hasSolution && hasTOML {
			report.completeSets++
			report.categoryCounts[category]++
		}

		return nil
	}); err != nil {
		return nil, err
	}

	if err := filepath.WalkDir(solutionsDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".go" || strings.HasSuffix(strings.TrimSuffix(filepath.Base(path), ".go"), "_test") {
			return nil
		}

		relPath, err := filepath.Rel(solutionsDir, path)
		if err != nil {
			return err
		}
		category := filepath.Dir(relPath)
		name := strings.TrimSuffix(filepath.Base(path), ".go")
		exerciseFile := filepath.Join(exercisesDir, category, name+".go")

		if !fileExists(exerciseFile) {
			report.orphanedSolutions = append(report.orphanedSolutions, relPath)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	for category, count := range report.categoryCounts {
		if count < 3 {
			report.metadataErrors = append(report.metadataErrors, fmt.Sprintf("%s has only %d complete exercise sets", category, count))
		}
	}

	sort.Strings(report.missingSolutions)
	sort.Strings(report.missingTOMLs)
	sort.Strings(report.orphanedSolutions)
	sort.Strings(report.metadataErrors)

	return report, nil
}

func (r *report) validateMetadata(path, category, name string) {
	var meta metadata
	if _, err := toml.DecodeFile(path, &meta); err != nil {
		r.metadataErrors = append(r.metadataErrors, fmt.Sprintf("%s: invalid TOML: %v", path, err))
		return
	}

	if meta.Exercise.Name != name {
		r.metadataErrors = append(r.metadataErrors, fmt.Sprintf("%s: exercise.name is %q, want %q", path, meta.Exercise.Name, name))
	}
	if meta.Exercise.Category != category {
		r.metadataErrors = append(r.metadataErrors, fmt.Sprintf("%s: exercise.category is %q, want %q", path, meta.Exercise.Category, category))
	}
	if meta.Exercise.Order <= 0 {
		r.metadataErrors = append(r.metadataErrors, fmt.Sprintf("%s: exercise.order must be positive", path))
	} else {
		orders := r.categoryOrderByPath[category]
		if orders == nil {
			orders = make(map[int]string)
			r.categoryOrderByPath[category] = orders
		}
		if existing, ok := orders[meta.Exercise.Order]; ok {
			r.metadataErrors = append(r.metadataErrors, fmt.Sprintf("%s: duplicate order %d also used by %s", path, meta.Exercise.Order, existing))
		} else {
			orders[meta.Exercise.Order] = path
		}
	}
	if meta.Exercise.Difficulty < 1 || meta.Exercise.Difficulty > 5 {
		r.metadataErrors = append(r.metadataErrors, fmt.Sprintf("%s: exercise.difficulty must be between 1 and 5", path))
	}
}

func printReport(r *report) {
	fmt.Println("🔍 GoForGo Exercise Completeness Check")
	fmt.Println("======================================")
	fmt.Println()
	fmt.Println("📊 Analyzing exercise files...")
	fmt.Println("🔍 Checking for orphaned solutions...")
	fmt.Println("📋 Exercise count by category:")
	fmt.Println("------------------------------")

	categories := make([]string, 0, len(r.categoryCounts))
	for category := range r.categoryCounts {
		categories = append(categories, category)
	}
	sort.Strings(categories)
	for _, category := range categories {
		readableCategory := strings.ReplaceAll(category, "_", " ")
		fmt.Printf("  %-25s: %d complete sets\n", readableCategory, r.categoryCounts[category])
	}

	fmt.Println()
	fmt.Println("📈 Summary Statistics")
	fmt.Println("====================")
	fmt.Printf("Total exercise files: %d\n", r.totalExercises)
	fmt.Printf("Complete exercise sets: %d\n", r.completeSets)
	fmt.Printf("Missing solutions: %d\n", len(r.missingSolutions))
	fmt.Printf("Missing TOML files: %d\n", len(r.missingTOMLs))
	fmt.Printf("Orphaned solutions: %d\n", len(r.orphanedSolutions))
	fmt.Printf("Metadata errors: %d\n", len(r.metadataErrors))
	if r.totalExercises > 0 {
		fmt.Printf("Completion rate: %.1f%%\n", float64(r.completeSets)*100/float64(r.totalExercises))
	}
	fmt.Println()

	printList("❌ Missing solution files:", r.missingSolutions)
	printList("❌ Missing TOML files:", r.missingTOMLs)
	printList("⚠️  Orphaned solution files:", r.orphanedSolutions)
	printList("❌ Metadata errors:", r.metadataErrors)

	if r.ok() {
		fmt.Println("✅ All exercises are complete! GoForGo is ready for production.")
	}
}

func printList(title string, items []string) {
	if len(items) == 0 {
		return
	}
	fmt.Println(title)
	for _, item := range items {
		fmt.Printf("  - %s\n", item)
	}
	fmt.Println()
}

func (r *report) ok() bool {
	return len(r.missingSolutions) == 0 &&
		len(r.missingTOMLs) == 0 &&
		len(r.orphanedSolutions) == 0 &&
		len(r.metadataErrors) == 0
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
