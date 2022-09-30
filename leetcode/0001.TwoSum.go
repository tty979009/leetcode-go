package leetcode

/*
题目:
	Given an array of integers, return indices of the two numbers such that they add up to a specific target.
	You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
	Given nums = [2, 7, 11, 15], target = 9,
	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1]

题目大意:
	在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。

解题思路:
	这道题最优的做法时间复杂度是 O(n)。
	顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，如果找到了，直接返回 2 个数字的下标即可。如果找不到，就把这个数字存入 map 中，等待扫到“另一半”数字的时候，再数值取出来返回结果。
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

/*
problem：
	1. 两种循环的区别
	2. 不需要使用else
	3. 返回nil
*/

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
