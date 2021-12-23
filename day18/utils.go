package day18

import (
	"fmt"
	"strconv"
)

var HEAD *Node

type Node struct {
	left, right, parent *Node
	depth               int
	val                 int
}

func (node *Node) String() string {
	if node.isSimple() {
		return fmt.Sprintf("%d", node.val)
	}
	return fmt.Sprintf("[%v,%v]", node.left, node.right)
}

func (node *Node) isSimple() bool {
	return node.left == nil && node.right == nil
}

func (node *Node) isLeft() bool {
	if node.parent == nil {
		return true
	}

	return node.parent.left == node
}

func (node *Node) leftMost() *Node {
	if node.isSimple() {
		return node
	}
	return node.left.leftMost()
}

func (node *Node) rightMost() *Node {
	if node.isSimple() {
		return node
	}
	return node.right.rightMost()
}

func (node *Node) leftMostInRight() *Node {
	if node.parent == nil {
		return nil
	}

	if node.parent.right != node {
		if node.parent.right.isSimple() {
			return node.parent.right
		}
		return node.parent.right.leftMost()
	} else {
		return node.parent.leftMostInRight()
	}
}

func (node *Node) rightMostInLeft() *Node {
	if node.parent == nil {
		return nil
	}

	if node.parent.left != node {
		if node.parent.left.isSimple() {
			return node.parent.left
		}
		return node.parent.left.rightMost()
	} else {
		return node.parent.rightMostInLeft()
	}
}

func (n *Node) join(other *Node) *Node {
	if n.parent != nil || other.parent != nil {
		panic("not a root node")
	}

	parent := &Node{
		left:  n,
		right: other,
	}
	n.parent = parent
	other.parent = parent

	nodes := make([]*Node, 0)
	nodes = append(nodes, n)
	nodes = append(nodes, other)
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		node.depth += 1
		if node.left != nil {
			nodes = append(nodes, node.left)
		}
		if node.right != nil {
			nodes = append(nodes, node.right)
		}
		node = nil
	}

	return parent
}

func (node *Node) explode() *Node {
	if node.isSimple() {
		return node
	}
	right := node.leftMostInRight()
	if right != nil {
		right.val += node.right.val
	}

	left := node.rightMostInLeft()
	if left != nil {
		left.val += node.left.val
	}

	new_node := &Node{parent: node.parent, val: 0, depth: node.depth}
	node = nil
	return new_node
}

func (node *Node) split() *Node {
	if node.val < 10 {
		return node
	}
	parent := &Node{
		parent: node.parent,
		depth:  node.depth,
	}
	left := &Node{
		val:    node.val / 2,
		parent: parent,
		depth:  node.depth + 1,
	}
	right := &Node{
		val:    (node.val + 1)/2,
		parent: parent,
		depth:  node.depth + 1,
	}

	parent.left = left
	parent.right = right
	node = nil
	return parent
}

func findExploding(nodes []*Node) *Node {
	for _, node := range nodes {
		if node.depth > 4 {
			return node.parent
		}
	}
	return nil
}

func findSpilt(nodes []*Node) *Node {
	for _, node := range nodes {
		if node.val >= 10 {
			return node
		}
	}
	return nil
}

func (node *Node) explore() {
	nodes := HEAD.flatten()
	modified := true
	for modified {
		modified = false
		node := findExploding(nodes)
		if node == nil {
			node = findSpilt(nodes)
			if node != nil {
				parent := node.parent
				if node.isLeft() {
					parent.left = node.split()
				} else {
					parent.right = node.split()
				}
				node = nil
				modified = true
			}
		} else {
			parent := node.parent
			if node.isLeft() {
				parent.left = node.explode()
			} else {
				parent.right = node.explode()
			}
			node = nil
			modified = true
		}

		if modified {
			nodes = HEAD.flatten()
		}
	}
}

func (node *Node) value() int {
	sum := 0
	if node.left.isSimple() {
		sum += 3 * node.left.val
	} else {
		sum += 3 * node.left.value()
	}

	if node.right.isSimple() {
		sum += 2 * node.right.val
	} else {
		sum += 2 * node.right.value()
	}
	return sum
}

func(node *Node) flatten() (nodes []*Node) {
	if node.left.isSimple() {
		nodes = append(nodes, node.left)
	} else {
		nodes = append(nodes, node.left.flatten()...)
	}

	if node.right.isSimple() {
		nodes = append(nodes, node.right)
	} else {
		nodes = append(nodes, node.right.flatten()...)
	}

	return
}

type stream struct {
	s   string
	idx int
}

func (s *stream) next() string {
	if s.idx >= len(s.s) {
		return ""
	}
	c := s.s[s.idx]
	s.idx++
	return string(c)
}

func newStream(s string) *stream {
	return &stream{
		s:   s,
		idx: 0,
	}
}

func newNode(s string) (*Node, error) {
	return processPair(newStream(s), nil, 0)
}

func processPair(s *stream, parent *Node, depth int) (*Node, error) {
	val := s.next()
	var err error
	switch val {
	case "[":
		node := &Node{
			depth:  depth,
			parent: parent,
		}
		node.left, err = processPair(s, node, depth+1)
		if err != nil {
			return nil, err
		}
		s.next()
		node.right, err = processPair(s, node, depth+1)
		if err != nil {
			return nil, err
		}
		s.next()
		return node, nil
	default:
		v, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		node := &Node{
			depth:  depth,
			parent: parent,
			val:    v,
		}
		return node, nil
	}
}
