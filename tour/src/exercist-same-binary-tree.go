package main

import "github.com/kangxie-colorado/golang-primer/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {
	t := tree.New(1)
	ch := make(chan int, 10)
	go Walk(t, ch)

	for e := range ch {
		print(e)
	}
}
