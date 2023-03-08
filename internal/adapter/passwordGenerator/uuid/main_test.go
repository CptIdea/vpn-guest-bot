package uuidPasswordGenerator

import (
	"testing"

	"github.com/google/uuid"
)

func TestGenerator_Generate(t *testing.T) {
	generator := New()

	password := generator.Generate()

	// Ensure that the generated password is a valid UUID v4.
	_, err := uuid.Parse(password)
	if err != nil {
		t.Errorf("Generator.Generate() produced an invalid UUID: %v", err)
	}
}
