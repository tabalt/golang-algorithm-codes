package chapter_2

type SinglyNode struct {
	Data int
	Next *SinglyNode
}

type SinglyLinkdedList struct {
	pHead *SinglyNode
	pLast *SinglyNode

	size int
}

func NewSinglyLinkdedList() *SinglyLinkdedList {
	return &SinglyLinkdedList{
		pHead: nil,
		pLast: nil,
		size:  0,
	}
}

func (sll *SinglyLinkdedList) Size() int {
	return sll.size
}

// T(n) = n = O(n)
// S(n) = 1 = O(1)
func (sll *SinglyLinkdedList) Get(idx int) *SinglyNode {
	if idx < 0 || idx >= sll.size {
		return nil
	}

	tmpNode := sll.pHead
	for i := 0; i < idx; i++ {
		tmpNode = tmpNode.Next
	}

	return tmpNode
}

// T(n) = 1 = O(1) 	// ignore possible lookup operations
// S(n) = 1 = O(1)
func (sll *SinglyLinkdedList) Insert(idx, val int) {
	if idx < 0 || idx > sll.size {
		return
	}

	insertNode := &SinglyNode{
		Data: val,
	}

	if sll.size == 0 {
		sll.pHead = insertNode
		sll.pLast = insertNode
	} else if idx == 0 {
		insertNode.Next = sll.pHead
		sll.pHead = insertNode
	} else if idx == sll.size {
		sll.pLast.Next = insertNode
		sll.pLast = insertNode
	} else {
		prevNode := sll.Get(idx - 1)
		insertNode.Next = prevNode.Next
		prevNode.Next = insertNode
	}

	sll.size++
}

// T(n) = n = O(1) 	// ignore possible lookup operations
// S(n) = 1 = O(1)
func (sll *SinglyLinkdedList) Delete(idx int) *SinglyNode {
	if idx < 0 || idx >= sll.size {
		return nil
	}

	var deletedNode *SinglyNode
	if idx == 0 {
		deletedNode = sll.pHead

		sll.pHead = sll.pHead.Next
	} else if idx == sll.size-1 {
		deletedNode = sll.pLast

		prevNode := sll.Get(idx - 1)
		prevNode.Next = nil
		sll.pLast = prevNode
	} else {
		prevNode := sll.Get(idx - 1)
		deletedNode = prevNode.Next

		prevNode.Next = deletedNode.Next
	}
	sll.size--

	return deletedNode

}
