package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"

	"github.com/sirupsen/logrus"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasksList []Task, numWorkers int, errorLimit int) error {
	logrus.Printf("tasksList.count=%d numWorkers=%d errorLimit=%d", len(tasksList), numWorkers, errorLimit)

	if len(tasksList) == 0 {
		return nil
	}

	if numWorkers < 1 {
		return nil
	}

	if errorLimit <= 0 {
		errorLimit = len(tasksList)
	}

	jobs := make(chan Task)
	errCh := make(chan error, len(tasksList))

	wg := sync.WaitGroup{}
	exitCh := make(chan struct{})
	defer func() {
		close(exitCh)
		wg.Wait()
	}()

	wg.Add(numWorkers)
	for w := 0; w < numWorkers; w++ {
		go worker(w, jobs, errCh, exitCh, &wg)
	}
	for _, task := range tasksList {
		jobs <- task

		if len(errCh) >= errorLimit {
			return ErrErrorsLimitExceeded
		}
	}

	return nil
}

func worker(w int, jobs <-chan Task, errCh chan<- error, exitCh chan struct{}, wg *sync.WaitGroup) {
	logrus.Printf("worker start %d ", w)
	for {
		select {
		case <-exitCh:
			wg.Done()
			return
		case t := <-jobs:
			if err := t(); err != nil {
				errCh <- err
			}
		}
	}
}
