package main

import "fmt"

// Find maximum in each sliding window of size k
// Example: [1,3,-1,-3,5,3,6,7], k=3 → [3,3,5,5,6,7]
func slidingWindowMax(nums []int, k int) []int {
	if k == 0 || len(nums) == 0 {
		return []int{}
	}
	res := []int{}
	deque := []int{}
	for i := 0; i < len(nums); i++ {
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}

// Find minimum in each sliding window of size k
func slidingWindowMin(nums []int, k int) []int {
	if k == 0 || len(nums) == 0 {
		return []int{}
	}
	res := []int{}
	deque := []int{}
	for i := 0; i < len(nums); i++ {
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] > nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}

// Find minimum sum of subarray of length k
func minSumSubarray(nums []int, k int) int {
	if k == 0 || len(nums) < k {
		return 0
	}
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	minSum := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum < minSum {
			minSum = sum
		}
	}
	return minSum
}

// Count subarrays with sum equal to target
func countSubarraysWithSum(nums []int, target int) int {
	count := 0
	prefixSum := 0
	m := make(map[int]int)
	m[0] = 1
	for _, num := range nums {
		prefixSum += num
		if val, ok := m[prefixSum-target]; ok {
			count += val
		}
		m[prefixSum]++
	}
	return count
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	result := slidingWindowMax(nums, k)
	fmt.Printf("Sliding window max (k=%d): %v\n", k, result)

	minSum := minSumSubarray([]int{4, 2, 1, 7, 8, 1, 2, 8, 1, 0}, 3)
	fmt.Println("Min sum subarray:", minSum)
}
