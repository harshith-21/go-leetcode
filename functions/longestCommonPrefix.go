package functions

import "sort"

func LongestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
