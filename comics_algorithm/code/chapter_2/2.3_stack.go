package chapter_2

//Stack based on array storage

type ArrayStack struct {
	storage *Array
}

func NewArrayStack(n int) *ArrayStack {
	return &ArrayStack{
		NewArray(n),
	}
}

func (as *ArrayStack) Push(val int) {
	as.storage.Insert(as.storage.Size(), val)
}

func (as *ArrayStack) Pop() int {
	return as.storage.Delete(as.storage.Size() - 1)
}

//Stack based on list storage

type ListStack struct {
	storage *SinglyLinkdedList
}

func NewListStack() *ListStack {
	return &ListStack{
		NewSinglyLinkdedList(),
	}
}

func (ls *ListStack) Push(val int) {
	ls.storage.Insert(ls.storage.Size(), val)
}

func (ls *ListStack) Pop() int {
	node := ls.storage.Delete(ls.storage.Size() - 1)
	return node.Data
}
