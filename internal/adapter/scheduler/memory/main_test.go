package memoryScheduler

import (
	"testing"
	"time"
)

func TestMemoryScheduler_AddTask(t *testing.T) {
	// Create a new scheduler.
	s := &scheduler{
		tasks: make(map[string]*taskModel),
	}

	// Add a new taskModel to the scheduler.
	err := s.AddTask(
		"task1", func() error {
			return nil
		}, time.Now().Add(time.Minute),
	)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	// Try to add a taskModel with the same name.
	err = s.AddTask(
		"task1", func() error {
			return nil
		}, time.Now().Add(time.Minute),
	)
	if err == nil {
		t.Errorf("expected an error, but got none")
	}
}

func TestMemoryScheduler_CancelTask(t *testing.T) {
	// Create a new scheduler.
	s := &scheduler{
		tasks: make(map[string]*taskModel),
	}

	// Add a new taskModel to the scheduler.
	err := s.AddTask(
		"task1", func() error {
			return nil
		}, time.Now().Add(time.Minute),
	)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	// Cancel the taskModel.
	s.CancelTask("task1")

	// Try to cancel the taskModel again.
	s.CancelTask("task1")

	// Make sure the taskModel has been removed.
	if _, ok := s.tasks["task1"]; ok {
		t.Errorf("expected taskModel to be cancelled, but it's still in the map")
	}
}

func TestMemoryScheduler_RunTasks(t *testing.T) {
	// Create a new scheduler.
	s := &scheduler{
		tasks: make(map[string]*taskModel),
	}

	// Add a new taskModel to the scheduler that should be run.
	task1Run := false
	err := s.AddTask(
		"task1", func() error {
			task1Run = true
			return nil
		}, time.Now().Add(-time.Minute),
	)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	// Add a new taskModel to the scheduler that should be cancelled.
	err = s.AddTask(
		"task2", func() error {
			return nil
		}, time.Now().Add(-time.Minute),
	)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	// Run the tasks.
	s.runTasks()

	// Make sure the first taskModel has been run.
	if !task1Run {
		t.Errorf("expected task1 to be run, but it wasn't")
	}

	// Make sure the second taskModel has been removed.
	if _, ok := s.tasks["task2"]; ok {
		t.Errorf("expected task2 to be removed, but it's still in the map")
	}
}
