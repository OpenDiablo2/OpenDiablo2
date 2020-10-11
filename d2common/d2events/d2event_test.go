package d2events

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func Test_EventEmitter_On(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	ee := NewEventEmitter()

	eventX := "x only"
	eventY := "y only"
	eventBoth := "both"

	var x, y int

	ee.On(eventX, func(args ...interface{}) {
		x++
	})

	ee.On(eventY, func(args ...interface{}) {
		y++
	})

	ee.On(eventBoth, func(args ...interface{}) {
		ee.Emit(eventX)
		ee.Emit(eventY)
	})

	ee.Emit(eventX)

	if x != 1 {
		t.Error("listener function not called")
	}

	if y != 0 {
		t.Error("listener function incorrectly called")
	}

	ee.Emit(eventY)
	ee.Emit(eventY)

	if x != 1 {
		t.Error("listener function incorrectly called")
	}

	if y != 2 {
		t.Error("listener function not called")
	}

	ee.Emit(eventBoth)

	if x != 2 {
		t.Error("listener function not called")
	}

	if y != 3 {
		t.Error("listener function not called")
	}
}

func Benchmark_EventEmitter(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	ee := NewEventEmitter()

	e1 := "testing"

	wg := &sync.WaitGroup{}

	for idx := 0; idx < b.N; idx++ {
		fn := func(args ...interface{}) {
			args[0].(*sync.WaitGroup).Done()
		}

		ee.Once(e1, fn)
		wg.Add(1)
	}

	ee.Emit(e1, wg)

	wg.Wait()

	if len(ee.listeners) > 0 {
		b.Error("listener count should be 0")
	}
}
