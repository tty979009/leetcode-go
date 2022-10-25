package leetcode

/*
problem：
	1. 两种循环的区别
	2. 不需要使用else
	3. 返回nil
*/

func TwoSumM(nums []int, target int) []int {
	m := make(map[int]int)
	for i, value := range nums {
		if v, ok := m[target-value]; ok {
			return []int{v, i}
		} else {
			m[value] = i
		}
	}
	return []int{}
}

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}
