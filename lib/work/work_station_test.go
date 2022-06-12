package work

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWorkStation(t *testing.T) {
	type Result struct {
		idx int
	}
	assert := assert.New(t)
	workStation := NewWorkStation[int, Result]()

	inputs := []int{0, 1, 2, 3}

	workStation.Run(inputs, func(i int) (Result, error) {
		if i == 3 {
			return Result{}, errors.New("invalid input")
		}
		return Result{idx: i}, nil
	})

	time.Sleep(50 * time.Millisecond)

	expectedResults := []Result{
		{idx: 0},
		{idx: 1},
		{idx: 2},
	}
	assert.Equal(expectedResults, workStation.results)
	expectedErrors := []error{
		errors.New("invalid input"),
	}
	assert.Equal(expectedErrors, workStation.errors)
}

func TestWorkStationQuantity(t *testing.T) {
	type Result struct {
		idx int
	}
	assert := assert.New(t)
	workStation := NewWorkStation[int, Result]()
	workStation.SetWorkersQty(2)

	inputs := []int{0, 1}

	workStation.Run(inputs, func(i int) (Result, error) {
		if i == 0 {
			time.Sleep(30 * time.Millisecond)
		}

		return Result{idx: i}, nil
	})

	time.Sleep(50 * time.Millisecond)

	expectedResults := []Result{
		{idx: 1},
		{idx: 0},
	}
	assert.Equal(expectedResults, workStation.results)
}
