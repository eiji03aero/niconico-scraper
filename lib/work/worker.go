package work

type worker[Input, Result any] struct {
	inputCh  <-chan Input
	resultCh chan<- Result
	errCh    chan<- error
	cancelCh <-chan error
}

func newWorker[Input, Result any](
	inputCh <-chan Input,
	resultCh chan<- Result,
	errCh chan<- error,
	cancelCh <-chan error,
) *worker[Input, Result] {
	return &worker[Input, Result]{
		inputCh:  inputCh,
		resultCh: resultCh,
		errCh:    errCh,
		cancelCh: cancelCh,
	}
}

func (w *worker[Input, Result]) Run(workFn func(input Input) (Result, error)) {
	for {
		select {
		case input := <-w.inputCh:
			result, err := workFn(input)
			if err != nil {
				w.errCh <- err
			} else {
				w.resultCh <- result
			}
		case <-w.cancelCh:
			break
		}
	}
}
