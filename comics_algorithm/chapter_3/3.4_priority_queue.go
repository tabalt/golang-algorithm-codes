package chapter_3

type PriotiryQueue struct {
	*BinaryHeap

	storage []int
}

func NewPriotiryQueue() *PriotiryQueue {
	return &PriotiryQueue{
		BinaryHeap: &BinaryHeap{},
		storage:    []int{},
	}
}

// 入队
func (pq *PriotiryQueue) In(v int) {
	pq.storage = append(pq.storage, v)
	pq.UpAdjustChild(pq.storage, len(pq.storage)-1)
}

// 出队
func (pq *PriotiryQueue) Out() int {
	if len(pq.storage) <= 0 {
		return -999
	}

	top := pq.storage[0]

	//最后一个节点值补位根节点
	pq.storage[0] = pq.storage[len(pq.storage)-1]
	pq.storage = pq.storage[:len(pq.storage)-1]

	pq.DownAdjustParent(pq.storage, 0, len(pq.storage)-1)

	return top
}
