package storage

import (
	"context"
	"sync"
	"sync/atomic"
)

// Result represents the Size function result
type Result struct {
	// Total Size of File objects
	Size int64
	// Count is a count of File objects processed
	Count int64
}

type DirSizer interface {
	// Size calculate a size of given Dir, receive a ctx and the root Dir instance
	// will return Result or error if happened
	Size(ctx context.Context, d Dir) (Result, error)
}

// sizer implement the DirSizer interface
type sizer struct {
	// maxWorkersCount number of workers for asynchronous run
	//maxWorkersCount int
	//sem             *semaphore.Weighted
	// TODO: add other fields as you wish
}

// NewSizer returns new DirSizer instance
func NewSizer() DirSizer {
	return &sizer{}
}

func (a *sizer) Size(ctx context.Context, d Dir) (Result, error) {
	// TODO: implement this
	new_ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	dirs, files, err := d.Ls(new_ctx)
	if err != nil {
		return Result{}, err
	}

	r := Result{}
	for _, file := range files {
		size, err := file.Stat(new_ctx)
		if err != nil {
			return Result{}, err
		}
		atomic.AddInt64(&r.Size, size)
		atomic.AddInt64(&r.Count, 1)
	}

	errCh := make(chan error, 1)
	wg := new(sync.WaitGroup)

	for _, dir := range dirs {
		wg.Add(1)
		go func(dir Dir) {
			defer wg.Done()

			select {
			case <-new_ctx.Done():
				return
			default:
				res2, err := a.Size(new_ctx, dir)
				if err != nil {
					select {
					case errCh <- err: // в канале уже что-то было
						close(errCh)
						cancel()
						return
					default:
						return
					}
				}
				atomic.AddInt64(&r.Size, res2.Size)
				atomic.AddInt64(&r.Count, res2.Count)
			}
		}(dir)
	}
	wg.Wait()
	select {
	case err := <-errCh:
		return Result{}, err
	default:
		return r, nil
	}
}
