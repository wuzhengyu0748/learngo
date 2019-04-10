package tree

func (node *Node) Traverse() {
	node.TraverseFunc(func (n *Node) {
		n.Print()
	})
}

func (node *Node) TraverseFunc (f func(n *Node)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(n *Node) {
			out <-n
		})
		close(out)
	}()
	return out
}