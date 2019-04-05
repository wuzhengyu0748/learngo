package queue

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	last := (*q)[len(*q) - 1]
	*q = (*q)[:len(*q) - 1]
	return last
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}