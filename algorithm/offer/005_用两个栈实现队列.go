package main

type Queue struct {
	stack1 []int
	stack2 []int
}

func (q Queue) push(x int) {
	q.stack1 = append(q.stack1, x)
}

func (q Queue) pop() int {
	if len(q.stack2) == 0 {
		for _, x := range q.stack1 {
			q.stack2 = append(q.stack2, x)
		}
	}
	x := q.stack2[len(q.stack2)-1]
	q.stack2 = q.stack2[:len(q.stack2)-1]
	return x
}
