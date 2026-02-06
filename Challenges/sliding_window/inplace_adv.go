package main

import "fmt"

// Move all zeros to end while maintaining order of non-zeros
// Example: [0,1,0,3,12] → [1,3,12,0,0]
func moveZeros(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}

// Remove duplicates from sorted array, return new length
// Example: [1,1,2,2,3] → [1,2,3], return 3
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

// Merge two sorted arrays into first array
// nums1 has enough space to hold both
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

// Rotate array right by k steps in-place
// Example: [1,2,3,4,5], k=2 → [4,5,1,2,3]
func rotateRight(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse := func(arr []int, start, end int) {
		for start < end {
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end--
		}
	}
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func main3() {
	// Test your functions
	nums1 := []int{0, 1, 0, 3, 12}
	moveZeros(nums1)
	fmt.Println("After move zeros:", nums1)

	nums2 := []int{1, 1, 2, 2, 3, 3, 3}
	newLen := removeDuplicates(nums2)
	fmt.Printf("After remove duplicates: %v (len=%d)\n", nums2[:newLen], newLen)
}
