package round

import (
	"reflect"
	"sync"
	"testing"
)

func TestMakeSequence(t *testing.T) {
	seq := makeSeq(5)
	expected := []int{1, 2, 3, 4, 0}
	if !reflect.DeepEqual(expected, seq) {
		t.Fatalf("expected=%v\nseq%v\n", expected, seq)
	}
}

func TestRoundRobin(t *testing.T) {
	rr := NewRoundRobin(10)

	for i := 0; i < 10; i++ {
		if n := <-rr(); n != i {
			t.Fatalf("it should return seuence number 0..4 but at index %d got %d", i, n)
		}
	}
	for i := 0; i < 10; i++ {
		if n := <-rr(); n != i {
			t.Fatalf("it should return seuence number 0..4 but at index %d got %d", i, n)
		}
	}
}

func TestRoundRobinHitByConcurrency(t *testing.T) {
	rr := NewRoundRobin(10)

	var a, b, c int
	wg := sync.WaitGroup{}
	wg.Add(30)

	go func() {
		for i := 0; i < 10; i++ {
			a += <-rr()
			wg.Done()
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			b += <-rr()
			wg.Done()
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c += <-rr()
			wg.Done()
		}
	}()

	wg.Wait()

	sum := a + b + c

	if 135 != sum {
		t.Errorf("lost some\n(a=%d)\n (b=%d)\n (c=%d)\n", a, b, c)
	}
}

func TestRoundRobin2HitByConcurrency(t *testing.T) {
	rr := NewRoundRobin2(10)

	var a, b, c int
	wg := sync.WaitGroup{}
	wg.Add(30)

	go func() {
		for i := 0; i < 10; i++ {
			a += rr()
			wg.Done()
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			b += rr()
			wg.Done()
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c += rr()
			wg.Done()
		}
	}()

	wg.Wait()

	sum := a + b + c

	if 135 != sum {
		t.Errorf("lost some\n(a=%d)\n (b=%d)\n (c=%d)\n", a, b, c)
	}
}
func BenchmarkRR(b *testing.B) {
	rr := NewRoundRobin(10)

	for i := 0; i < b.N; i++ {
		<-rr()
	}
}

func BenchmarkRR2(b *testing.B) {
	rr := NewRoundRobin2(10)

	for i := 0; i < b.N; i++ {
		rr()
	}
}

func BenchmarkRRChan(b *testing.B) {
	rr := NewRoundRobinChan(10)

	for i := 0; i < b.N; i++ {
		<-rr()
	}
}
