package round

func NewRoundRobin(n int) func() <-chan int {
	seq := []int{}
	cur := 0
	ch := make(chan int, 1)

	for i := 0; i < (n - 1); i++ {
		seq = append(seq, i+1)
	}
	seq = append(seq, 0)

	return func() <-chan int {
		ch <- cur
		cur = seq[cur]
		return ch
	}
}
