package main

import (
	"fmt"
	"math"
)

type Node[T any] struct {
	left  *Node[T]
	right *Node[T]
	val   T
}

type Tree[T any] struct {
	cmp  func(T, T) int
	root *Node[T]
}

func preorder[T any](node *Node[T]) {
	if node == nil {
		return
	}

	fmt.Println(node.val)

	preorder(node.left)
	preorder(node.right)
}

func (tr *Tree[T]) insert(root *Node[T], node Node[T]) *Node[T] {
	if tr.root == nil {
		tr.root = &node
	} else {
		if root == nil {
			return &node
		}

		cmp := tr.cmp(root.val, node.val)

		if cmp > 0 {
			root.right = tr.insert(root.right, node)
		}
		if cmp < 0 {
			root.left = tr.insert(root.left, node)
		}
	}

	return root
}

func Map[E any](s []E, f func(E, int, []E) E) []E {
	res := make([]E, 0)

	for i, v := range s {
		res = append(res, f(v, i, s))
	}

	return res
}

func Square(i float64) float64 {
	return math.Pow(i, float64(2))
}

func main() {
	c := make(chan *[16]byte, 1)

	c <- new([16]byte)
	close(c)

	s := []float64{1, 2, 3}

	fmt.Println(Map(s, func(v float64, i int, slice []float64) float64 {
		return Square(v)
	}))

	tr := Tree[int]{}

	tr.cmp = func(t1 int, t2 int) int {
		return t2 - t1
	}

	tr.insert(tr.root, Node[int]{
		val: 2,
	})

	tr.insert(tr.root, Node[int]{
		val: 1,
	})

	tr.insert(tr.root, Node[int]{
		val: 3,
	})

	preorder(tr.root)

	// go func() {
	// 	dA := new([16]byte)
	//
	// 	for false {
	// 		_, err := rand.Read(dA[:])
	//
	// 		if err != nil {
	// 			close(c)
	// 		}
	//
	// 		// c <- dA
	// 		// dA, dB = dB, dA
	// 	}
	// }()

	for data := range c {
		fmt.Println(data)
	}

}
