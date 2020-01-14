package chapter_3

type BinaryHeap struct {
}

//将无序的完全二叉树调整成二叉堆，本质是所有非叶子节点依次下沉
func (bh *BinaryHeap) Build(arr []int) {

	//从最后一个非叶子节点开始，依次做下沉调整
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		bh.DownAdjustParent(arr, i, len(arr))
	}
}

//堆排序, 最大堆排出从小到大，最小堆排出从大到小的序列
//时间复杂度 O(n) + O(nlogn) = O(logn)
func (bh *BinaryHeap) Sort(arr []int) {
	//构建二叉堆
	bh.Build(arr)

	//循环删除堆顶元素，移动到未排序的尾部，调整产生新对顶
	for i := len(arr) - 1; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = arr[0]
		arr[0] = tmp

		bh.DownAdjustParent(arr, 0, i)
	}
}

// 父节点下沉
func (bh *BinaryHeap) DownAdjustParent(arr []int, parentIndex, length int) {
	childIndex := 2*parentIndex + 1
	for childIndex < length {
		//存在右孩子，且右孩子比左孩子小，则定位到右孩子
		if childIndex+1 < length && arr[childIndex+1] < arr[childIndex] {
			childIndex++
		}

		// 父节点小于等于小孩子，则不需要下沉
		if arr[parentIndex] <= arr[childIndex] {
			break
		}

		//交换小孩子节点和父节点的值
		tmp := arr[parentIndex]
		arr[parentIndex] = arr[childIndex]
		arr[childIndex] = tmp

		//继续比对孩子节点的孩子节点
		parentIndex = childIndex
		childIndex = 2*parentIndex + 1
	}
}

//子节点上浮
func (bh *BinaryHeap) UpAdjustChild(arr []int, childIndex int) {
	parentIndex := (childIndex - 1) / 2
	for childIndex > 0 {
		//子节点大于等于父节点，则不需要上浮
		if arr[childIndex] >= arr[parentIndex] {
			break
		}

		//交换孩子节点和父节点的值
		tmp := arr[childIndex]
		arr[childIndex] = arr[parentIndex]
		arr[parentIndex] = tmp

		//继续比对父亲节点的父亲节点
		childIndex = parentIndex
		parentIndex = (childIndex - 1) / 2
	}
}
