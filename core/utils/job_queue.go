package utils

import (
	"fmt"
	"reflect"
)

func NewJobProcessor(numWorkers int) *JobProcessor {
	return &JobProcessor{
		JobQueue:   make(chan Job),
		Shutdown:   make(chan struct{}),
		IsShutdown: false,
	}
}

func (jp *JobProcessor) Enqueue(job Job) error {
	jp.Mu.Lock()
	defer jp.Mu.Unlock()

	if jp.IsShutdown {
		return fmt.Errorf("job processor is shutdown, cannot enqueue new jobs")
	}

	jp.JobQueue <- job
	return nil
}

func (jp *JobProcessor) StartWorkers(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		jp.Wg.Add(1)
		go jp.Worker()
	}
}

func (jp *JobProcessor) StopWorkers() {
	close(jp.JobQueue)
	close(jp.Shutdown)
	jp.Wg.Wait()
}

func (jp *JobProcessor) Wait() {
	jp.Wg.Wait()
}

func (jp *JobProcessor) Worker() {
	defer jp.Wg.Done()

	for {
		select {
		case job, ok := <-jp.JobQueue:
			if !ok {
				return
			}
			jp.processJob(job)
		case <-jp.Shutdown:
			return
		}
	}
}

func (jp *JobProcessor) processJob(job Job) ([]reflect.Value, error) {
	functionValue := reflect.ValueOf(job.Function)
	if functionValue.Kind() != reflect.Func {
		return nil, fmt.Errorf("error: invalid function type for job %s", job.ID)

	}

	args := make([]reflect.Value, len(job.Args))
	for i, arg := range job.Args {
		args[i] = reflect.ValueOf(arg)
	}

	resultValues := functionValue.Call(args)

	if len(resultValues) > 0 {
		return resultValues, nil
	} else {
		return nil, nil
	}
}
