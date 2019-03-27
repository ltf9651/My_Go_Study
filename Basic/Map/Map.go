package main

import "fmt"

func maps() {
	m := map[string]string{
		"name":     "leon",
		"language": "go",
		"d":        "ss",
	}
	fmt.Println(m)

	m2 := make(map[string]string) // var m2 map[string]string  m2 == nil
	fmt.Println(m2)
	for k, v := range m {
		fmt.Println(k, v) // 无序
	}

	name, ok := m["name"]
	fmt.Println(name, ok)
	na1me, ok := m["na1me"]
	fmt.Println(na1me, ok)
	if name, ok := m["na1me"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key doesn't exists")
	}

	delete(m, "d")
	fmt.Println(m)
}

func main() {
	//寻找最长不含有重复字符的字符串
	str := "啊是啊是的"
	lastOccured := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(str) {
		if lastOccured[ch] >= start {
			start = lastOccured[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	fmt.Println(maxLength)
}
