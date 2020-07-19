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
		b := []byte(s)

		// sort todo
		sort.Slice(b, func(i, j int) bool {
			return s[i] < s[j]
		})

		ss := string(b)
		index, ok := indexMap[ss]
		fmt.Println(indexMap)
		if ok {
			res[index] = append(res[index], strs[i])
		} else {
			indexMap[ss] = l
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

	for i, v := range ss {
		fmt.Println("No.", i, " = ", v)
	}

}
