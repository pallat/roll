package round

import (
	"testing"
)

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
