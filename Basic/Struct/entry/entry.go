package main

import (
	"My_Go_Study/Basic/Struct/quene"
	"My_Go_Study/Basic/Struct/tree"
	"fmt"

	"golang.org/x/tools/container/intsets"
)

type MyTreeNode struct {
	node *tree.Node
}

func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := MyTreeNode{myNode.node.Left}
	right := MyTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func testSparse() {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(101110)
	fmt.Println(s.Has(1))
	fmt.Println(s.Has(11))
	fmt.Println(s)
}

func main() {
	var s tree.Node
	s = tree.Node{Value: 3}
	s.Left = &tree.Node{}
	s.Right = &tree.Node{9, nil, nil}
	s.Traverse()
	s.SetValue(8888)
	s.Print()
	fmt.Println(s)
	myS := MyTreeNode{&s}
	myS.postOrder()
	fmt.Println()

	fmt.Println("quene")
	q := quene.Quene{1}
	q.Push(2)
	q.Push("abc")
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q)
	q.Pop()
	fmt.Println(q.IsEmpty())
	q.Pop()
	fmt.Println(q.IsEmpty())

	fmt.Println("testSparse")
	testSparse()

	nodeCount := 0
	s.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("nodeCount:", nodeCount)

	c := s.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("maxNode:", maxNode)
}
