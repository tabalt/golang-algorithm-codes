package xiaohui

//singly linked list with head node
type HeadSinglyLinkdedList struct {
	Head  *SinglyNode
	pLast *SinglyNode

	Size int
}

func NewHeadSinglyLinkdedList() *HeadSinglyLinkdedList {
	return &HeadSinglyLinkdedList{
		Head: &SinglyNode{
			Next: nil,
		},
		pLast: nil,
		Size:  0,
	}
}

// T(n) = n = O(n)
// S(n) = 1 = O(1)
func (hsll *HeadSinglyLinkdedList) Get(idx int) *SinglyNode {
	if idx < 0 || idx >= hsll.Size {
		return nil
	}

	tmpNode := hsll.Head.Next
	for i := 0; i < idx; i++ {
		tmpNode = tmpNode.Next
	}

	return tmpNode
}

// T(n) = 1 = O(1) 	// ignore possible lookup operations
// S(n) = 1 = O(1)
func (hsll *HeadSinglyLinkdedList) Insert(idx, val int) {
	if idx < 0 || idx > hsll.Size {
		return
	}

	insertNode := &SinglyNode{
		Data: val,
	}
	if hsll.pLast == nil {
		hsll.pLast = insertNode
	}

	if idx == 0 {
		insertNode.Next = hsll.Head.Next
		hsll.Head.Next = insertNode
	} else if idx == hsll.Size {
		hsll.pLast.Next = insertNode
		hsll.pLast = insertNode
	} else {
		prevNode := hsll.Get(idx - 1)
		insertNode.Next = prevNode.Next
		prevNode.Next = insertNode
	}

	hsll.Size++
}

// T(n) = n = O(1) 	// ignore possible lookup operations
// S(n) = 1 = O(1)
func (hsll *HeadSinglyLinkdedList) Delete(idx int) *SinglyNode {
	if idx < 0 || idx >= hsll.Size {
		return nil
	}

	var deletedNode *SinglyNode
	if idx == 0 {
		deletedNode = hsll.Head.Next

		hsll.Head.Next = deletedNode.Next
	} else if idx == hsll.Size-1 {
		prevNode := hsll.Get(idx - 1)
		deletedNode = prevNode.Next

		prevNode.Next = nil
		hsll.pLast = prevNode
	} else {
		prevNode := hsll.Get(idx - 1)
		deletedNode = prevNode.Next

		prevNode.Next = deletedNode.Next
	}
	hsll.Size--

	return deletedNode

}
