package slidingwindow

/*
438. 找到字符串中所有字母异位词 | 力扣  | https://leetcode.com/problems/find-all-anagrams-in-a-string/  |

给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

示例 1:

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
 示例 2:

输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
提示:

1 <= s.length, p.length <= 3 * 104
s 和 p 仅包含小写字母
*/

func findAnagrams(s, p string) []int {
	// boundary contitions
	if len(s) == 0 || len(p) == 0 {
		return nil
	}

	// times for the variable in the windows
	windows := make(map[rune]int)
	// number of the target item
	needs := make(map[rune]int)
	for _, c := range p {
		needs[c]++
	}
	// number of item that meet the requirement
	var valid int
	// result
	res := make([]int, 0)
	// left, right point
	left, right := 0, 0

	for right < len(s) {
		c := rune(s[right])
		right++

		if _, ok := needs[c]; ok {
			windows[c]++
			// variable meet the requirement
			if windows[c] == needs[c] {
				valid++
			}
		}

		for right-left >= len(p) {
			// find an anagrams
			if valid == len(needs) {
				res = append(res, left)
			}

			// move left
			d := rune(s[left])
			left++
			if _, ok := needs[d]; ok {
				// variable delete and not meet the requirement
				if windows[d] == needs[d] {
					valid--
				}
				windows[d]--
			}
		}
	}

	return res
}
