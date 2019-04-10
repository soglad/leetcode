/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 *
 * https://leetcode-cn.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (45.67%)
 * Total Accepted:    324K
 * Total Submissions: 709.5K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 * 
 * 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
 * 
 * 示例:
 * 
 * 给定 nums = [2, 7, 11, 15], target = 9
 * 
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 * 
 * 
 */
func twoSum(nums []int, target int) []int {
	v2idx := make(map[int][]int)
	for i, v := range nums {
		idx := v2idx[v]
		idx = append(idx, i)
		v2idx[v] = idx
	}
	for i, v := range nums {
		complet := target - v
		idx, ok := v2idx[complet]
		if ok {
			if len(idx) ==1 && idx[0] != i {
				return []int{i, idx[0]}
			}

			if len(idx) == 2 {
				return idx
			}
		}
	}
	return []int{}
}

