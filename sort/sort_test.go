package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	var nums []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 30; i++ {
		nums = append(nums, rand.Intn(200))
	}
	BubbleSort(nums)
	fmt.Println(nums)
}

func TestInsertSort(t *testing.T) {
	var nums []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 30; i++ {
		nums = append(nums, rand.Intn(200))
	}
	InsertSort(nums)
	fmt.Println(nums)
}

func TestSelectSort(t *testing.T) {
	var nums []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 30; i++ {
		nums = append(nums, rand.Intn(200))
	}
	SelectSort(nums)
	fmt.Println(nums)
}

func TestQuickSort(t *testing.T) {
	var nums []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 30; i++ {
		nums = append(nums, rand.Intn(200))
	}
	QuickSort(nums)
	fmt.Println(nums)
}

func TestMergeSort(t *testing.T) {
	var nums []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(200))
	}
	nums = MergeSort(nums)
	fmt.Println(nums)
}

func TestHeapSort(t *testing.T) {
	var nums []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(200))
	}
	fmt.Println(nums)
	HeapSort(nums)
	fmt.Println(nums)
}


