package utils

import (
	"fmt"
	"time"
)

// NewCronJobRunner creates a new instance of CronJobRunner.
func NewCronJobRunner(interval time.Duration) *CronJobRunner {
	return &CronJobRunner{
		interval: interval,
		stop:     make(chan struct{}),
	}
}

// AddJob adds a new job to the cron job runner.
func (c *CronJobRunner) AddJob(job func()) {
	c.jobs = append(c.jobs, job)
}

// Start runs all the registered jobs periodically at the given interval.
func (c *CronJobRunner) Start() {
	c.wg.Add(1)

	go func() {
		defer c.wg.Done()

		for {
			select {
			case <-time.After(c.interval):
				for _, job := range c.jobs {
					go job()
				}
			case <-c.stop:
				return
			}
		}
	}()
}

// Stop stops the cron job runner.
func (c *CronJobRunner) Stop() {
	close(c.stop)
	c.wg.Wait()
}

// Example usage:
func main() {
	// Create a new cron job runner with a 2-second interval
	cronJobRunner := NewCronJobRunner(2 * time.Second)

	// Add your methods on the fly
	cronJobRunner.AddJob(func() {
		fmt.Println("Executing Method 1...")
		// Add logic for Method 1 here
	})

	cronJobRunner.AddJob(func() {
		fmt.Println("Executing Method 2...")
		// Add logic for Method 2 here
	})

	// Start the cron job runner
	cronJobRunner.Start()

	// Let it run for 10 seconds
	time.Sleep(10 * time.Second)

	// Stop the cron job runner
	cronJobRunner.Stop()
}
