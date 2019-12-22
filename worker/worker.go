package worker

import (
	"context"
)

type Worker struct {
	ctx     context.Context
	cancel  context.CancelFunc
	chDone  chan bool
	stopped bool
}

type Callback func(ctx context.Context, w *Worker)

func New(f Callback) *Worker {
	ctx, cancel := context.WithCancel(context.Background())
	w := Worker{ctx: ctx, cancel: cancel, chDone: make(chan bool)}
	return (&w).doit(f)
}

func (this *Worker) doit(f func(ctx context.Context, w *Worker)) *Worker {
	go func() {
		for {
			select {
			case <-this.ctx.Done():
				this.chDone <- true
				return
			default:
				f(this.ctx, this)
			}
		}
	}()

	return this
}

func (this *Worker) Shutdown(ctx context.Context) error {
	if this.stopped {
		return nil
	}

	this.stopped = true

	ctxb := ctx
	if ctxb == nil {
		ctxb = context.Background()
	}

	this.cancel()

	select {
	case <-this.chDone:
		return nil
	case <-ctxb.Done():
		return ctxb.Err()
	}
}
