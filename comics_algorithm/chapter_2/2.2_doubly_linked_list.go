package chapter_2

type DoublyNode struct {
	Data int
	Prev *DoublyNode
	Next *DoublyNode
}

//doubly linked list with head node
type DoublyLinkdedList struct {
	Head *DoublyNode

	pLast *DoublyNode
	size  int
}

func NewDoublyLinkdedList() *DoublyLinkdedList {
	return &DoublyLinkdedList{
		Head: &DoublyNode{
			Prev: nil,
			Next: nil,
		},

		pLast: nil,
		size:  0,
	}
}

func (dll *DoublyLinkdedList) Size() int {
	return dll.size
}

// T(n) = n = O(n)
// S(n) = 1 = O(1)
func (dll *DoublyLinkdedList) Get(idx int) *DoublyNode {
	if idx < 0 || idx >= dll.size {
		return nil
	}

	tmpNode := dll.Head.Next
	for i := 0; i < idx; i++ {
		tmpNode = tmpNode.Next
	}

	return tmpNode
}

// T(n) = 1 = O(1) 	// ignore possible lookup operations
// S(n) = 1 = O(1)
func (dll *DoublyLinkdedList) Insert(idx, val int) {
	if idx < 0 || idx > dll.size {
		return
	}

	insertNode := &DoublyNode{
		Data: val,
	}

	if dll.pLast == nil {
		dll.pLast = insertNode
	}

	if idx == 0 {
		// 头插法
		if dll.Head.Next != nil {
			dll.Head.Next.Prev = insertNode
		}
		insertNode.Prev = dll.Head

		insertNode.Next = dll.Head.Next
		dll.Head.Next = insertNode
	} else if idx == dll.size {
		// 尾插法
		insertNode.Prev = dll.pLast

		dll.pLast.Next = insertNode
		dll.pLast = insertNode
	} else {
		prevNode := dll.Get(idx - 1)

		prevNode.Next.Prev = insertNode
		insertNode.Prev = prevNode

		insertNode.Next = prevNode.Next
		prevNode.Next = insertNode
	}

	dll.size++
}

// T(n) = n = O(1) 	// ignore possible lookup operations
// S(n) = 1 = O(1)
func (dll *DoublyLinkdedList) Delete(idx int) *DoublyNode {
	if idx < 0 || idx >= dll.size {
		return nil
	}

	var deletedNode *DoublyNode
	if idx == 0 {
		deletedNode = dll.Head.Next

		deletedNode.Next.Prev = dll.Head
		dll.Head.Next = deletedNode.Next
	} else if idx == dll.size-1 {
		deletedNode = dll.pLast

		prevNode := deletedNode.Prev
		prevNode.Next = nil
		dll.pLast = prevNode
	} else {
		prevNode := dll.Get(idx - 1)
		deletedNode = prevNode.Next
		deletedNode.Next.Prev = prevNode

		prevNode.Next = deletedNode.Next
	}
	dll.size--

	return deletedNode

}
