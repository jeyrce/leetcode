package archive

/**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 方法1：求和比较
func twoSum(nums []int, target int) []int {
	max := len(nums) - 1
	for i := 0; i <= max; i++ {
		// 如果只剩两个元素，则一定就是最后两个
		if i == max-1 {
			return []int{i, i + 1}
		}
		// 顺次将元素组合（不要重复排列）
		for j := i + 1; j <= max; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 方法二：利用map数据结构快速找到差值索引，值保存为key，索引保存为value
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for index, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{index, v}
		} else {
			m[num] = index
		}
	}
	return []int{}
}
