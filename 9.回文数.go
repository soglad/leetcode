/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 *
 * https://leetcode-cn.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (55.90%)
 * Total Accepted:    94.8K
 * Total Submissions: 169.5K
 * Testcase Example:  '121'
 *
 * 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
 * 
 * 示例 1:
 * 
 * 输入: 121
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: -121
 * 输出: false
 * 解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
 * 
 * 
 * 示例 3:
 * 
 * 输入: 10
 * 输出: false
 * 解释: 从右向左读, 为 01 。因此它不是一个回文数。
 * 
 * 
 * 进阶:
 * 
 * 你能不将整数转为字符串来解决这个问题吗？
 * 
 */
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    d := 1
    for x/d >= 10 {
        d *= 10
    }
    for x != 0 {
        l := x / d
        r := x % 10
        if l != r {
            return false
        }
        x = (x % d) / 10 //remove first and last number
        d /= 100         //d need to down 2 level as 2 numbers has been removed 
    }
    return true
}

