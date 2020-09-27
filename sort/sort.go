package sort

// 冒泡排序
// 基本思路：比较相邻的两个元素，若逆序，则交换。对于每轮遍历，若没有发生遍历，则表明数组已有序
// 时间复杂度：o(N^2)
// 空间复杂度：o(1)
// 最差情况：N(N-1)/2 次
// 最好情况：N-1 次，数组已为有序，仅需一轮遍历即可
// 是否稳定：是
func BubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		flag := false
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				tmp := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = tmp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

// 插入排序
// 基本思路：0...i-1 已经有序，对于第 i 个元素，将前面比 nums[i] 大的依次后移，找到第 i 个元素的正确位置插入
// 时间复杂度：o(N^2)
// 空间复杂度：o(1)
// 最差情况：N(N-1)/2 次，即数组为逆序
// 最好情况：N-1 次，即数组已是有序
// 是否稳定：是
func InsertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		val := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > val {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = val
	}
}

// 选择排序
// 基本思路：对数组进行 N - 1 轮遍历，每轮选出最大的元素与尾部元素交换
// 时间复杂度：o(N^2)
// 空间复杂度：o(1)
// 最差情况：
// 最好情况：
// 是否稳定：否
func SelectSort(nums []int) {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		max := 0
		for j := 0; j < length-i; j++ {
			if nums[max] < nums[j] {
				max = j
			}
		}
		tmp := nums[length-i-1]
		nums[length-i-1] = nums[max]
		nums[max] = tmp
	}
}

// 快速排序
// 基本思路：基于分治法。设置一个数为枢轴，在每轮递归中根据枢轴将数组分为两部分，左半部分小于枢轴，右半部分大于枢轴。
//         将数组按枢轴左右两个子数组递归进行上述操作，每一次递归调用中都有一个元素都放在正确的位置，即枢轴。
// 时间复杂度：平均时间复杂度为 o(N*logN)，最差时间复杂度为 o(N^2)
// 空间复杂度：o(1)
// 最差情况：由于是递归将数组分为两个子数组，若左右两个子数组越均匀，则只需 logN 次递归调用，即复杂度为 N*logN。
//         数组已经有序（正序或者逆序），此时需要遍历 N 次，在递归中数组长度从 N 到 1。
// 最好情况：数组分布均匀，只需 logN 次遍历
// 是否稳定：否
func QuickSort(nums []int) {
	left, right := 0, len(nums)-1
	if left < right {
		pivot := nums[left]
		i, j := left, right
		for i < j {
			for i < j && nums[j] > pivot {
				j--
			}
			if i < j {
				nums[i] = nums[j]
				i++
			}
			for i < j && nums[i] < pivot {
				i++
			}
			if i < j {
				nums[j] = nums[i]
				j--
			}
		}
		nums[i] = pivot
		QuickSort(nums[:i])
		QuickSort(nums[j+1:])
	}
}

// 希尔排序
func ShellSort(nums []int) {

}

// 归并排序
// 基本思路：利用归并的思想实现的排序方法，该算法采用经典的分治（divide-and-conquer）策略。
// 分治法将问题分(divide)成一些小的问题然后递归求解，而治(conquer)的阶段则将分的阶段得到的
// 各答案"修补"在一起，即分而治之)。
// 时间复杂度：o(N*logN)，
// 空间复杂度：o(N)
// 是否稳定：是
func MergeSort(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(nums1, nums2 []int) []int {
	var tmp []int
	l1, l2 := 0, 0
	for l1 < len(nums1) && l2 < len(nums2) {
		if nums1[l1] <= nums2[l2] {
			tmp = append(tmp, nums1[l1])
			l1++
		} else {
			tmp = append(tmp, nums2[l2])
			l2++
		}
	}
	tmp = append(tmp, nums1[l1:]...)
	tmp = append(tmp, nums2[l2:]...)
	return tmp
}

// 堆排序
// 基本思路：将待排序序列构造成一个大顶堆，此时整个数组的最大值就是堆顶的根节点。将其与末尾元素进行
// 交换，此时末尾就为最大值。然后将剩余 N-1 个元素重新构造成一个大顶堆，这样会得到n个元素的次小值。
// 如此反复执行，便能得到一个有序序列
// 时间复杂度：o(N*logN)，
// 空间复杂度：o(1)
// 是否稳定：是
func HeapSort(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, len(nums))
	}
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		adjustHeap(nums, 0, i)
	}
}

func adjustHeap(nums []int, pos, length int) {
	for {
		child := pos*2 + 1
		if child > length-1 {
			break
		}
		if child < length-1 && nums[child+1] > nums[child] {
			child++
		}
		if nums[pos] < nums[child] {
			nums[pos], nums[child] = nums[child], nums[pos]
			pos = child
		} else {
			break
		}
	}
}
