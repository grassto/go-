package main

import (
	"fmt"
	"sort"
	"strings"
)

// 分割消息
type sm struct {
	Content string
	IsFile  bool
}

func main() {
	files := []string{
		"A", "B", "C",
	}
	if len(files) > 0 {
		// 包含文件，尝试分拆
		messageContent := "[A]B[A][C]EFG[D]DGH[A]"
		var fileIndexMap = make(map[int]string, 0)
		for _, fileID := range files {
			preLength := 0
			startIndex := strings.Index(messageContent, "["+fileID+"]")
			for startIndex >= 0 {
				fileIndexMap[startIndex+preLength] = fileID
				preLength = preLength + len(fileID) + 2
				startIndex = strings.Index(messageContent[preLength:], "["+fileID+"]")
			}
		}
		// 此时已经得到所有文件的位置
		// 排序
		var keys []int
		for key := range fileIndexMap {
			keys = append(keys, key)
		}
		sort.Ints(keys)
		// 分拆消息
		var sms []*sm
		preIndex := 0
		for _, key := range keys {
			if key > preIndex {
				sms = append(sms, &sm{
					Content: messageContent[preIndex:key],
					IsFile:  false,
				})
			}
			sms = append(sms, &sm{
				Content: fileIndexMap[key],
				IsFile:  true,
			})
			preIndex = key + len(fileIndexMap[key]) + 2
		}
		if len(messageContent[preIndex:]) > 0 {
			sms = append(sms, &sm{
				Content: messageContent[preIndex:],
				IsFile:  false,
			})
		}

		for _, sm := range sms {

			if sm.IsFile {

				fmt.Println("[" + sm.Content + "]")
			} else {
				fmt.Println(sm.Content)
			}

		}

	}
}
