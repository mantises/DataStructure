package sort

// 插入排序：
// 基本思路：0...i-1 已经有序，对于第 i 个元素，将前面比 nums[i] 大的依次后移，找到第 i 个元素的正确位置插入
// 时间复杂度：o(n^2)
// 空间复杂度：o(1)
// 最差情况比较次数：n(n-1)/2 次，即数组为逆序
// 最好情况比较次数：n-1 次，即数组已是有序
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

func SelectSort(nums []int) {

}

func QuickSort(nums []int) {

}

func ShellSort(nums []int) {

}

func HeapSort(nums []int) {

}
