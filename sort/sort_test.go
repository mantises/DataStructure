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
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(200))
	}
	fmt.Println(nums)
	BubbleSort(nums)
	fmt.Println(nums)
}

func TestInsertSort(t *testing.T) {
	var nums []int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(100))
	}
	fmt.Println(nums)
	InsertSort(nums)
	fmt.Println(nums)
}

func TestSelectSort(t *testing.T) {
	var nums []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(200))
	}
	fmt.Println(nums)
	SelectSort(nums)
	fmt.Println(nums)
}

func TestQuickSort(t *testing.T) {
	var nums []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(200))
	}
	fmt.Println(nums)
	QuickSort(nums)
	fmt.Println(nums)
}

func TestMergeSort(t *testing.T) {
	var nums []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(200))
	}
	fmt.Println(nums)
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

func TestShellSort(t *testing.T) {
	var nums []int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(20)%(i+1))
	}
	fmt.Println(nums)
	ShellSort(nums)
	fmt.Println(nums)
}
