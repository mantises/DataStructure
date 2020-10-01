package priority_queue

type FuncCmp func(x interface{}, y interface{}) bool

type PriorityQueue struct {
	val []interface{}
	cmp FuncCmp
}

func NewPriorityQueue(cmp FuncCmp) PriorityQueue {
	return PriorityQueue{
		val: make([]interface{}, 0, 0),
		cmp: cmp,
	}
}

func (h *PriorityQueue) Push(x interface{})  {
	h.val = append(h.val, x)
	for i := len(h.val)/2 - 1; i >= 0; i-- {
		h.adjustHeap(h.val, i, len(h.val))
	}
}

func (h *PriorityQueue) Pop() interface{} {
	if len(h.val) == 0 {
		return nil
	}
	ret := h.val[0]
	h.val = h.val[1:]
	for i := len(h.val)/2 - 1; i >= 0; i-- {
		h.adjustHeap(h.val, i, len(h.val))
	}
	return ret
}

func (h *PriorityQueue) adjustHeap(nums []interface{}, pos, length int) {
	for {
		child := pos*2 + 1
		if child > length-1 {
			break
		}
		if child < length-1 && h.cmp(nums[child+1], nums[child]) {
			child++
		}
		if h.cmp(nums[child], nums[pos]) {
			nums[pos], nums[child] = nums[child], nums[pos]
			pos = child
		} else {
			break
		}
	}
}

func (h *PriorityQueue) Sort(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		h.adjustHeap(h.val, 0, i)
	}
}
