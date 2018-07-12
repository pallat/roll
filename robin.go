package round

import (
	"sync"
)

func NewRoundRobin(n int) func() <-chan int {
	cur := 0
	ch := make(chan int, 1)
	seq := makeSeq(n)
	mux := sync.Mutex{}

	return func() <-chan int {
		mux.Lock()
		ch <- cur
		cur = seq[cur]
		mux.Unlock()
		return ch
	}
}

func makeSeq(n int) []int {
	seq := []int{}
	for i := 0; i < (n - 1); i++ {
		seq = append(seq, i+1)
	}
	seq = append(seq, 0)
	return seq
}
