package worker

import (
	"time"
	"errors"
	"context"
	"testing"
)

type SomeTest struct {
	Variable bool
	Done     chan bool
}

func compareResults(v *SomeTest) (bool, error) {
	worker := New(func(ctx context.Context, w *Worker, o *[]Iface) {
		if sb, ok := (*o)[0].(*SomeTest); ok {
			sb.Variable = true
			sb.Done <- true
		}
		w.Shutdown(nil)
	}, &[]Iface{
		v,
	})

	select {
	case <-time.After(5 * time.Second):
		return false, errors.New("TIMEOUT")
	case <-v.Done:
		return v.Variable, worker.Shutdown(nil)
	}

	return v.Variable, worker.Shutdown(nil)
}

func TestGoroutineAndChangeVariable(t *testing.T) {
	someVar := SomeTest{Variable: false, Done: make(chan bool)}
	boolVar, err := compareResults(&someVar)

	if boolVar != true {
		t.Fatalf("should modify variable\n")
	}

	if err != nil {
		t.Fatalf("should be nil\n")
	}
}
