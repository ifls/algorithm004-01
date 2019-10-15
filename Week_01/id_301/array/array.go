package array

import "fmt"

func removeDuplicates(nums []int) int {
	length := len(nums)
	index := 0
	for i := 1; i < length; i++ {
		if nums[index] != nums[i] {
			index += 1
			nums[index] = nums[i]
		}
	}
	return index + 1
}

/**
解法思想 ： 当一个数组长度为n，移动后k次，当k<n的时候很好理解，当k>=n的时候，其实本质上是移动了k对n求余次数
*/
func RotateOne(nums []int, k int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	for i := 0; i < k; i++ {
		tmp := nums[length-1]
		for j := length - 2; j >= 0; j-- {
			nums[j+1] = nums[j]
		}
		nums[0] = tmp
	}
}

/**
解法思想 ： 移动次数简化同上面分析。然后我们从最数组的第一位开始，挨个一次将元素移动到目标位置，移动n次，那么所有元素都将到达最后的位置.
		  如果内层刚好移动到
*/
func RotateTwo(nums []int, k int) {
	length := len(nums)
	if length <= 1 || length == k || k == 0 {
		return
	}
	k = k % length
	count := 0
	for i := 0; count < length; i++ {
		curr := i
		pre := nums[i]
		for {
			next := (curr + k) % length
			temp := nums[next]
			nums[next] = pre
			pre = temp
			curr = next
			count++
			fmt.Println(curr, "|", i, "|", count)
			if curr == i {
				break
			}
		}
	}
}

func RotateThree(nums []int, k int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	k = k % length
	reverseArray(nums, 0, length-1)
	reverseArray(nums, 0, k-1)
	reverseArray(nums, k, length-1)
}

func reverseArray(nums []int, start, end int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	if start >= end {
		return
	}
	if start > length-1 {
		start = 0
	}

	if end > length-1 {
		end = length - 1
	}
	j := end
	for i := start; i < j; i++ {
		nums[i], nums[j] = nums[j], nums[i]
		j--
	}
}
