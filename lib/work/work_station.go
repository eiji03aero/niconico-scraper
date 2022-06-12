package work

import "fmt"

type workStation[Input, Result any] struct {
	workersQty int
	results    []Result
	errors     []error
}

func NewWorkStation[Input, Result any]() *workStation[Input, Result] {
	return &workStation[Input, Result]{
		workersQty: 1,
		results:    []Result{},
		errors:     []error{},
	}
}

func (w *workStation[Input, Result]) SetWorkersQty(qty int) {
	w.workersQty = qty
}

func (w *workStation[Input, Result]) Run(inputs []Input, workFn func(input Input) (Result, error)) {
	inputCh := make(chan Input)
	resultCh := make(chan Result)
	errCh := make(chan error)
	cancelCh := make(chan error)

	go func() {
		for {
			select {
			case result := <-resultCh:
				w.results = append(w.results, result)
			case err := <-errCh:
				w.errors = append(w.errors, err)
			}
		}
	}()

	fmt.Println("about to initialize workers")
	// initialize workers
	for i := 0; i < w.workersQty; i++ {
		worker := newWorker(inputCh, resultCh, errCh, cancelCh)
		go worker.Run(workFn)
	}

	fmt.Println("about to send")
	// send work to workers
	for _, input := range inputs {
		inputCh <- input
	}

	// close input channels
	cancelCh <- nil
}
