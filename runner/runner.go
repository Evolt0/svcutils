package runner

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
)

var monitoredSignals = []os.Signal{
	syscall.SIGHUP,
	syscall.SIGINT,
	syscall.SIGTERM,
	syscall.SIGQUIT,
}

type Runner struct {
	tasks []Task
}

type Task interface {
	Name() string
	Run(ctx context.Context) (err error)
}

func NewRunner(tasks ...Task) *Runner {
	return &Runner{tasks: tasks}
}

func (r *Runner) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, monitoredSignals...)

	defer func() { signal.Stop(quit); close(quit) }()

	go func() { <-quit; cancel() }()

	var wg sync.WaitGroup

	for _, t := range r.tasks {
		task := t

		wg.Add(1)

		go func() {
			defer wg.Done()
			logrus.Infof("%s is starting", task.Name())

			if err := task.Run(ctx); err != nil {
				logrus.Errorf("%s run with error: %v", task.Name(), err)
			}

			logrus.Infof("%s is stopped", task.Name())
			// Cancel the context to stop all other tasks.
			cancel()
		}()
	}

	wg.Wait()
}
