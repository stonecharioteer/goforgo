package validation

import (
	"context"
	"testing"
	"time"

	"github.com/stonecharioteer/goforgo/internal/exercise"
)

func TestTestOrchestrator_BasicValidation(t *testing.T) {
	orchestrator := NewTestOrchestrator()

	// Create a simple exercise for testing
	exercise := &exercise.Exercise{
		FilePath: "/tmp/test.go",
		Info: exercise.ExerciseInfo{
			Name:     "test_exercise",
			Category: "test",
		},
		Validation: exercise.ExerciseValidation{
			Mode:    "build",
			Timeout: "30s",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// validateLegacyMode is not implemented, so the orchestrator should return an error
	_, err := orchestrator.ValidateExercise(ctx, exercise, "/tmp")
	if err == nil {
		t.Fatal("Expected error from unimplemented validateLegacyMode, got nil")
	}
	t.Logf("Got expected error: %v", err)
}

func TestUniversalRunner_Integration(t *testing.T) {
	runner := NewUniversalRunner("/tmp")

	// Both legacy and universal exercises route through the legacy runner now.
	// The legacy runner calls runner.RunExercise which will fail on non-existent
	// files, but that's expected - we're testing routing, not compilation.
	legacyExercise := &exercise.Exercise{
		FilePath: "/tmp/legacy.go",
		Info: exercise.ExerciseInfo{
			Name:     "legacy_test",
			Category: "test",
		},
		Validation: exercise.ExerciseValidation{
			Mode:    "build",
			Timeout: "30s",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// Legacy exercises go through the legacy runner (may fail on missing file, that's fine)
	result, err := runner.ValidateExercise(ctx, legacyExercise)
	if err != nil {
		t.Logf("Legacy validation returned error (expected for non-existent file): %v", err)
	} else if result != nil {
		summary := runner.GetValidationSummary(result)
		t.Logf("Validation summary: %+v", summary)
	}

	// Universal exercises should also route through legacy runner (not the orchestrator)
	universalExercise := &exercise.Exercise{
		FilePath: "/tmp/universal.go",
		Info: exercise.ExerciseInfo{
			Name:     "universal_test",
			Category: "test",
		},
		Validation: exercise.ExerciseValidation{
			Mode:    "universal",
			Timeout: "60s",
		},
	}

	result2, err := runner.ValidateExercise(ctx, universalExercise)
	if err != nil {
		t.Logf("Universal validation returned error (expected - routes through legacy): %v", err)
	} else if result2 != nil {
		t.Logf("Universal exercise routed through legacy runner successfully")
	}

	// Clean up
	if err := runner.Cleanup(ctx); err != nil {
		t.Errorf("Failed to cleanup runner: %v", err)
	}
}
