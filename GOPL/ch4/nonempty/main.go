package main

import "fmt"

func main() {
	data := []string{"one", "", "three", "", "four"}
	fmt.Printf("%q\n", nonempty2(data))

	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
