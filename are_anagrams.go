package sprint

import (
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	normalize := func(s string) string {
		s = strings.ToLower(s)
		s = strings.ReplaceAll(s, " ", "")
		return s
	}

	s1 := normalize(str1)
	s2 := normalize(str2)
	runeS1 := []rune(s1)
	runeS2 := []rune(s2)
	sort.Slice(runeS1, func(i, j int) bool { return runeS1[i] < runeS1[j] })
	sort.Slice(runeS2, func(i, j int) bool { return runeS2[i] < runeS2[j] })
	return string(runeS1) == string(runeS2)
}
