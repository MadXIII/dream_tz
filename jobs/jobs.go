package jobs

import (
	"sync"

	"github.com/madxiii/dream_tz/models"
)

type Job struct {
	ch chan models.Values
	wg *sync.WaitGroup
	mu *sync.Mutex
}

func newJob(size int) *Job {
	return &Job{
		ch: make(chan models.Values, size),
		wg: &sync.WaitGroup{},
		mu: &sync.Mutex{},
	}
}

func GetSum(vals []models.Values, limit int) int {
	j := newJob(len(vals))
	var sum int

	for i := 0; i < limit; i++ {
		go j.worker(&sum)
	}

	for i := 0; i < len(vals); i++ {
		j.wg.Add(1)
		j.ch <- vals[i]
	}

	close(j.ch)
	j.wg.Wait()

	return sum
}

func (j Job) worker(sum *int) {
	for job := range j.ch {
		j.mu.Lock()
		*sum += job.A + job.B
		j.wg.Done()
		j.mu.Unlock()
	}
}
