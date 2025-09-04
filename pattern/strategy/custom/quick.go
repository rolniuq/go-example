package custom

import "fmt"

type QuickSort struct{}

func (q *QuickSort) Sort(nums []int) {
	q.helper(nums, 0, len(nums)-1)
	fmt.Println("sorted using quick sort: ", nums)
}

func (q *QuickSort) helper(nums []int, low, high int) {
	if low < high {
		p := q.partition(nums, low, high)
		q.helper(nums, low, p-1)
		q.helper(nums, p+1, high)
	}
}

func (q *QuickSort) partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low
	for j := low; j < high; j++ {
		if nums[i] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[high] = nums[high], nums[i]

	return i
}
