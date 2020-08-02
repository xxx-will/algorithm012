package main

import "container/list"

// bfs 模版书写
// 提交后，时间超限 还需优化

func ladderLength(beginWord string, endWord string, wordList []string) int {

	l := list.New()

	l.PushBack(beginWord)
	visited := make(map[string]struct{})

	level := 0

	for l.Len() != 0 {
		lCache := make([]string, 0)
		for k := 0; k < l.Len(); k++ {
			w := l.Front()
			wStr := w.Value.(string)
			l.MoveToBack(w)
			for i := 0; i < len(wordList); i++ {
				cnt := 0
				for j := 0; j < len(wStr); j++ {
					if wStr[j] != wordList[i][j] {
						cnt++
					}
					if cnt > 1 {
						break
					}
				}

				if cnt == 1 {
					if wordList[i] == endWord {
						return level
					}
					_, ok := visited[wordList[i]]
					if !ok {
						lCache = append(lCache, wordList[i])
						visited[wordList[i]] = struct{}{}
					}
				}
			}
		}

		level++
		for _, v := range lCache {
			l.PushBack(v)
		}

	}

	return 0
}
