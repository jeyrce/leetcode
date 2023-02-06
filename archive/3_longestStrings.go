package archive

// 最长字符字串
// 
// 示例 1:
// 
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
// 
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
// 
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
// 
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters

// 思路: 从第一个开始，用map保存字符串，如果出现重复则map清零，并移动开始位置到下一个直到所有检查完
func lengthOfLongestSubstring(s string) int {
	var (
		max int
	)
	for i := 0; i <= len(s)-1; i++ {
		k := string(s[i])
		m := make(map[string]struct{}, len(s)-i-1)
		m[k] = struct{}{}
		// 检查是否有重复字符
		for j := i + 1; j <= len(s)-1; j++ {
			next := string(s[j])
			if _, ok := m[next]; !ok {
				m[next] = struct{}{}
				continue
			}
		}
		// 一轮走完检查字符串长度
		if cur := len(m); max < cur {
			max = cur
		}
	}
	return max
}
