package Collections

type Queue[T any] []T

func (queue *Queue[T]) Clear() {
	*queue = Queue[T]{}
}

func (queue *Queue[T]) Enqueue(item T) {
	*queue = append(*queue, item)
}

func (queue *Queue[T]) Dequeue() (item T) {
	item, *queue = (*queue)[0], (*queue)[1:]
	return item
}

func (queue *Queue[T]) TryDequeue() (item T, ok bool) {
	if len(*queue) > 0 {
		item, *queue, ok = (*queue)[0], (*queue)[1:], true
	}
	return item, ok
}

func (queue Queue[T]) Peek() (item T) {
	return (queue)[0]
}

func (queue Queue[T]) TryPeek() (item T, ok bool) {
	if len(queue) > 0 {
		item, ok = (queue)[0], true
	}
	return item, ok
}

func (queue Queue[T]) Count() (count int) {
	return len(queue)
}

func (queue Queue[T]) GetEnumerator() IEnumerator[T] {
	return &enumerator[T]{
		index:  0,
		object: queue,
	}
}
