package scheduler

import (
	"context"
	"log"
	"time"

	services "history-rate/service"
)

// Task is to hold task function
type TaskFn func(ctx context.Context) error

// ISchedulerService base methods scheduler service interface
type ISchedulerService interface {
	services.Service
	// Add adding task to scheduler by given task and delay time
	AddTask(t TaskFn, delay time.Duration, maxRetryFromPanic int)
}

type task struct {
	fn                TaskFn
	delay             time.Duration
	maxRetryFromPanic int
	countPanic        int
}

type schedulerService struct {
	// buffer of task channel
	ch chan *task
}

// NewService returns new scheduler service
func NewService() ISchedulerService {
	return &schedulerService{
		// create task buffer max 10 channel
		ch: make(chan *task, 10),
	}
}

// AddTask implementing ISchedulerService.AddTask
func (s *schedulerService) AddTask(fn TaskFn, delay time.Duration, maxRetryFromPanic int) {
	if maxRetryFromPanic > 1 {
		maxRetryFromPanic = maxRetryFromPanic - 1
	}
	t := &task{
		fn:                fn,
		delay:             delay,
		maxRetryFromPanic: maxRetryFromPanic,
	}

	s.ch <- t
}

// Run implementing services.Service
func (s *schedulerService) Run(ctx context.Context) error {
	log.Println("Scheduler service started...")
	for {
		select {
		case <-ctx.Done():
			return nil
		case t := <-s.ch:
			go s.createWorker(ctx, t)
		default:
			continue
		}
	}
}

// createWorker create new worker by given task and delay
func (s *schedulerService) createWorker(ctx context.Context, t *task) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			s.runFn(ctx, t)
			time.Sleep(t.delay)
			if (t.maxRetryFromPanic > -1) && (t.countPanic > t.maxRetryFromPanic) {
				log.Println("exit: max retry from panic reached")
				return
			}
		}
	}
}

func (s *schedulerService) runFn(ctx context.Context, t *task) {
	defer func(t *task) {
		p := recover()
		if p == nil {
			return
		}
		log.Println("job raised panic:", p)
		t.countPanic = t.countPanic + 1
	}(t)

	if err := t.fn(ctx); err != nil {
		log.Println("failed when run task with error", err)
	}
}
