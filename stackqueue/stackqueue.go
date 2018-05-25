package stackqueue

import (
	"github.com/golang-collections/collections/stack"
	"github.com/golang-collections/collections/queue"
)

type Queue struct {
	stack1 *stack.Stack
	stack2 *stack.Stack
}

func NewQueue() *Queue {
	queue := &Queue{
		stack1: stack.New(),
		stack2: stack.New(),
	}
	return queue
}

func (queue *Queue) Enqueue(ele interface{}) {
	queue.stack1.Push(ele)
}

func (queue *Queue) Dequeue() interface{} {
	if queue.stack2.Len() < 1 {

		for {
			if queue.stack1.Len() < 1 {
				break
			}
			queue.stack2.Push(queue.stack1.Pop())
		}
	}

	if queue.stack2.Len() < 1 {
		return nil
	}

	return queue.stack2.Pop()
}

type Stack struct {
	queue1 *queue.Queue
	queue2 *queue.Queue
}

func NewStack() *Stack {
	stack := &Stack{
		queue1: queue.New(),
		queue2: queue.New(),
	}
	return stack
}

func (stack *Stack) Push(ele interface{}) {
	// 同时空
	if stack.queue1.Len() < 1 && stack.queue2.Len() < 1 {
		stack.queue1.Enqueue(ele)
		return
	}

	// 往有数据的队列上加
	if stack.queue1.Len() < 1 {
		stack.queue2.Enqueue(ele)
	} else {
		stack.queue1.Enqueue(ele)
	}
}

func (stack *Stack) Pop() interface{} {
	// 同时空
	if stack.queue1.Len() < 1 && stack.queue2.Len() < 1 {
		return nil
	}

	var haveQueue *queue.Queue
	var noQueue *queue.Queue

	if stack.queue1.Len() < 1 {
		noQueue = stack.queue1
		haveQueue = stack.queue2
	} else {
		noQueue = stack.queue2
		haveQueue = stack.queue1
	}

	if haveQueue.Len() == 1 {
		return haveQueue.Dequeue()
	}

	haveQueueLen := haveQueue.Len()

	for i := 0; i < haveQueueLen-1; i++ {
		noQueue.Enqueue(haveQueue.Dequeue())
	}

	return haveQueue.Dequeue()
}
