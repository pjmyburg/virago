package queue

import "testing"

func TestQueueSimple(t *testing.T) {
	q := New()

	for i := 0; i < minLength; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < minLength; i++ {
		if q.Peek().(int) != i {
			t.Error("peek", i, "had value", q.Peek())
		}
		x := q.Dequeue()
		if x != i {
			t.Error("dequeue", i, "had value", x)
		}
	}
}

func TestQueueWrapping(t *testing.T) {
	q := New()

	for i := 0; i < minLength; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 3; i++ {
		q.Dequeue()
		q.Enqueue(minLength + i)
	}

	for i := 0; i < minLength; i++ {
		if q.Peek().(int) != i+3 {
			t.Error("peek", i, "had value", q.Peek())
		}
		q.Dequeue()
	}
}

func TestQueueReturnToQueue(t *testing.T) {
	q := New()

	for i := 0; i < minLength; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 3; i++ {
		q.ReturnToQueue(minLength + i)
	}

	if q.Depth() != minLength+3 {
		t.Error("Expected Depth to be", minLength+1)
	}

	for i := 2; i >= 0; i-- {
		if q.Peek().(int) != minLength+i {
			t.Error("peek", i, "had value", q.Peek())
		}
		q.Dequeue()
	}
	for i := 0; i < minLength; i++ {
		if q.Peek().(int) != i {
			t.Error("peek", i, "had value", q.Peek())
		}
		q.Dequeue()
	}
}

func TestQueueReturnToQueueResize(t *testing.T) {
	q := New()

	for i := 0; i < minLength; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < minLength+1; i++ {
		q.ReturnToQueue(minLength + i)
	}

	if q.Depth() != (2*minLength)+1 {
		t.Error("Expected Depth to be", (2*minLength)+1)
	}

	for i := minLength; i >= 0; i-- {
		if q.Peek().(int) != minLength+i {
			t.Error("peek", i, "had value", q.Peek())
		}
		q.Dequeue()
	}
	for i := 0; i < minLength; i++ {
		if q.Peek().(int) != i {
			t.Error("peek", i, "had value", q.Peek())
		}
		q.Dequeue()
	}
}

func TestQueueLength(t *testing.T) {
	q := New()

	if q.Depth() != 0 {
		t.Error("empty queue length not 0")
	}

	for i := 0; i < 1000; i++ {
		q.Enqueue(i)
		if q.Depth() != i+1 {
			t.Error("adding: queue with", i, "elements has length", q.Depth())
		}
	}
	for i := 0; i < 1000; i++ {
		q.Dequeue()
		if q.Depth() != 1000-i-1 {
			t.Error("removing: queue with", 1000-i-i, "elements has length", q.Depth())
		}
	}
}

func TestQueueGet(t *testing.T) {
	q := New()

	for i := 0; i < 1000; i++ {
		q.Enqueue(i)
		for j := 0; j < q.Depth(); j++ {
			if q.get(j).(int) != j {
				t.Errorf("index %d doesn't contain %d", j, j)
			}
		}
	}
}

func TestQueueGetNegative(t *testing.T) {
	q := New()

	for i := 0; i < 1000; i++ {
		q.Enqueue(i)
		for j := 1; j <= q.Depth(); j++ {
			if q.get(-j).(int) != q.Depth()-j {
				t.Errorf("index %d doesn't contain %d", -j, q.Depth()-j)
			}
		}
	}
}

func TestQueueGetOutOfRangePanics(t *testing.T) {
	q := New()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assertPanics(t, "should panic when negative index", func() {
		q.get(-4)
	})

	assertPanics(t, "should panic when index greater than length", func() {
		q.get(4)
	})
}

func TestQueuePeekOutOfRangePanics(t *testing.T) {
	q := New()

	assertPanics(t, "should panic when peeking empty queue", func() {
		q.Peek()
	})

	q.Enqueue(1)
	q.Dequeue()

	assertPanics(t, "should panic when peeking emptied queue", func() {
		q.Peek()
	})
}

func TestQueueDequeueOutOfRangePanics(t *testing.T) {
	q := New()

	assertPanics(t, "should panic when removing empty queue", func() {
		q.Dequeue()
	})

	q.Enqueue(1)
	q.Dequeue()

	assertPanics(t, "should panic when removing emptied queue", func() {
		q.Dequeue()
	})
}

func assertPanics(t *testing.T, name string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%s: didn't panic as expected", name)
		}
	}()

	f()
}

// General warning: Go's benchmark utility (go test -bench .) increases the number of
// iterations until the benchmarks take a reasonable amount of time to run; memory usage
// is *NOT* considered. On my machine, these benchmarks hit around ~1GB before they've had
// enough, but if you have less than that available and start swapping, then all bets are off.

func BenchmarkQueueSerial(b *testing.B) {
	q := New()
	for i := 0; i < b.N; i++ {
		q.Enqueue(nil)
	}
	for i := 0; i < b.N; i++ {
		q.Peek()
		q.Dequeue()
	}
}

func BenchmarkQueueGet(b *testing.B) {
	q := New()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.get(i)
	}
}

func BenchmarkQueueTickTock(b *testing.B) {
	q := New()
	for i := 0; i < b.N; i++ {
		q.Enqueue(nil)
		q.Peek()
		q.Dequeue()
	}
}
