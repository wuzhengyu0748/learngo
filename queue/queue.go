package queue

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	last := (*q)[len(*q) - 1]
	*q = (*q)[:len(*q) - 1]
	return last
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}