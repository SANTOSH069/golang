package main

type Queue struct {
	qu []int
}

func (q *Queue) AddFront(data int) {
	q.qu = append(q.qu, data)
}

func (q *Queue) AddLast(data int) {
	val := []int{}
	val = append(val, data)
	q.qu = append(q.qu, val...)
}

func (q *Queue) Dequeue() {
	ln := len(q.qu)
	q.qu = q.qu[:ln]
}

func (q *Queue) DelFirst() {
	q.qu = q.qu[1:]
}
