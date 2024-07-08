package main

// 暴力法
// 时间复杂度 O(n^2)
// 空间复杂度 O(1)
func removeElement1(nums []int, val int) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		if nums[i] == val {
			for j := i + 1; j < size; j++ {
				nums[j-1] = nums[j]
			}
			i--
			size--
		}
	}
	return size
}

// 快慢指针法
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func removeElement2(nums []int, val int) int {
	// 初始化慢指针 slow
	slow := 0
	// 通过 for 循环移动快指针 fast
	// 当 fast 指向的元素等于 val 时，跳过
	// 否则，将该元素写入 slow 指向的位置，并将 slow 后移一位
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] == val {
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}

	return slow
}

// 相向双指针法
func removeElement3(nums []int, val int) int {
	// 有点像二分查找的左闭右闭区间 所以下面是<=
	left := 0
	right := len(nums) - 1
	for left <= right {
		// 不断寻找左侧的val和右侧的非val 找到时交换位置 目的是将val全覆盖掉
		for left <= right && nums[left] != val {
			left++
		}
		for left <= right && nums[right] == val {
			right--
		}
		//各自找到后开始覆盖 覆盖后继续寻找
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}
	return left
}

func main() {
	println(removeElement1([]int{3, 2, 2, 3}, 3))
	println(removeElement2([]int{3, 2, 2, 3}, 3))
	println(removeElement3([]int{3, 2, 2, 3}, 3))
}
