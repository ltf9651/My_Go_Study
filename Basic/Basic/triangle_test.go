// godoc -http :6060
package main

import (
	"My_Go_Study/Basic/Struct/quene"
	"fmt"
	"testing"
)

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{5, 12, 15},
	}
	for _, tt := range tests {
		if actual := calTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); "+
				"got %d; expected %d",
				tt.a, tt.b, actual, tt.c)
		}
	}
}

func BenchmarkTriangle(b *testing.B) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
	}
	for index := 0; index < b.N; index++ {
		for _, tt := range tests {
			if actual := calTriangle(tt.a, tt.b); actual != tt.c {
				b.Errorf("calcTriangle(%d, %d); "+
					"got %d; expected %d",
					tt.a, tt.b, actual, tt.c)
			}
		}
	}
}

func ExampleQuene() {
	q := quene.Quene{1}
	q.Push(1)
	q.Push(111)
	q.Push(11111)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())

	// Output:
	// false
	// 1
}
