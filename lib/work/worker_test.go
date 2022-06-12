package work

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWorker(t *testing.T) {
	type Result struct {
		idx int
	}

	inputCh := make(chan int)
	resultCh := make(chan Result)
	errorCh := make(chan error)
	cancelCh := make(chan error)

	results := []Result{}
	errs := []error{}

	go func() {
		for {
			select {
			case result := <-resultCh:
				results = append(results, result)
			case err := <-errorCh:
				errs = append(errs, err)
			}
		}
	}()

	worker := newWorker[int, Result](inputCh, resultCh, errorCh, cancelCh)
	go worker.Run(func(input int) (Result, error) {
		if input == 3 {
			return Result{}, errors.New("invalid input")
		}
		return Result{idx: input}, nil
	})
	for i := 0; i < 4; i++ {
		inputCh <- i
	}

	time.Sleep(50 * time.Millisecond)
	cancelCh <- nil

	assert := assert.New(t)
	expectedResults := []Result{
		{idx: 0},
		{idx: 1},
		{idx: 2},
	}
	assert.Equal(expectedResults, results)
	expectedErrors := []error{
		errors.New("invalid input"),
	}
	assert.Equal(expectedErrors, errs)
}
