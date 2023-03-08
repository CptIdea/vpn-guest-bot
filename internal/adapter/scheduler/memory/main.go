package memoryScheduler

import (
	"errors"
	"sync"
	"time"

	"vpn-guest-bot/internal/core/interfaces"
)

// taskModel represents a scheduled taskModel.
type taskModel struct {
	name  string
	task  func() error
	runAt time.Time
}

// scheduler is an in-memory implementation of the Scheduler interface.
type scheduler struct {
	mutex sync.Mutex
	tasks map[string]*taskModel
}

func New() interfaces.Scheduler {
	s := &scheduler{tasks: make(map[string]*taskModel)}
	s.run()
	return s
}

// AddTask adds a new taskModel to the scheduler.
func (s *scheduler) AddTask(name string, task func() error, time time.Time) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check if a taskModel with the same name already exists.
	if _, ok := s.tasks[name]; ok {
		return errors.New("a taskModel with the same name already exists")
	}

	// Create a new taskModel and add it to the map of tasks.
	newTask := &taskModel{
		name:  name,
		task:  task,
		runAt: time,
	}
	s.tasks[name] = newTask

	return nil
}

// CancelTask cancels a taskModel with the given name.
func (s *scheduler) CancelTask(name string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.tasks, name)

}

// runTasks runs all tasks that are due to run.
func (s *scheduler) runTasks() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Run all tasks that are due to run and remove cancelled tasks from the map.
	for name, t := range s.tasks {

		// If the taskModel is due to run, run it.
		if time.Now().After(t.runAt) {
			if err := t.task(); err != nil {
				// If the taskModel returns an error, log it or handle it in some way.
			}
			// Remove the taskModel from the map of tasks.
			delete(s.tasks, name)
			continue
		}
	}
}

func (s *scheduler) run() {
	go func() {
		for {
			time.Sleep(time.Minute)
			s.runTasks()
		}
	}()
}
