package aux

import (
	"sort"
	"strings"
)

// 求交集
func InteractStrings(lists ...[]string) []string {
	for i := range lists {
		if len(lists[i]) == 0 {
			return nil
		}
		sort.Strings(lists[i])
	}
	return interactList(lists, 0, len(lists)-1)
}

// 求差集
func SubstractStrings(list1 []string, list2 []string) []string {
	sort.Strings(list1)
	sort.Strings(list2)
	var res []string
	var i, j int
	for i < len(list1) && j < len(list2) {
		if list1[i] < list2[j] {
			res = append(res, list1[i])
			i++
		} else if list1[i] == list2[j] {
			i++
			j++
		} else {
			j++
		}
	}
	if i < len(list1) {
		res = append(res, list1[i:]...)
	}
	return res
}

// 求并集
func UnionStrings(lists ...[]string) []string {
	for i := range lists {
		sort.Strings(lists[i])
	}
	return andList(lists, 0, len(lists)-1)
}

// 是否相等
func EqualStrings(list1, list2 []string) bool {
	if len(list1) != len(list2) {
		return false
	}
	sort.Strings(list1)
	sort.Strings(list2)
	for i := 0; i < len(list1); i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// 去重
func UniqStrings(list []string) []string {
	if len(list) <= 1 {
		return list
	}
	memo := make(map[string]int)
	for _, e := range list {
		memo[e] = 1
	}
	i := 0
	for str := range memo {
		list[i] = str
		i++
	}
	return list[:i]
}

// 是否包含
func ContainString(list []string, target string) bool {
	for _, e := range list {
		if e == target {
			return true
		}
	}
	return false
}

// 复制数组
func CopyStrings(list []string) []string {
	list2 := make([]string, len(list))
	copy(list2, list)
	return list2
}

// 移除某个字符串
func RemoveString(list []string, str string) []string {
	offset := 0
	for i, ele := range list {
		if ele == str {
			offset++
		} else if offset > 0 {
			list[i-offset] = ele
		}
	}
	return list[:len(list)-offset]
}

// 是否包含
func ContainStrings(list []string, sub []string) bool {
	memo := make(map[string]int)
	for _, e := range list {
		memo[e] = 1
	}
	for _, e := range sub {
		if _, ok := memo[e]; !ok {
			return false
		}
	}
	return true
}

// 文本是否包含任意特征串
func ContainsAny(text string, sub []string) bool {
	for _, str := range sub {
		if strings.Contains(text, str) {
			return true
		}
	}
	return false
}

// 是否包含，不区分大小写
func ContainStringIgnoreCase(list []string, target string) bool {
	target = strings.ToLower(target)
	for _, e := range list {
		if strings.ToLower(e) == target {
			return true
		}
	}
	return false
}

// 按大小分组
func PartitionStrings(list []string, size int) [][]string {
	var res [][]string
	for start := 0; start < len(list); start += size {
		end := start + size
		if end > len(list) {
			end = len(list)
		}
		res = append(res, list[start:end])
	}
	return res
}

// 转换元素
func MapStrings(list []string, f func(string) string) []string {
	for i, str := range list {
		list[i] = f(str)
	}
	return list
}

// 是否纯数字
func IsDigit(str string) bool {
	data := []byte(str)
	for _, b := range data {
		if b < 48 || b > 57 {
			return false
		}
	}
	return true
}

// below are helpers
func interactStrings(list1 []string, list2 []string) []string {
	if len(list1) == 0 || len(list2) == 0 {
		return nil
	}
	var res []string
	for i, j := 0, 0; i < len(list1) && j < len(list2); {
		if list1[i] == list2[j] {
			res = append(res, list1[i])
			i++
			j++
		} else if list1[i] > list2[j] {
			j++
		} else {
			i++
		}
	}
	return res
}

func interactList(lists [][]string, start, end int) []string {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	var left, right []string
	switch mid - start {
	case 0:
		left = lists[start]
	case 1:
		left = interactStrings(lists[start], lists[mid])
	default:
		left = interactList(lists, start, mid)
	}
	mid++
	if mid <= end {
		switch end - mid {
		case 0:
			right = lists[mid]
		case 1:
			right = interactStrings(lists[mid], lists[end])
		default:
			right = interactList(lists, mid, end)
		}
		return interactStrings(left, right)
	} else {
		return left
	}
}

func andList(lists [][]string, start, end int) []string {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	var left, right []string
	switch mid - start {
	case 0:
		left = lists[start]
	case 1:
		left = andStrings(lists[start], lists[mid])
	default:
		left = andList(lists, start, mid)
	}
	mid++
	if mid <= end {
		switch end - mid {
		case 0:
			right = lists[mid]
		case 1:
			right = andStrings(lists[mid], lists[end])
		default:
			right = andList(lists, mid, end)
		}
	}
	return andStrings(left, right)
}

func andStrings(list1 []string, list2 []string) []string {
	var res []string
	var i, j int
	for i < len(list1) && j < len(list2) {
		if list1[i] == list2[j] {
			res = append(res, list1[i])
			i++
			j++
		} else if list1[i] > list2[j] {
			res = append(res, list2[j])
			j++
		} else {
			res = append(res, list1[i])
			i++
		}
	}
	if i < len(list1) {
		res = append(res, list1[i:]...)
	}
	if j < len(list2) {
		res = append(res, list2[j:]...)
	}
	return res
}

// UnderlineLowercase 将大写单词转化成小写加下划线
func UnderlineLowercase(name string) string {
	data := []byte(name)
	var res []byte
	var i int
	for i < len(data) {
		if data[i] >= 65 && data[i] <= 90 {
			start := i
			i++
			for i < len(data) {
				if data[i] < 65 || data[i] > 90 {
					break
				}
				i++
			}
			res = append(res, byte(95))
			if i < len(data) && i-start >= 2 {
				res = append(res, data[start:i-1]...)
				res = append(res, byte(95), data[i-1])
			} else {
				res = append(res, data[start:i]...)
			}
			continue
		}
		res = append(res, data[i])
		i++
	}
	if len(res) > 0 && res[0] == byte(95) {
		res = res[1:]
	}
	for i, b := range res {
		if b >= 65 && b <= 90 {
			res[i] = b + 32
		}
	}
	return string(res)
}
