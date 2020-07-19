package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {

	indexMap := make(map[string]int) //key: 拍完序的string  value: return 的index

	l := 0

	res := make([]([]string), 0)
	for i, s := range strs {
		fmt.Println("sort -pre", s)
		// todo
		sort.Slice([]byte(s), func(i, j int) bool {
			flag := s[i] < s[j]
			fmt.Println("flag is ", flag)
			return flag
		})
		fmt.Println("sort-", s)
		index, ok := indexMap[s]
		if ok {
			res[index] = append(res[index], strs[i])
		} else {
			indexMap[s] = l
			arr := []string{strs[i]}

			res = append(res, arr)
			l++
		}
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	ss := groupAnagrams(strs)

	fmt.Println(ss)
}
