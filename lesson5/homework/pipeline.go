package executor

import (
	"context"
)

type (
	In  <-chan any
	Out = In
)

type Stage func(in In) (out Out)

func ExecutePipeline(ctx context.Context, in In, stages ...Stage) Out {
	// TODO
	in2 := make(chan any)

	go func() {
		arr := make([]any, 0)
		for {
			select {
			case <-ctx.Done():
				close(in2)
				return
			case i, ok := <-in:
				if !ok {
					for _, elem := range arr {
						in2 <- elem
					}
					close(in2)
					return
				}
				arr = append(arr, i)
			}
		}
	}()

	var in3 <-chan any = in2
	for _, stage := range stages {
		in3 = stage(in3)
	}
	return in3
}
