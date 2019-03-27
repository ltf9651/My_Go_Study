package quene

// type Quene []int
type Quene []interface{} //interface可以接受任何类型

func (q *Quene) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Quene) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Quene) IsEmpty() bool {
	return len(*q) == 0
}
