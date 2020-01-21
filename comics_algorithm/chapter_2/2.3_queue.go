package chapter_2

//Queue based on array storage

type ArrayQueue struct {
	storage *Array
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		NewArray(n),
	}
}

func (aq *ArrayQueue) Enqueue(val int) {
	aq.storage.Insert(aq.storage.Size(), val)
}

func (aq *ArrayQueue) Dequeue() int {
	return aq.storage.Delete(0)
}

//Queue based on list storage

type ListQueue struct {
	storage *SinglyLinkdedList
}

func NewListQueue() *ListQueue {
	return &ListQueue{
		NewSinglyLinkdedList(),
	}
}

func (lq *ListQueue) Enqueue(val int) {
	lq.storage.Insert(lq.storage.Size(), val)
}

func (lq *ListQueue) Dequeue() int {
	node := lq.storage.Delete(0)
	return node.Data
}
