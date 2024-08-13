package main

import "fmt"

type Queue struct {
	queue []int
}

func NewQueue() Queue {
	return Queue{queue: []int{}}
}

func (q *Queue) Enqueue(value int) {
	q.queue = append(q.queue, value)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.queue) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	} else if len(q.queue) == 1 {
		item := q.queue[0]
		q.queue = []int{}
		return item, nil
	}

	temp := q.queue[0]
	q.queue = q.queue[1:]
	return temp, nil
}

func main() {
	queueValues := NewQueue()
	queueValues.Enqueue(1)
	queueValues.Enqueue(2)
	fmt.Println(queueValues.queue)

	fmt.Println(queueValues.Dequeue())
	fmt.Println(queueValues.Dequeue())

}
