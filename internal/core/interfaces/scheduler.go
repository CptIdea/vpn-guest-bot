package interfaces

import "time"

type Scheduler interface {
	AddTask(name string, task func() error, time time.Time) error
	CancelTask(name string)
}
