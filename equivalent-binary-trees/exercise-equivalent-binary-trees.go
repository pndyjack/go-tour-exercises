package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// to close channel, separate recursive call
	walkRecursively(t, ch)
	close(ch)
}

func walkRecursively(t *tree.Tree, ch chan int) {
	// assuming inorder traversal
	if t != nil {
		if t.Left != nil {
			walkRecursively(t.Left, ch)
		}
		ch <- t.Value
		if t.Right != nil {
			walkRecursively(t.Right, ch)
		}
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(tree1, tree2 *tree.Tree) bool {
	chTree1 := make(chan int, 10)
	chTree2 := make(chan int, 10)
	go Walk(tree1, chTree1)
	go Walk(tree2, chTree2)
	mapTree1 := getValuesFromChannel(chTree1)
	for valTree2 := range chTree2 {
		_, ok := mapTree1[valTree2]
		if len(mapTree1) == 0 || !ok {
			return false
		}
		delete(mapTree1, valTree2)
	}
	if len(mapTree1) != 0 {
		return false
	}
	return true
}

func getValuesFromChannel(ch chan int) map[int]int {
	values := make(map[int]int)
	for val := range ch {
		values[val] = val
	}
	return values
}

func main() {
	t := tree.Tree{
		Value: 2,
		Left: &tree.Tree{
			Value: 1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	fmt.Println(Same(&t, tree.New(2)))
}
