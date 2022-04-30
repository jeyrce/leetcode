package archive

// 最短单词
func findClosest(words []string, word1 string, word2 string) int {
	var (
		minLength int
		loc1      = make([]int, 0)
		loc2      = make([]int, 0)
	)
	for idx, item := range words {
		if word1 == item {
			loc1 = append(loc1, idx)
		}
		if word2 == item {
			loc2 = append(loc2, idx)
		}
	}
	for i := 0; i <= len(loc1); i++ {
		for j := 0; j <= len(loc2); j++ {
			c := loc1[i] - loc2[j]
			if c < 0 {
				c = -c
			}
			if c < minLength {
				minLength = c
			}
			if minLength < 2 {
				return minLength
			}
		}
	}
	return minLength
}
