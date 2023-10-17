package question

import (
	"strings"
	"unicode"
)

// 请编写一个方法，将字符串中的空格全部替换为“%20”。
// 假定该字符串有足够的空间存放新增的字符，并且知道字符串的真实长度(小于等于20)，同时保证字符串由【大小写的英文字母组成】。 给定一个string为原始的串，返回替换后的string。

func replaceBlank(s string) (string, bool) {
	if len(s) > 20 {
		return "", false
	}

	for _, v := range s {
		if string(v) != " " && !unicode.IsLetter(v) {
			return " ", false
		}
	}

	return strings.ReplaceAll(s, " ", "%20"), true
}
