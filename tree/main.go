package main

const (
	defaultValue = -1
)

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func (t *Tree) insertNode(node *Tree) *Tree {
	if t == nil {
		return node
	}

	if t.Val == defaultValue {
		t.Val = node.Val
		return t
	}

	if node.Val > t.Val {
		if t.Right == nil {
			t.Right = &Tree{Val: defaultValue}
		}
		t.Right = t.Right.insertNode(node)
	}
	if node.Val < t.Val {
		if t.Left == nil {
			t.Left = &Tree{Val: defaultValue}
		}
		t.Left = t.Left.insertNode(node)
	}

	return t
}

func (t *Tree) Create(values []int) *Tree {
	root := &Tree{Val: defaultValue, Left: nil, Right: nil}
	for _, value := range values {
		node := &Tree{Val: value}
		root.insertNode(node)
	}

	return root
}

func (t *Tree) Print() {
	if t == nil {
		return
	}

	t.Left.Print()
	println(t.Val)
	t.Right.Print()
}

func main() {
	tree := &Tree{}
	tree = tree.Create([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	tree.Print()
}
