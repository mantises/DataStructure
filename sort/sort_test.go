package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestInsertSort(t *testing.T) {
	var nums []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 30; i ++ {
		nums = append(nums, rand.Intn(200))
	}
	InsertSort(nums)
	fmt.Println(nums)
}
