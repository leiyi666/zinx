package No23

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	if k == 0 {
		return nil
	} else if k == 1 {
		return lists[0]
	}
	arr := make([]*ListNode, 0)
	for i := 0; i < k; i++ {
		if lists[i] != nil {
			arr = append(arr, lists[i])
		} else {
			arr = append(arr, &ListNode{Val: 100000, Next: nil})
		}
	}
	createHeap(arr)
	head := ListNode{
		Val:  0,
		Next: nil,
	}
	var pre *ListNode
	pre = &head
	for arr[0].Val != 100000 {
		pre.Next = arr[0]
		if arr[0].Next != nil {
			arr[0] = arr[0].Next
		} else {
			arr[0] = &ListNode{Val: 100000, Next: nil}
		}
		heapSort(arr, 0, k)
		pre = pre.Next
	}
	return head.Next
}

func createHeap(arr []*ListNode) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapSort(arr, i, len(arr))
	}
}

func heapSort(arr []*ListNode, i int, length int) {
	temp := arr[i]
	for k := i*2 + 1; k < length; k = k*2 + 1 {
		if k+1 < length && arr[k].Val > arr[k+1].Val {
			k++
		}
		if arr[k].Val < temp.Val {
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}
	arr[i] = temp
}
